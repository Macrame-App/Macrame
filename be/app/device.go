package app

import (
	"be/app/helper"
	"be/app/structs"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func GetServerIP(w http.ResponseWriter, r *http.Request) {
	ifs, err := net.Interfaces()
	if err != nil {
		log.Println(err)
	}
	for _, ifi := range ifs {
		addrs, err := ifi.Addrs()
		if err != nil {
			log.Println(err)
		}
		for _, addr := range addrs {
			var ip net.IP

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip != nil && ip.To4() != nil {
				json.NewEncoder(w).Encode(ip.String())
				return
			}
		}
	}
}

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

		// log.Println(device, ext)

		if _, ok := devices[device]; !ok {
			devices[device] = make(map[string]interface{})
		}

		if ext == ".json" {
			devices[device]["settings"] = readDeviceSettings(filePath)
		}
		if ext == ".key" {
			devices[device]["key"] = true
		}
	}

	result := map[string]interface{}{
		"devices": devices,
	}

	log.Println(result)

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

func DeviceAccessCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("device access check")
	var req structs.Check

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, errSett := os.Stat("devices/" + req.Uuid + ".json")
	_, errKey := os.Stat("devices/" + req.Uuid + ".key")

	if (errSett == nil) && (errKey == nil) {
		log.Println("authorized")
		json.NewEncoder(w).Encode("authorized")
	} else if (errSett == nil) && (errKey != nil) {
		log.Println("unauthorized")
		json.NewEncoder(w).Encode("unauthorized")
	} else {
		log.Println("unauthorized")
		json.NewEncoder(w).Encode("unlinked")
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

	json.NewEncoder(w).Encode("unauthorized")
}

func PingLink(w http.ResponseWriter, r *http.Request) {
	log.Println("ping link")
	var req structs.Check
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		json.NewEncoder(w).Encode(false)
		return
	}

	key, keyErr := os.ReadFile("devices/" + req.Uuid + ".key")
	pin, pinErr := os.ReadFile("devices/" + req.Uuid + ".tmp")

	encryptedKey, encErr := helper.EncryptAES(string(pin), string(key))

	log.Println(encryptedKey, string(pin), string(key))

	if keyErr == nil && pinErr == nil && encErr == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(encryptedKey))
		return
	}

	json.NewEncoder(w).Encode(false)
}

func StartLink(w http.ResponseWriter, r *http.Request) {
	log.Println("start link")
	var req structs.Check

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pin := fmt.Sprintf("%04d", rand.Intn(10000))

	deviceKey := helper.GenerateKey()

	err = helper.SaveDeviceKey(req.Uuid, deviceKey)

	if err == nil && helper.TempPinFile(req.Uuid, pin) {
		json.NewEncoder(w).Encode(pin)
	}

	return
}

func PollLink(w http.ResponseWriter, r *http.Request) {
	var req structs.Check

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		json.NewEncoder(w).Encode(false)
		return
	}

	if helper.CheckPinFile(req.Uuid) {
		json.NewEncoder(w).Encode(true)
		return
	}

	json.NewEncoder(w).Encode(false)
}

func RemoveLink(data string, w http.ResponseWriter, r *http.Request) {
	req := &structs.Check{}
	_, err := helper.ParseRequest(req, data, r)

	if err != nil {
		json.NewEncoder(w).Encode(false)
		return
	}

	err = os.Remove("devices/" + req.Uuid + ".key")

	if err != nil {
		json.NewEncoder(w).Encode(false)
		return
	}

	json.NewEncoder(w).Encode(true)
}

func Handshake(w http.ResponseWriter, r *http.Request) {
	var req structs.Handshake

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return
	}

	deviceKey, err := helper.GetKeyByUuid(req.Uuid)

	if err != nil {
		json.NewEncoder(w).Encode(false)
		return
	}

	decryptShake, _ := helper.DecryptAES(deviceKey, req.Shake)

	if decryptShake == getDateStr() {
		os.Remove("devices/" + req.Uuid + ".tmp")

		json.NewEncoder(w).Encode(true)
		return
	} else {
		os.Remove("devices/" + req.Uuid + ".key")
	}

	json.NewEncoder(w).Encode(false)
}

func getDateStr() string {
	date := time.Now()
	year, month, day := date.Date()
	formattedDate := fmt.Sprintf("%04d%02d%02d", year, month, day)
	return formattedDate
}
