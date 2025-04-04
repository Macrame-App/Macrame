package macro

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

type Step struct {
	Type      string `json:"type"`
	Key       string `json:"key"`
	Code      string `json:"code"`
	Location  int    `json:"location"`
	Direction string `json:"direction"`
	Value     int    `json:"value"`
}

var newMacro struct {
	Name  string `json:"name"`
	Steps []Step `json:"steps"`
}

func Save(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	log.Println(string(body))

	err = json.Unmarshal(body, &newMacro)

	if err != nil {
		panic(err)
	}

	stepsJSON, err := json.Marshal(newMacro.Steps)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("../macros/"+makeValidFilename(newMacro.Name)+".json", stepsJSON, 0644)
	if err != nil {
		panic(err)
	}
}

func makeValidFilename(s string) string {
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

func List(w http.ResponseWriter, r *http.Request) {
	log.Println("listing macros")
	dir := "../macros"
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var fileNames []string

	for _, file := range files {
		filename := filepath.Base(file.Name())
		filename = strings.TrimSuffix(filename, filepath.Ext(filename))
		filename = strings.Replace(filename, "_", " ", -1)

		fileNames = append(fileNames, filename)
	}

	json.NewEncoder(w).Encode(fileNames)
}

func Delete(w http.ResponseWriter, r *http.Request) {}

func Play(w http.ResponseWriter, r *http.Request) {
	type MacroRequest struct {
		Macro string `json:"macro"`
	}

	var req MacroRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	macro := req.Macro

	macroFile, err := readMacroFile(fmt.Sprintf("../macros/%s.json", makeValidFilename(macro)))

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	playMacro(macroFile)
	// fmt.Println(macroFile)
}

func readMacroFile(filename string) (steps []Step, err error) {

	log.Println(filename)
	// Let's first read the `config.json` file
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `steps`
	err = json.Unmarshal(content, &steps)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return steps, nil
}

func playMacro(steps []Step) {
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
