//go:build linux

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
		return hotkey.Mod1, nil
	case "win":
		return hotkey.Mod4, nil
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
		"pgup":       hotkey.Key(0xff55),
		"pageup":     hotkey.Key(0xff55),
		"pgdn":       hotkey.Key(0xff56),
		"pagedown":   hotkey.Key(0xff56),
		"pgdown":     hotkey.Key(0xff56),
		"insert":     hotkey.Key(0xff63),
		"ins":        hotkey.Key(0xff63),
		"home":       hotkey.Key(0xff50),
		"end":        hotkey.Key(0xff57),
		"scrlk":      hotkey.Key(0xff14),
		"scrolllock": hotkey.Key(0xff14),
		"pause":      hotkey.Key(0xff13),
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
		"kp_space":     hotkey.Key(0xff80),
		"kp_tab":       hotkey.Key(0xff89),
		"kp_enter":     hotkey.Key(0xff8d),
		"kp_f1":        hotkey.Key(0xff91),
		"kp_f2":        hotkey.Key(0xff92),
		"kp_f3":        hotkey.Key(0xff93),
		"kp_f4":        hotkey.Key(0xff94),
		"kp_home":      hotkey.Key(0xff95),
		"kp_left":      hotkey.Key(0xff96),
		"kp_up":        hotkey.Key(0xff97),
		"kp_right":     hotkey.Key(0xff98),
		"kp_down":      hotkey.Key(0xff99),
		"kp_prior":     hotkey.Key(0xff9a),
		"kp_pageup":    hotkey.Key(0xff9a),
		"kp_next":      hotkey.Key(0xff9b),
		"kp_pagedown":  hotkey.Key(0xff9b),
		"kp_end":       hotkey.Key(0xff9c),
		"kp_begin":     hotkey.Key(0xff9d),
		"kp_insert":    hotkey.Key(0xff9e),
		"kp_ins":       hotkey.Key(0xff9e),
		"kp_delete":    hotkey.Key(0xff9f),
		"kp_del":       hotkey.Key(0xff9f),
		"kp_equal":     hotkey.Key(0xffbd),
		"kp_multiply":  hotkey.Key(0xffaa),
		"kp_star":      hotkey.Key(0xffaa),
		"kp_add":       hotkey.Key(0xffab),
		"kp_plus":      hotkey.Key(0xffab),
		"kp_separator": hotkey.Key(0xffac),
		"kp_comma":     hotkey.Key(0xffac), // amurica
		"kp_subtract":  hotkey.Key(0xffad),
		"kp_minus":     hotkey.Key(0xffad),
		"kp_decimal":   hotkey.Key(0xffae),
		"kp_dot":       hotkey.Key(0xffae), // amurica
		"kp_point":     hotkey.Key(0xffae), // amurica
		"kp_divide":    hotkey.Key(0xffaf),
		"kp_slash":     hotkey.Key(0xffaf),
		"kp_0":         hotkey.Key(0xffb0),
		"kp_1":         hotkey.Key(0xffb1),
		"kp_2":         hotkey.Key(0xffb2),
		"kp_3":         hotkey.Key(0xffb3),
		"kp_4":         hotkey.Key(0xffb4),
		"kp_5":         hotkey.Key(0xffb5),
		"kp_6":         hotkey.Key(0xffb6),
		"kp_7":         hotkey.Key(0xffb7),
		"kp_8":         hotkey.Key(0xffb8),
		"kp_9":         hotkey.Key(0xffb9),
	}

	k, found := keyMap[strings.ToLower(key)]
	if !found {
		return 0, fmt.Errorf("invalid key: %s", key)
	}
	return k, nil
}
