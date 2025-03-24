package app

import (
	"be/app/helper"
	"be/app/structs"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
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

var (
	mu    sync.Mutex
	queue = make(map[string][]structs.RemoteWebhook)
)

func PollRemote(w http.ResponseWriter, r *http.Request) {

}

func DeviceAccessCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("device access check")
	var req structs.Check

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, errSett := os.Stat("devices/" + req.Uuid + ".json")
	_, errKey := os.Stat("devices/" + req.Uuid + ".pem")

	log.Println(errSett, errKey)

	if (errSett == nil) && (errKey == nil) {
		log.Println("authorized")
		json.NewEncoder(w).Encode("authorized")
	} else if (errSett == nil) && (errKey != nil) {
		log.Println("requested")
		json.NewEncoder(w).Encode("requested")
	} else {
		log.Println("unauthorized")
		json.NewEncoder(w).Encode("unauthorized")
	}

	return
}

func DeviceAccessRequest(w http.ResponseWriter, r *http.Request) {
	var req structs.Request

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deviceSettings := structs.Settings{Name: req.Name, Type: req.Type}

	settingsJSON, err := json.Marshal(deviceSettings)
	if err != nil {
		log.Println(err)
		return
	}

	err = os.WriteFile("devices/"+req.Uuid+".json", settingsJSON, 0644)

	if err != nil {
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(true)
}

func PingLink(w http.ResponseWriter, r *http.Request) {
	log.Println("ping link")
	var req structs.Check
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		json.NewEncoder(w).Encode(false)
		return
	}

	var filename = "devices/" + req.Uuid + ".tmp"

	_, err = os.ReadFile(filename)

	if err == nil {
		json.NewEncoder(w).Encode(true)
		return
	}

	json.NewEncoder(w).Encode(false)
	return
}

func StartLink(w http.ResponseWriter, r *http.Request) {
	var req structs.Check

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pin := fmt.Sprintf("%04d", rand.Intn(10000))

	if helper.TempPinFile(req.Uuid, pin) {
		json.NewEncoder(w).Encode(pin)
	}
}

func Handshake(w http.ResponseWriter, r *http.Request) {
	var req structs.Handshake

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println(req.Shake)

}

func DeviceAuth(w http.ResponseWriter, r *http.Request) bool {
	return true
}
