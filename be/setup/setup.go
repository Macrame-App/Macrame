package main

import "be/app/helper"

func main() {
	helper.CreateConfigFile("../public/config.js")
	helper.CheckFeDevDir()

	return
}
