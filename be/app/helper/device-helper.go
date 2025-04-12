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
