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

func CheckPinFile(encryptedData []byte) bool {
	return false
}
