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

	"be/app/helper"
	"be/app/structs"
)

func SaveMacro(w http.ResponseWriter, r *http.Request) {
	var newMacro structs.NewMacro

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

	err = os.WriteFile("../macros/"+helper.FormatMacroFileName(newMacro.Name)+".json", stepsJSON, 0644)
	if err != nil {
		panic(err)
	}
}

func ListMacros(w http.ResponseWriter, r *http.Request) {
	log.Println("listing macros")
	dir := "../macros"
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Println(err)
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

func DeleteMacro(w http.ResponseWriter, r *http.Request) {}

func PlayMacro(w http.ResponseWriter, r *http.Request) {
	var req structs.MacroRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	macro := req.Macro

	var filename = helper.FormatMacroFileName(macro)
	var filepath = fmt.Sprintf("../macros/%s.json", filename)

	macroFile, err := helper.ReadMacroFile(filepath)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.RunMacroSteps(macroFile)
}
