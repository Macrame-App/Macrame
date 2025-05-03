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

package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"macrame/app/helper"
	"macrame/app/structs"
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

	macroFile, err := helper.ReadMacroFile(fmt.Sprintf("macros/%s.json", filename))

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

	err = os.WriteFile("macros/"+helper.FormatMacroFileName(newMacro.Name)+".json", stepsJSON, 0644)
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
	dir := "macros"
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

func DeleteMacro(w http.ResponseWriter, r *http.Request) {
	var req structs.MacroRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		MCRMLog("DeleteMacro Decode Error: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var filename = helper.FormatMacroFileName(req.Macro)

	err = os.Remove("macros/" + filename + ".json")

	if err != nil {
		MCRMLog("DeleteMacro Remove Error: ", err)
		json.NewEncoder(w).Encode(false)
		return
	}
	log.Println("Deleted Macro:", req.Macro)
	json.NewEncoder(w).Encode(true)
}

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
	var filepath = fmt.Sprintf("macros/%s.json", filename)

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

	macroFile, err := helper.ReadMacroFile(fmt.Sprintf("macros/%s.json", req.Macro))

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
