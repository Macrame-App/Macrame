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
	"be/app/helper"
	"be/app/structs"
	"encoding/json"
	"fmt"
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
		MCRMLog(err)
		return
	}

	for _, ifi := range ifs {
		// Skip interfaces that are down
		if ifi.Flags&net.FlagUp == 0 {
			continue
		}
		// Skip loopback interfaces
		if ifi.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := ifi.Addrs()
		if err != nil {
			MCRMLog(err)
			continue
		}

		for _, addr := range addrs {
			var ip net.IP

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.To4() == nil {
				continue
			}

			// Skip APIPA (169.254.x.x) addresses
			if ip.IsLinkLocalUnicast() {
				continue
			}

			// Found a good IP, return it
			json.NewEncoder(w).Encode(ip.String())
			return
		}
	}
}

func DeviceList(w http.ResponseWriter, r *http.Request) {
	dir := "devices"

	files, err := os.ReadDir(dir)
	if err != nil {
		os.MkdirAll(dir, 0600)
		files = nil
		MCRMLog("DeviceList Error: ", err)
	}

	devices := make(map[string]map[string]interface{})

	for _, file := range files {
		filePath := dir + "/" + file.Name()
		ext := filepath.Ext(filePath)
		device := strings.TrimSuffix(file.Name(), ext)

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

	json.NewEncoder(w).Encode(result)
}

func readDeviceSettings(filepath string) (settings structs.Settings) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		MCRMLog("readDeviceSettings Error: ", err)
	}

	err = json.Unmarshal(data, &settings)
	if err != nil {
		MCRMLog("readDeviceSettings JSON Error: ", err)
	}

	return settings
}

func DeviceAccessCheck(w http.ResponseWriter, r *http.Request) {
	var req structs.Check

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		MCRMLog("DeviceAccessCheck Error: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, errSett := os.Stat("devices/" + req.Uuid + ".json")
	_, errKey := os.Stat("devices/" + req.Uuid + ".key")

	if (errSett == nil) && (errKey == nil) {
		json.NewEncoder(w).Encode("authorized")
	} else if (errSett == nil) && (errKey != nil) {
		MCRMLog("DeviceAccessCheck: UUID: ", req.Uuid, "; Access: Unauthorized")
		json.NewEncoder(w).Encode("unauthorized")
	} else {
		MCRMLog("DeviceAccessCheck: UUID: ", req.Uuid, "; Access: Unlinked")
		json.NewEncoder(w).Encode("unlinked")
	}
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
		MCRMLog("DeviceAccessRequest JSON Error: ", err)
		return
	}

	err = os.WriteFile("devices/"+req.Uuid+".json", settingsJSON, 0644)

	if err != nil {
		MCRMLog("DeviceAccessRequest Error: ", err)
		return
	}

	json.NewEncoder(w).Encode("unauthorized")
}

func PingLink(w http.ResponseWriter, r *http.Request) {
	var req structs.Check
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		MCRMLog("PingLink Error: ", err)
		json.NewEncoder(w).Encode(false)
		return
	}

	key, keyErr := os.ReadFile("devices/" + req.Uuid + ".key")
	pin, pinErr := os.ReadFile("devices/" + req.Uuid + ".tmp")

	encryptedKey, encErr := helper.EncryptAES(string(pin), string(key))

	if keyErr == nil && pinErr == nil && encErr == nil {
		MCRMLog("PINGLINK FIXED")
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(encryptedKey))
		return
	} else {
		MCRMLog("PingLink Error: keyErr:", keyErr, "; pinErr:", pinErr, "; encErr:", encErr)
	}

	json.NewEncoder(w).Encode(false)
}

func StartLink(w http.ResponseWriter, r *http.Request) {
	var req structs.Check

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		MCRMLog("StartLink Error: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pin := fmt.Sprintf("%04d", rand.Intn(10000))

	deviceKey := helper.GenerateKey()

	errKey := helper.SaveDeviceKey(req.Uuid, deviceKey)
	savedPin, errPin := helper.TempPinFile(req.Uuid, pin)

	if errKey == nil && errPin == nil && savedPin {
		json.NewEncoder(w).Encode(pin)
	} else {
		MCRMLog("StartLink Error: errKey:", errKey, "; errPin:", errPin)
	}
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
		MCRMLog("RemoveLink ParseRequest Error: ", err)
		json.NewEncoder(w).Encode(false)
		return
	}

	err = os.Remove("devices/" + req.Uuid + ".key")

	if err != nil {
		MCRMLog("RemoveLink Remove Error: ", err)
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
		MCRMLog("Handshake GetKeyByUuid Error: ", err)
		json.NewEncoder(w).Encode(false)
		return
	}

	decryptShake, _ := helper.DecryptAES(deviceKey, req.Shake)

	helper.RemovePinFile(req.Uuid)

	if decryptShake == getDateStr() {
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
