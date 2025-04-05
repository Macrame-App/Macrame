package app

import (
	"be/app/structs"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func PanelList(w http.ResponseWriter, r *http.Request) {
	panelDirs, err := os.ReadDir("../panels")
	if err != nil {
		json.NewEncoder(w).Encode(false)
		return
	}

	var panelList []interface{}

	for _, panelDir := range panelDirs {
		if panelDir.IsDir() {
			// log.Println(getPanelInfo(panelDir.Name()))
			panelList = append(panelList, getPanelInfo(panelDir.Name()))
		}
	}

	if len(panelList) == 0 {
		json.NewEncoder(w).Encode(false)
		return
	}

	json.NewEncoder(w).Encode(panelList)
}

func getPanelInfo(dirname string) structs.PanelInfo {
	var panelInfo structs.PanelInfo

	jsonFile, err := os.ReadFile("../panels/" + dirname + "/panel.json")

	if err != nil {
		panelInfo.Name = strings.Replace(dirname, "_", " ", -1)
		panelInfo.Description = "null"
		panelInfo.AspectRatio = "null"
		panelInfo.Macros = make(map[string]string)
	} else {
		err = json.Unmarshal(jsonFile, &panelInfo)
		if err != nil {
			log.Println(err)
		}
	}

	thumb := getPanelThumb(dirname)
	panelInfo.Thumb = thumb

	return panelInfo
}

func getPanelThumb(dirname string) string {
	re := regexp.MustCompile(`^thumb\.(jpg|jpeg|png)$`)
	files, _ := os.ReadDir("../panels/" + dirname)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if re.MatchString(filepath.Base(file.Name())) {
			return filepath.Base(file.Name())
		}
	}

	return ""
}

func GetPanel(w http.ResponseWriter, r *http.Request) {
	html, _ := os.ReadFile("../panels/test_panel/index.html")
	css, _ := os.ReadFile("../panels/test_panel/output.css")

	type Response struct {
		Html string `json:"html"`
		Css  string `json:"css"`
	}

	resp := Response{Html: string(html), Css: string(css)}
	json.NewEncoder(w).Encode(resp)
}
