/*
Macrame is a program that enables the user to create keyboard macros and button panels.
The macros are saved as simple JSON files and can be linked to the button panels. The panels can
be created with HTML and CSS.

Copyright (C) 2025 Jesse Malotaux

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

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
	if name == "\\" {
		return "Backslash"
	}

	for key, value := range translations {
		if value == name {
			return key
		}
	}
	return name
}
