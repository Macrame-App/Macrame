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
