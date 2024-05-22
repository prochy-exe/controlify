package main

import (
	_ "controlify/assets"
	"controlify/keymap"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"fyne.io/systray"
	"github.com/gen2brain/beeep"
	"github.com/gorilla/websocket"
	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
)

//go:embed assets/icon.ico
var windowsIconData []byte

//go:embed assets/icon.png
var linuxIconData []byte

var (
	clients     = make(map[*websocket.Conn]string)
	mu          sync.Mutex
	clientsByID = make(map[string]*websocket.Conn)
	exePath     = filepath.Dir(func() string { p, _ := os.Executable(); return p }())
	upgrader    = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	imgPath       string
	isCLI         string
	iconData      []byte
	beeepIconName string
)

const (
	beeepTitle = "Controlify"
)

func createTempImage() (string, error) {
	// Get the temporary directory
	tmpDir := os.TempDir()

	// Construct the path for the temporary file
	tmpFilePath := filepath.Join(tmpDir, beeepIconName)

	// Write embedded image data to the temporary file
	err := os.WriteFile(tmpFilePath, iconData, 0644)
	if err != nil {
		return "", fmt.Errorf("error writing to temporary file: %v", err)
	}

	// Return the path to the temporary file
	return tmpFilePath, nil
}

func runWebSocketServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleConnections)

	server := &http.Server{
		Addr:    "localhost:8999",
		Handler: mux,
	}

	go sendHeartbeats()

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sendHeartbeats() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		mu.Lock()
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte("heartbeat"))
			if err != nil {
				removeClient(client)
			}
		}
		mu.Unlock()
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading to WebSocket: %v", err)
		return
	}

	defer func() {
		mu.Lock()
		defer mu.Unlock()
		delete(clients, ws)
		ws.Close()
	}()

	var initialMsg map[string]interface{}
	err = ws.ReadJSON(&initialMsg)
	if err != nil {
		log.Printf("error reading initial message: %v", err)
		return
	}

	log.Print(initialMsg)

	clientID, ok := initialMsg["clientID"].(string)
	if !ok || clientID == "" {
		log.Printf("clientID missing or invalid")
		return
	}

	if isCLI != "true" {
		if clientID == "spicetify-client" {
			beeep.Notify(beeepTitle, "Spotify client hooked!", imgPath)
		} else if clientID == "deej-client" {
			beeep.Notify(beeepTitle, "Deej connected!", imgPath)
		} else {
			beeep.Alert(beeepTitle, fmt.Sprintf("Unknown client connected with ID %v!", clientID), imgPath)
		}
	}

	clientsByID[clientID] = ws

	mu.Lock()
	clients[ws] = clientID
	mu.Unlock()

	for {
		var msg map[string]interface{}
		err := ws.ReadJSON(&msg)
		if err != nil {
			break
		}
		log.Printf("received message: %v %v", msg, clientID)
		forwardMessage(msg)
	}
}

func forwardMessage(msg map[string]interface{}) {
	targetClientID := "spicetify-client"

	mu.Lock()
	targetClient, ok := clientsByID[targetClientID]
	mu.Unlock()
	if !ok {
		log.Printf("target client not found: %s", targetClientID)
		return
	}

	err := targetClient.WriteJSON(msg)
	if err != nil {
		log.Printf("error sending message to client %s: %v", targetClientID, err)
	}
}

func removeClient(client *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()
	client.Close()
	delete(clients, client)
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--cli" {
		isCLI = "true"
	}

	if runtime.GOOS == "windows" {
		iconData = windowsIconData
		beeepIconName = "icon.ico"
	} else {
		iconData = linuxIconData
		beeepIconName = "icon.png"
	}

	imgPath, _ = createTempImage()

	config, err := loadConfig(filepath.Join(exePath, "config.json"))
	if err != nil {
		errMsg := "Config not found, hotkey function is disabled"
		if isCLI != "true" {
			beeep.Alert(beeepTitle, errMsg, imgPath)
		} else {
			log.Print(errMsg)
		}
	} else {
		initHotkeys(config)
	}

	// Connect to WebSocket server
	go runWebSocketServer()

	if len(os.Args) > 1 && os.Args[1] == "--cli" || isCLI == "true" {
		select {}
	} else if len(os.Args) > 1 && os.Args[1] != "" {
		log.Printf("Run executable with --cli to run without system tray and in console.")
	} else {
		// Initialize system tray
		systray.Run(onReady, onExit)
	}
}

func initHotkeys(config map[string]string) {
	// Set up hotkey
	mainthread.Init(func() {
		for event, shortcut := range config {
			if shortcut == "" {
				continue
			} else {
				go func(event, shortcut string) {
					if err := registerHotkey(event, shortcut); err != nil {
						log.Printf("Error registering hotkey for %s: %v", event, err)
					}
				}(event, shortcut)
			}
		}
	})
}

func registerHotkey(event, shortcut string) error {
	// Parse shortcut string
	modifiers, key, err := parseShortcut(shortcut)
	if err != nil {
		return err
	}

	// Register hotkey
	hk := hotkey.New(modifiers, key)
	err = hk.Register()
	if err != nil {
		return err
	}

	log.Printf("Registered hotkey for %s: %s\n", event, shortcut)
	for {
		<-hk.Keyup()
		msg := map[string]interface{}{event: ""}
		forwardMessage(msg)
	}
}

func parseShortcut(shortcut string) ([]hotkey.Modifier, hotkey.Key, error) {
	parts := strings.Split(shortcut, "+")
	if len(parts) == 0 {
		return nil, 0, fmt.Errorf("invalid shortcut: %s", shortcut)
	}

	var modifiers []hotkey.Modifier
	for _, part := range parts[:len(parts)-1] {
		mod, err := keymap.ParseModifier(part)
		if err != nil {
			return nil, 0, err
		}
		modifiers = append(modifiers, mod)
	}

	key, err := keymap.ParseKey(parts[len(parts)-1])
	if err != nil {
		return nil, 0, err
	}

	return modifiers, key, nil
}

func loadConfig(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var shortcuts map[string]string
	err = json.NewDecoder(file).Decode(&shortcuts)
	if err != nil {
		return nil, err
	}

	return shortcuts, nil
}

func onReady() {
	err := beeep.Notify(beeepTitle, "Controlify running as a tray app", imgPath)
	if err != nil {
		panic(err)
	}

	// Set up system tray
	systray.SetIcon(iconData)
	systray.SetTitle(beeepTitle)
	systray.SetTooltip(beeepTitle)

	// Add menu items
	systray.AddMenuItem("Controlify", "")
	systray.AddSeparator()
	mQuit := systray.AddMenuItemCheckbox("Quit", "", false)
	go func() {
		for range mQuit.ClickedCh {
			systray.Quit()
		}
	}()
}

func onExit() {
	os.Exit(0)
}
