package parse

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"strconv"
	"strings"
)

func MakeKey(key string) ([]byte, error) {
	if len([]byte(key)) < 32 {
		return nil, errors.New("Error: Key must be at least 32 bytes")
	}
	return []byte(key[0:32]), nil
}

func StrBytesToBytes(inString string) ([]byte, error) {
	stringArr := strings.Split(inString[1:len(inString)-2], " ")
	outBytes := make([]byte, len(stringArr))
	for i, s := range stringArr {
		si, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		outBytes[i] = byte(si)
	}
	return outBytes, nil
}

func Encrypt(key []byte, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func Decrypt(key []byte, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}
