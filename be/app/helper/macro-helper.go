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
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

func FormatMacroFileName(s string) string {
	// Remove invalid characters
	re := regexp.MustCompile(`[\/\?\*\>\<\:\"\|
]`)
	s = re.ReplaceAllString(s, "")

	// Replace spaces with underscores
	s = strings.ReplaceAll(s, " ", "_")

	// Remove special characters
	re = regexp.MustCompile(`[!@#$%^&\(\)\[\]\{\}\~]`)
	s = re.ReplaceAllString(s, "")

	// Truncate the string
	if len(s) > 255 {
		s = s[:255]
	}

	return s
}

func ReadMacroFile(filename string) (steps []map[string]interface{}, err error) {
	content, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &steps)

	return steps, err
}

func RunMacroSteps(steps []map[string]interface{}) error {
	for _, step := range steps {
		switch step["type"] {
		case "key":
			keyCode := step["code"].(string)

			if strings.Contains(keyCode, "|") {
				keyCode = handleToggleCode(keyCode, step["direction"].(string))
			}

			robotgo.KeyToggle(keyCode, step["direction"].(string))
		case "delay":

			time.Sleep(time.Duration(step["value"].(float64)) * time.Millisecond)

		default:
			return errors.New("RunMacroSteps Unknown step type: %v" + fmt.Sprint(step["type"]))
		}
	}
	return nil
}

var toggleCodes = map[string]bool{}

func handleToggleCode(keyCode string, direction string) string {
	splitCodes := strings.Split(keyCode, "|")

	if direction == "down" {
		if _, ok := toggleCodes[splitCodes[0]]; !ok {
			toggleCodes[splitCodes[0]] = true
			return splitCodes[0]
		}
		return splitCodes[1]
	}

	if direction == "up" {
		if toggleCodes[splitCodes[0]] {
			toggleCodes[splitCodes[0]] = false
			return splitCodes[0]
		}
		delete(toggleCodes, splitCodes[0])
		return splitCodes[1]
	}

	return ""
}
