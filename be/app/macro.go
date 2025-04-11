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

	stepsJSON, err := json.Marshal(newMacro.Steps)
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

func ListMacros(w http.ResponseWriter, r *http.Request) {
	dir := "../macros"
	files, err := os.ReadDir(dir)
	if err != nil {
		MCRMLog("ListMacros ReadDir Error: ", err)
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
