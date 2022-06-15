package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var salt = []byte{75, 63, 69, 62, 64, 66, 20, 73, 75, 70, 65, 72, 20, 74, 68, 65}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func ModifyPasswordLen(password string) string {
	if len(password) < 16 {
		password = password + string(salt[:16-len(password)])
	} else if len(password) > 16 {
		password = password[:16]
	}

	return password
}

func Encrypt(text, password string) (string, error) {
	password = ModifyPasswordLen(password)
	block, err := aes.NewCipher([]byte(password))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, salt)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func Decrypt(text, password string) (string, error) {
	password = ModifyPasswordLen(password)
	block, err := aes.NewCipher([]byte(password))
	if err != nil {
		return "", err
	}
	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, salt)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
