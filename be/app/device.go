package app

import (
	"be/app/structs"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func DeviceList(w http.ResponseWriter, r *http.Request) {
	log.Println("device list")
	dir := "devices"
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	devices := make(map[string]map[string]interface{})

	for _, file := range files {
		filePath := dir + "/" + file.Name()
		ext := filepath.Ext(filePath)
		device := strings.TrimSuffix(file.Name(), ext)

		log.Println(device, ext)

		if _, ok := devices[device]; !ok {
			devices[device] = make(map[string]interface{})
		}

		if ext == ".json" {
			devices[device]["settings"] = readDeviceSettings(filePath)
		}
		if ext == ".pem" {
			devices[device]["key"] = true
		}
	}

	result := map[string]interface{}{
		"devices": devices,
	}

	json.NewEncoder(w).Encode(result)
}

func readDeviceSettings(filepath string) (settings structs.Settings) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(data, &settings)
	if err != nil {
		log.Println(err)
	}
	log.Println(settings)
	return settings
}

func DeviceAccess(w http.ResponseWriter, r *http.Request) bool {
	return true
}

func DeviceAuth(w http.ResponseWriter, r *http.Request) bool {
	return true
}
