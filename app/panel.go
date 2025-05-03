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
	"encoding/base64"
	"encoding/json"
	"macrame/app/helper"
	"macrame/app/structs"
	"net/http"
	"os"
	"strings"
)

func PanelList(w http.ResponseWriter, r *http.Request) {
	panelDirs, err := os.ReadDir("panels")
	if err != nil {
		MCRMLog("PanelList ReadDir Error: ", err)
		json.NewEncoder(w).Encode(false)
		return
	}

	var panelList []interface{}

	for _, panelDir := range panelDirs {
		if panelDir.IsDir() {
			panelList = append(panelList, getPanelInfo(panelDir.Name()))
		}
	}

	if len(panelList) == 0 {
		MCRMLog("PanelList: No panels found")
		json.NewEncoder(w).Encode(false)
		return
	}

	json.NewEncoder(w).Encode(panelList)
}

func getPanelInfo(dirname string) structs.PanelInfo {
	var panelInfo structs.PanelInfo

	jsonFile, err := os.ReadFile("panels/" + dirname + "/panel.json")

	if err != nil {
		panelInfo.Name = strings.Replace(dirname, "_", " ", -1)
		panelInfo.Description = "null"
		panelInfo.AspectRatio = "null"
		panelInfo.Macros = make(map[string]string)
	} else {
		err = json.Unmarshal(jsonFile, &panelInfo)
		if err != nil {
			MCRMLog("getPanelInfo Unmarshal Error: ", err)
		}
	}

	panelInfo.Dir = dirname

	thumb := getPanelThumb(dirname)
	panelInfo.Thumb = thumb

	return panelInfo
}

func getPanelThumb(dirname string) string {
	extensions := []string{".jpg", ".jpeg", ".png", ".webp"}

	for _, ext := range extensions {
		filename := "thumbnail" + ext
		file, err := os.Open("panels/" + dirname + "/" + filename)
		if err != nil {
			MCRMLog("getPanelThumb Open Error: ", err)
			continue
		}
		defer file.Close()

		return encodeImg(file)
	}

	return ""
}

func getPanelCode(dirname string) (html string, css string) {
	htmlBytes, _ := os.ReadFile("panels/" + dirname + "/index.html")
	cssBytes, _ := os.ReadFile("panels/" + dirname + "/output.css")

	return string(htmlBytes), string(cssBytes)
}

func encodeImg(file *os.File) string {
	contents, err := os.ReadFile(file.Name())
	if err != nil {
		MCRMLog("encodeImg ReadFile Error: ", err)
		return ""
	}

	encoded := base64.StdEncoding.EncodeToString(contents)
	return encoded
}

func GetPanel(data string, w http.ResponseWriter, r *http.Request) {
	req := &structs.PanelRequest{}

	_, err := helper.ParseRequest(req, data, r)
	if err != nil {
		MCRMLog("GetPanel ParseRequest Error: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var response = structs.PanelResponse{}

	panelInfo := getPanelInfo(req.Dir)
	panelHtml, panelCss := getPanelCode(req.Dir)

	response.Dir = panelInfo.Dir
	response.Name = panelInfo.Name
	response.Description = panelInfo.Description
	response.AspectRatio = panelInfo.AspectRatio
	response.Macros = panelInfo.Macros
	response.Thumb = panelInfo.Thumb
	response.HTML = panelHtml
	response.Style = panelCss

	json.NewEncoder(w).Encode(response)
}

func SavePanelJSON(w http.ResponseWriter, r *http.Request) {
	var req structs.PanelSaveJSON
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		MCRMLog("SavePanelJSON Decode Error: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filePath := "panels/" + req.Dir + "/panel.json"

	req.Dir = ""

	// Marshal the data to JSON without the dir field
	jsonData, err := json.Marshal(req)
	if err != nil {
		MCRMLog("SavePanelJSON Marshal Error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		MCRMLog("SavePanelJSON WriteFile Error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
