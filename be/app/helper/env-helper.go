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
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

var configPath = "../public/config.js"

func EnvGet(key string) string {
	if !configFileExists() {
		CreateConfigFile(configPath)
		CheckFeDevDir()
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Println("Error reading config.js:", err)
		return ""
	}

	raw := strings.TrimSpace(string(data))
	raw = strings.TrimPrefix(raw, "window.__CONFIG__ = ")
	raw = strings.TrimSuffix(raw, ";")

	var config map[string]string
	if err := json.Unmarshal([]byte(raw), &config); err != nil {
		log.Println("Error parsing config.js:", err)
		return ""
	}

	return config[key]
}

func configFileExists() bool {
	_, err := os.Stat(configPath)
	return err == nil
}

func CheckFeDevDir() {
	log.Println("Checking FE dev directory...")
	_, err := os.Stat("../fe")

	if err != nil {
		return
	}

	copyConfigToFe()
}

func copyConfigToFe() {
	data, err := os.ReadFile("../public/config.js")

	if err != nil {
		log.Println("Error reading config.js:", err)
		return
	}

	if err := os.WriteFile("../fe/config.js", data, 0644); err != nil {
		log.Println("Error writing config.js:", err)
	}
}

func CreateConfigFile(filename string) {
	port, _ := findOpenPort()
	saltKey := GenerateKey()
	salt := saltKey[:28]
	iv := GenerateRandomIntegerString(16)

	config := map[string]string{
		"MCRM__PORT": port,
		"MCRM__SALT": salt,
		"MCRM__IV":   iv,
	}

	jsonData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Println("Error creating config JSON:", err)
		return
	}
	jsData := "window.__CONFIG__ = " + string(jsonData) + ";"

	log.Println("Created JS config:", jsData)

	if err := os.WriteFile(filename, []byte(jsData), 0644); err != nil {
		log.Println("Error writing config.js:", err)
	}
}

func findOpenPort() (string, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return "", err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return "", err
	}
	defer l.Close()
	return strconv.Itoa(l.Addr().(*net.TCPAddr).Port), nil
}
