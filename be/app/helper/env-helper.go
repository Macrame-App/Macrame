package helper

import (
	"log"
	"net"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func EnvGet(key string) string {
	envFile := "../.env"
	fileExists := func() bool {
		_, err := os.Stat(envFile)
		return err == nil
	}
	if !fileExists() {
		createEnvFile(envFile)
	}
	err := godotenv.Load(envFile)
	if err != nil {
		return ""
	}
	return os.Getenv("VITE_" + key)
}

func createEnvFile(filename string) {
	log.Println("Creating .env file...")
	file, err := os.Create(filename)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	// You can add some default values to the .env file here if needed
	// For example:
	port, err := findOpenPort()
	saltKey := GenerateKey()
	salt := saltKey[:28]
	iv := GenerateRandomIntegerString(16)

	log.Println(err, saltKey, iv)

	_, err = file.WriteString("VITE_MCRM__PORT=" + string(port) + "\nVITE_MCRM__SALT=" + salt + "\nVITE_MCRM__IV=" + iv)
	if err != nil {
		log.Fatal(err)
	}
}

func findOpenPort() (string, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return "", err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return "", err
	}
	defer l.Close()
	return strconv.Itoa(l.Addr().(*net.TCPAddr).Port), nil
}
