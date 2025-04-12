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
	re := regexp.MustCompile(`[\/\?\*\>\<\:\\"\|\n]`)
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
