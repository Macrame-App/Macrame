package main

import "be/app/helper"

func main() {
	helper.CreateConfigFile("../public/config.js")
	helper.CheckFeDevDir()

	port := helper.EnvGet("MCRM__PORT")

	helper.MakeCaddyFile("CaddyFile", port)

	return
}
