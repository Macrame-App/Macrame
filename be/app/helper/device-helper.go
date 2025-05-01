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

package helper

import (
	"errors"
	"os"
	"time"
)

func TempPinFile(Uuid string, pin string) (bool, error) {
	err := os.WriteFile("devices/"+Uuid+".tmp", []byte(pin), 0644)
	if err != nil {
		return false, errors.New("TempPinFile Error: " + err.Error())
	}

	time.AfterFunc(1*time.Minute, func() {
		os.Remove("devices/" + Uuid + ".tmp")
	})

	return true, nil
}

func RemovePinFile(Uuid string) error { return os.Remove("devices/" + Uuid + ".tmp") }

func CheckPinFile(Uuid string) bool {
	_, err := os.Stat("devices/" + Uuid + ".tmp")
	return err == nil
}

func SaveDeviceKey(Uuid string, key string) error {
	err := os.WriteFile("devices/"+Uuid+".key", []byte(key), 0644)

	if err != nil {
		return err
	}

	return nil
}

func GetKeyByUuid(Uuid string) (string, error) {
	data, err := os.ReadFile("devices/" + Uuid + ".key")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
