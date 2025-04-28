package helper

import (
	"strings"
)

var translations = map[string]string{
	"ArrowUp":            "up",
	"ArrowDown":          "down",
	"ArrowRight":         "right",
	"ArrowLeft":          "left",
	"Meta":               "cmd",
	"MetaLeft":           "lcmd",
	"MetaRight":          "rcmd",
	"Alt":                "alt",
	"AltLeft":            "lalt",
	"AltRight":           "ralt",
	"Control":            "ctrl",
	"ControlLeft":        "lctrl",
	"ControlRight":       "rctrl",
	"Shift":              "shift",
	"ShiftLeft":          "lshift",
	"ShiftRight":         "rshift",
	"AudioVolumeMute":    "audio_mute",
	"AudioVolumeDown":    "audio_vol_down",
	"AudioVolumeUp":      "audio_vol_up",
	"MediaTrackPrevious": "audio_prev",
	"MediaTrackNext":     "audio_next",
	"MediaPlayPause":     "audio_play|audio_pause",
	"Numpad0":            "num0",
	"Numpad1":            "num1",
	"Numpad2":            "num2",
	"Numpad3":            "num3",
	"Numpad4":            "num4",
	"Numpad5":            "num5",
	"Numpad6":            "num6",
	"Numpad7":            "num7",
	"Numpad8":            "num8",
	"Numpad9":            "num9",
	"NumLock":            "num_lock",
	"NumpadDecimal":      "num.",
	"NumpadAdd":          "num+",
	"NumpadSubtract":     "num-",
	"NumpadMultiply":     "num*",
	"NumpadDivide":       "num/",
	"NumpadEnter":        "num_enter",
	"Clear":              "num_clear",
	"BracketLeft":        "[",
	"BracketRight":       "]",
	"Quote":              "'",
	"Semicolon":          ";",
	"Backquote":          "`",
	"Backslash":          "\\",
	"IntlBackslash":      "\\",
	"Slash":              "/",
	"Comma":              ",",
	"Period":             ".",
	"Equal":              "=",
	"Minus":              "-",
}

func Translate(code string) string {
	if val, ok := translations[code]; ok {
		return val
	}
	return strings.ToLower(code)
}

func ReverseTranslate(name string) string {
	for key, value := range translations {
		if value == name {
			return key
		}
	}
	return name
}

// func Translate(code string) string {
// 	translations := map[string]string{
// 		"ArrowUp":            "up",
// 		"ArrowDown":          "down",
// 		"ArrowRight":         "right",
// 		"ArrowLeft":          "left",
// 		"Meta":               "cmd",
// 		"MetaLeft":           "lcmd",
// 		"MetaRight":          "rcmd",
// 		"Alt":                "alt",
// 		"AltLeft":            "lalt",
// 		"AltRight":           "ralt",
// 		"Control":            "ctrl",
// 		"ControlLeft":        "lctrl",
// 		"ControlRight":       "rctrl",
// 		"Shift":              "shift",
// 		"ShiftLeft":          "lshift",
// 		"ShiftRight":         "rshift",
// 		"AudioVolumeMute":    "audio_mute",
// 		"AudioVolumeDown":    "audio_vol_down",
// 		"AudioVolumeUp":      "audio_vol_up",
// 		"MediaTrackPrevious": "audio_prev",
// 		"MediaTrackNext":     "audio_next",
// 		"MediaPlayPause":     "audio_play|audio_pause",
// 		"Numpad0":            "num0",
// 		"Numpad1":            "num1",
// 		"Numpad2":            "num2",
// 		"Numpad3":            "num3",
// 		"Numpad4":            "num4",
// 		"Numpad5":            "num5",
// 		"Numpad6":            "num6",
// 		"Numpad7":            "num7",
// 		"Numpad8":            "num8",
// 		"Numpad9":            "num9",
// 		"NumLock":            "num_lock",
// 		"NumpadDecimal":      "num.",
// 		"NumpadAdd":          "num+",
// 		"NumpadSubtract":     "num-",
// 		"NumpadMultiply":     "num*",
// 		"NumpadDivide":       "num/",
// 		"NumpadEnter":        "num_enter",
// 		"Clear":              "num_clear",
// 		"BracketLeft":        "[",
// 		"BracketRight":       "]",
// 		"Quote":              "'",
// 		"Semicolon":          ";",
// 		"Backquote":          "`",
// 		"Backslash":          "\\",
// 		"IntlBackslash":      "\\",
// 		"Slash":              "/",
// 		"Comma":              ",",
// 		"Period":             ".",
// 		"Equal":              "=",
// 		"Minus":              "-",
// 	}

// 	if translations[code] == "" {
// 		return strings.ToLower(code)
// 	}

// 	return translations[code]
// }

// Redundant translation because tolower can be used
// "Backspace":          "backspace",
// "Delete":             "delete",
// "Enter":              "enter",
// "Tab":                "tab",
// "Escape":             "esc",
// "Home":               "home",
// "End":                "end",
// "PageUp":             "pageup",
// "PageDown":           "pagedown",
// "F1":                 "f1",
// "F2":                 "f2",
// "F3":                 "f3",
// "F4":                 "f4",
// "F5":                 "f5",
// "F6":                 "f6",
// "F7":                 "f7",
// "F8":                 "f8",
// "F9":                 "f9",
// "F10":                "f10",
// "F11":                "f11",
// "F12":                "f12",
// "F13":                "f13",
// "F14":                "f14",
// "F15":                "f15",
// "F16":                "f16",
// "F17":                "f17",
// "F18":                "f18",
// "F19":                "f19",
// "F20":                "f20",
// "F21":                "f21",
// "F22":                "f22",
// "F23":                "f23",
// "F24":                "f24",
// "CapsLock":           "capslock",
// "Space":              "space",
// "PrintScreen":        "printscreen",
// "Insert":             "insert",
