package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"be/app/helper"
	"be/app/structs"
)

func CheckMacro(w http.ResponseWriter, r *http.Request) {
	var req structs.MacroRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		MCRMLog("OpenMacro Decode Error: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var filename = helper.FormatMacroFileName(req.Macro)

	macroFile, err := helper.ReadMacroFile(fmt.Sprintf("../macros/%s.json", filename))

	if macroFile != nil && err == nil {
		json.NewEncoder(w).Encode(true)
		return
	} else {
		MCRMLog("OpenMacro ReadMacroFile Error: ", err)
		json.NewEncoder(w).Encode(false)
		return
	}
}

func SaveMacro(w http.ResponseWriter, r *http.Request) {
	var newMacro structs.NewMacro

	body, err := io.ReadAll(r.Body)
	if err != nil {
		MCRMLog("SaveMacro ReadAll Error: ", err)
		return
	}

	err = json.Unmarshal(body, &newMacro)
	if err != nil {
		MCRMLog("SaveMacro Unmarshal Error: ", err)
		return
	}

	simplifiedSteps := make([]map[string]interface{}, 0)
	for _, step := range newMacro.Steps {
		simplifiedSteps = append(simplifiedSteps, simplifyMacro(step))
	}

	stepsJSON, err := json.Marshal(simplifiedSteps)
	if err != nil {
		MCRMLog("SaveMacro Marshal Error: ", err)
		return
	}

	err = os.WriteFile("../macros/"+helper.FormatMacroFileName(newMacro.Name)+".json", stepsJSON, 0644)
	if err != nil {
		MCRMLog("SaveMacro WriteFile Error: ", err)
		return
	}
}

func simplifyMacro(step structs.Step) map[string]interface{} {
	simplified := make(map[string]interface{})

	simplified["type"] = step.Type

	switch step.Type {
	case "delay":
		simplified["value"] = step.Value
	case "key":
		keyCode := step.Code

		if keyCode == "" || (strings.Contains(keyCode, "Digit")) {
			keyCode = step.Key
		} else if strings.Contains(keyCode, "Key") {
			keyCode = strings.Replace(keyCode, "Key", "", 1)
		}

		simplified["code"] = helper.Translate(keyCode)
		simplified["direction"] = step.Direction
	}

	return simplified
}

func ListMacros(w http.ResponseWriter, r *http.Request) {
	dir := "../macros"
	files, err := os.ReadDir(dir)
	if err != nil {
		MCRMLog("ListMacros ReadDir Error: ", err)
		json.NewEncoder(w).Encode(false)
		return
	}

	var macroList []structs.MacroInfo

	for _, file := range files {
		filename := filepath.Base(file.Name())
		macroname := strings.TrimSuffix(filename, filepath.Ext(filename))
		nicename := strings.Replace(macroname, "_", " ", -1)

		macroList = append(macroList, structs.MacroInfo{
			Name:      nicename,
			Macroname: macroname,
		})
	}

	json.NewEncoder(w).Encode(macroList)
}

func DeleteMacro(w http.ResponseWriter, r *http.Request) {}

func PlayMacro(data string, w http.ResponseWriter, r *http.Request) {
	req := &structs.MacroRequest{}
	_, err := helper.ParseRequest(req, data, r)

	if err != nil {
		MCRMLog("PlayMacro ParseRequest Error: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	macro := req.Macro

	MCRMLog("Playing Macro: ", macro)

	var filename = helper.FormatMacroFileName(macro)
	var filepath = fmt.Sprintf("../macros/%s.json", filename)

	macroFile, err := helper.ReadMacroFile(filepath)
	if err != nil {
		MCRMLog("PlayMacro ReadMacroFile Error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = helper.RunMacroSteps(macroFile)

	if err != nil {
		MCRMLog("PlayMacro RunMacroSteps Error: ", err)
		return
	}
}

func OpenMacro(w http.ResponseWriter, r *http.Request) {
	var req structs.MacroRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		MCRMLog("OpenMacro Decode Error: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	macroFile, err := helper.ReadMacroFile(fmt.Sprintf("../macros/%s.json", req.Macro))

	if err != nil {
		MCRMLog("OpenMacro ReadMacroFile Error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Walk through the macro file and reverse translate codes
	for i, action := range macroFile {
		if actionType, ok := action["type"].(string); ok && actionType == "key" {
			if code, ok := action["code"].(string); ok {
				macroFile[i]["code"] = helper.ReverseTranslate(code)
			}
		}
	}

	json.NewEncoder(w).Encode(macroFile)
}
