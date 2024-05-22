//go:build windows

package keymap

import (
	"fmt"
	"strings"

	"golang.design/x/hotkey"
)

func ParseModifier(modifier string) (hotkey.Modifier, error) {
	switch strings.ToLower(modifier) {
	case "ctrl":
		return hotkey.ModCtrl, nil
	case "shift":
		return hotkey.ModShift, nil
	case "alt":
		return hotkey.ModAlt, nil
	case "win":
		return hotkey.ModWin, nil
	default:
		return 0, fmt.Errorf("invalid modifier: %s", modifier)
	}
}

func ParseKey(key string) (hotkey.Key, error) {
	keyMap := map[string]hotkey.Key{
		"a":          hotkey.KeyA,
		"b":          hotkey.KeyB,
		"c":          hotkey.KeyC,
		"d":          hotkey.KeyD,
		"e":          hotkey.KeyE,
		"f":          hotkey.KeyF,
		"g":          hotkey.KeyG,
		"h":          hotkey.KeyH,
		"i":          hotkey.KeyI,
		"j":          hotkey.KeyJ,
		"k":          hotkey.KeyK,
		"l":          hotkey.KeyL,
		"m":          hotkey.KeyM,
		"n":          hotkey.KeyN,
		"o":          hotkey.KeyO,
		"p":          hotkey.KeyP,
		"q":          hotkey.KeyQ,
		"r":          hotkey.KeyR,
		"s":          hotkey.KeyS,
		"t":          hotkey.KeyT,
		"u":          hotkey.KeyU,
		"v":          hotkey.KeyV,
		"w":          hotkey.KeyW,
		"x":          hotkey.KeyX,
		"y":          hotkey.KeyY,
		"z":          hotkey.KeyZ,
		"0":          hotkey.Key0,
		"1":          hotkey.Key1,
		"2":          hotkey.Key2,
		"3":          hotkey.Key3,
		"4":          hotkey.Key4,
		"5":          hotkey.Key5,
		"6":          hotkey.Key6,
		"7":          hotkey.Key7,
		"8":          hotkey.Key8,
		"9":          hotkey.Key9,
		"up":         hotkey.KeyUp,
		"down":       hotkey.KeyDown,
		"left":       hotkey.KeyLeft,
		"right":      hotkey.KeyRight,
		"delete":     hotkey.KeyDelete,
		"del":        hotkey.KeyDelete,
		"space":      hotkey.KeySpace,
		"enter":      hotkey.KeyReturn,
		"return":     hotkey.KeyReturn,
		"tab":        hotkey.KeyTab,
		"escape":     hotkey.KeyEscape,
		"esc":        hotkey.KeyEscape,
		"pgup":       hotkey.Key(0x21),
		"pgdn":       hotkey.Key(0x22),
		"pgdown":     hotkey.Key(0x22),
		"insert":     hotkey.Key(0x2D),
		"ins":        hotkey.Key(0x2D),
		"home":       hotkey.Key(0x24),
		"end":        hotkey.Key(0x23),
		"scrlk":      hotkey.Key(0x91),
		"scrolllock": hotkey.Key(0x91),
		"pause":      hotkey.Key(0x13),
		"f1":         hotkey.KeyF1,
		"f2":         hotkey.KeyF2,
		"f3":         hotkey.KeyF3,
		"f4":         hotkey.KeyF4,
		"f5":         hotkey.KeyF5,
		"f6":         hotkey.KeyF6,
		"f7":         hotkey.KeyF7,
		"f8":         hotkey.KeyF8,
		"f9":         hotkey.KeyF9,
		"f10":        hotkey.KeyF10,
		"f11":        hotkey.KeyF11,
		"f12":        hotkey.KeyF12,
		"f13":        hotkey.KeyF13,
		"f14":        hotkey.KeyF14,
		"f15":        hotkey.KeyF15,
		"f16":        hotkey.KeyF16,
		"f17":        hotkey.KeyF17,
		"f18":        hotkey.KeyF18,
		"f19":        hotkey.KeyF19,
		"f20":        hotkey.KeyF20,

		//keypad bindings
		"kp_multiply":  hotkey.Key(0x6A),
		"kp_star":      hotkey.Key(0x6A),
		"kp_add":       hotkey.Key(0x6B),
		"kp_plus":      hotkey.Key(0x6B),
		"kp_separator": hotkey.Key(0x6C),
		"kp_comma":     hotkey.Key(0x6C), // amurica
		"kp_subtract":  hotkey.Key(0x6D),
		"kp_minus":     hotkey.Key(0x6D),
		"kp_decimal":   hotkey.Key(0x6E),
		"kp_dot":       hotkey.Key(0x6E), // amurica
		"kp_point":     hotkey.Key(0x6E), // amurica
		"kp_divide":    hotkey.Key(0x6F),
		"kp_slash":     hotkey.Key(0x6F),
		"kp_0":         hotkey.Key(0x60),
		"kp_1":         hotkey.Key(0x61),
		"kp_2":         hotkey.Key(0x62),
		"kp_3":         hotkey.Key(0x63),
		"kp_4":         hotkey.Key(0x64),
		"kp_5":         hotkey.Key(0x65),
		"kp_6":         hotkey.Key(0x66),
		"kp_7":         hotkey.Key(0x67),
		"kp_8":         hotkey.Key(0x68),
		"kp_9":         hotkey.Key(0x69),
	}

	k, found := keyMap[strings.ToLower(key)]
	if !found {
		return 0, fmt.Errorf("invalid key: %s", key)
	}
	return k, nil
}
