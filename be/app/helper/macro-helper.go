package helper

import (
	"encoding/json"
	"log"
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
	log.Println(filename)

	content, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	err = json.Unmarshal(content, &steps)

	return steps, err
}

func RunMacroSteps(steps []structs.Step) {
	for _, step := range steps {
		// log.Println(step)
		switch step.Type {
		case "key":
			robotgo.KeyToggle(step.Key, step.Direction)
			// log.Println("Toggling", step.Key, "to", step.Direction)
		case "delay":
			time.Sleep(time.Duration(step.Location) * time.Millisecond)
			// log.Println("Sleeping for", step.Value, "milliseconds")
		default:
			log.Println("Unknown step type:", step.Type)
		}
	}
}
