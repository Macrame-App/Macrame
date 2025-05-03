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
	systray.SetIcon(getIcon("favicon.ico"))
	systray.SetTitle("Macrame")
	systray.SetTooltip("Macrame - Server")

	ip, err := GetServerIp()

	if err == nil {
		systray.AddMenuItem("IP: "+ip+":"+helper.EnvGet("MCRM__PORT"), "Server IP")
	}

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

func getIcon(path string) []byte {
	icon, err := os.ReadFile(path)

	if err != nil {
		MCRMLog("getIcon Error: ", err)
	}

	return icon
}
