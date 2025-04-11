package helper

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"math"
	mathRand "math/rand"
	"strconv"
	"strings"
)

func EncryptAES(key string, plaintext string) (string, error) {
	origData := []byte(plaintext)

	// Create AES cipher
	block, err := aes.NewCipher(keyToBytes(key))
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()

	origData = PKCS5Padding(origData, blockSize)

	iv := []byte(EnvGet("MCRM__IV"))
	blockMode := cipher.NewCBCEncrypter(block, iv)

	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)

	cryptedString := base64.StdEncoding.EncodeToString(crypted)

	return cryptedString, nil
}

func DecryptAES(key string, cryptedText string) (string, error) {
	crypted, err := base64.StdEncoding.DecodeString(cryptedText)

	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyToBytes(key))
	if err != nil {
		return "", err
	}

	iv := []byte(EnvGet("MCRM__IV"))
	blockMode := cipher.NewCBCDecrypter(block, iv)

	origData := make([]byte, len(crypted))

	blockMode.CryptBlocks(origData, crypted)
	origData, err = PKCS5UnPadding(origData)

	if err != nil || len(origData) <= 3 {
		return "", errors.New("invalid key")
	}

	origDataString := string(origData)

	return origDataString, nil
}

func GenerateRandomString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

func GenerateRandomIntegerString(length int) string {
	min := int64(0)
	max := int64(math.Pow10(length))
	randInt := min + mathRand.Int63()%(max-min+1)
	return strconv.FormatInt(randInt, 10)
}

func GenerateKey() string {
	return strings.Replace(GenerateRandomString(24), "=", "", -1)
}

func keyToBytes(key string) []byte {
	// Convert key to bytes
	keyBytes := []byte(key)

	// If key is 4 characters, append salt
	if len(key) == 4 {
		keyBytes = []byte(key + EnvGet("MCRM__SALT"))
	}

	return keyBytes
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	unpadding := int(origData[length-1])

	if (unpadding >= length) || (unpadding == 0) {
		return nil, errors.New("unpadding error")
	}

	return origData[:(length - unpadding)], nil
}
