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
	"log"
	"os"
)

var logFile *os.File

func MCRMLogInit() {
	var err error
	// Open or create the log file with append permissions
	logFile, err = os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err) // Exit if we can't open the log file
	}

	// Optionally set log to write to file in addition to standard log output
	log.SetOutput(logFile)
}

func MCRMLog(v ...interface{}) {
	log.Println(v...) // Logs to terminal as well
}
