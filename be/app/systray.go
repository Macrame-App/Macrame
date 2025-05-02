package app

import (
	"be/app/helper"
	"os"

	"github.com/getlantern/systray"
)

func InitSystray() {
	go func() {
		systray.Run(OnReady, OnExit)
	}()
}

func OnReady() {
	systray.SetIcon(getFavicon())
	systray.SetTitle("Macrame")
	systray.SetTooltip("Macrame - Server")

	ip, _ := GetServerIp()

	systray.AddMenuItem("IP: "+ip+":"+helper.EnvGet("MCRM__PORT"), "")

	systray.AddSeparator()

	addMCRMItem("Dashboard", "/")
	addMCRMItem("Panels", "/panels")
	addMCRMItem("Macros", "/macros")
	addMCRMItem("Devices", "/devices")

	systray.AddSeparator()

	mQuit := systray.AddMenuItem("Quit Macrame", "Quit Macrame")
	go func() {
		<-mQuit.ClickedCh
		os.Exit(0)
	}()
}

func addMCRMItem(name, urlPath string) {
	m := systray.AddMenuItem(name, name)
	go func() {
		<-m.ClickedCh
		helper.OpenBrowser("http://localhost:" + helper.EnvGet("MCRM__PORT") + urlPath)
	}()
}

func OnExit() {
	systray.Quit()
}

func getFavicon() []byte {
	favicon, err := os.ReadFile("favicon.ico")

	if err != nil {
		MCRMLog("getFavicon Error: ", err)
	}

	return favicon
}
