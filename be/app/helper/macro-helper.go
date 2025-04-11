package helper

import (
	"encoding/json"
	"errors"
	"os"
	"regexp"
	"strings"
	"time"

	"be/app/structs"

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

func ReadMacroFile(filename string) (steps []structs.Step, err error) {
	content, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &steps)

	return steps, err
}

func RunMacroSteps(steps []structs.Step) error {
	for _, step := range steps {
		switch step.Type {
		case "key":
			err := robotgo.KeyToggle(step.Key, step.Direction)
			if err != nil {
				return errors.New("RunMacroSteps KeyToggle Error: " + err.Error())
			}
		case "delay":
			time.Sleep(time.Duration(step.Location) * time.Millisecond)
		default:
			return errors.New("RunMacroSteps Unknown step type:" + step.Type)
		}
	}
	return nil
}
