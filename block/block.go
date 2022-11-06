package block

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func Blockencrypt(paths string, key string) {
	err := filepath.Walk(paths,
		func(path string, _ os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			file, _ := os.Open(path)
			defer file.Close()
			liste, _ := file.Readdirnames(0)
			for _, isim := range liste {
				dosya, _ := ioutil.ReadFile(path + isim)

				a, _ := Encrypt(dosya, key)
				veris := []byte(a)
				ioutil.WriteFile(path+isim, veris, 0644)

			}

			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
}
func Blockdecryption(paths string, key string) {
	err := filepath.Walk(paths,
		func(path string, _ os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			file, _ := os.Open(path)
			defer file.Close()
			liste, _ := file.Readdirnames(0)
			for _, isim := range liste {
				dosya, _ := ioutil.ReadFile(path + isim)

				a, _ := Decrypt(string(dosya), key)
				veris := []byte(a)
				ioutil.WriteFile(path+isim, veris, 0644)

			}

			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
}

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(text []byte, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	plainText := text
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

// Decrypt method is to extract back the encrypted text
func Decrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
