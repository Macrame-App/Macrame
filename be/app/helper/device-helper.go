package helper

import (
	"log"
	"os"
	"time"
)

func TempPinFile(Uuid string, pin string) bool {
	log.Println("temp pin file", Uuid, pin)
	err := os.WriteFile("devices/"+Uuid+".tmp", []byte(pin), 0644)
	if err != nil {
		log.Println(err)
		return false
	}

	time.AfterFunc(1*time.Minute, func() {
		log.Println("deleting", Uuid, pin)
		os.Remove("devices/" + Uuid + ".tmp")
	})

	return true
}

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
