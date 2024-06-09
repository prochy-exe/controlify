# Controlify

An app that bridges the gap between you and easy Spotify playback control.

## Installation

To use this app, you need to have [Spicetify](https://spicetify.app/) and [controlify-plugin](https://github.com/prochy-exe/controlify-plugin) installed.

As for controlify itself, just download the zip file, extract it to any directory of your choosing and just run either the tray app or the cli app. You can also run the tray app with the --cli tag, that will run the app in the background with no tray icon and notifications.

# Features
- The entirety of toggle-able buttons in Spotify is mapped, to set these shortcuts just modify the config.json (WAYLAND IS NOT SUPPORTED!!!)
- Deej support, so you can directly change the volume of Spotify without affecting other audio sessions (that means you can control the volume of any Spotify Connect device!!!)

## Usage

I personally use this app in conjunction with deej. But the possibilities are endless, really, you can make an Arduino to simulate function keys or plug in a keypad or whatever you want really. Most of the keyboard is mappable (function keys up to 20, and keypad numbers are mapped too).

## Contributing

Pull requests are always welcome, so if you have any feature suggestions or fixes please feel free to submit them!

## Credits

- Icon was created from icons available at [Google Fonts](https://fonts.google.com/icons)
- External Go libraries used:
  - [systray used on Windows](https://github.com/getlantern/systray)
  - [systray used on Linux](https://github.com/fyne-io/systray)
  - [Beeep](https://github.com/gen2brain/beeep)
  - [Hotkey](https://pkg.go.dev/golang.design/x/hotkey)
  - [Websocket](https://github.com/gorilla/websocket)
