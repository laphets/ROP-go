package service

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/spf13/viper"
	"io"
	"strings"
)

func addBase64Padding(value string) string {
	m := len(value) % 4
	if m != 0 {
		value += strings.Repeat("=", 4-m)
	}

	return value
}

func removeBase64Padding(value string) string {
	return strings.Replace(value, "=", "", -1)
}

func Pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func Unpad(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])

	if unpadding > length {
		return nil, errors.New("unpad error. This could happen when incorrect encryption key is used")
	}

	return src[:(length - unpadding)], nil
}

func Encrypt(text string) (string, error) {
	key := []byte(viper.GetString("aes_secret"))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	msg := Pad([]byte(text))
	ciphertext := make([]byte, aes.BlockSize+len(msg))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(msg))

	finalHex := hex.EncodeToString(ciphertext)
	//finalMsg := removeBase64Padding(base64.URLEncoding.EncodeToString(ciphertext))
	return finalHex, nil
}

func Decrypt(text string) (string, error) {
	key := []byte(viper.GetString("aes_secret"))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	//decodedMsg, err := base64.URLEncoding.DecodeString(addBase64Padding(text))
	decodedMsg, err := hex.DecodeString(text)

	if err != nil {
		return "", err
	}

	if (len(decodedMsg) % aes.BlockSize) != 0 {
		return "", errors.New("blocksize must be multipe of decoded message length")
	}

	iv := decodedMsg[:aes.BlockSize]
	msg := decodedMsg[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(msg, msg)

	unpadMsg, err := Unpad(msg)
	if err != nil {
		return "", err
	}

	return string(unpadMsg), nil
}