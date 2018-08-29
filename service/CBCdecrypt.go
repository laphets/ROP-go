package service

import (
	"crypto/aes"
	"crypto/cipher"
	"bytes"
	"github.com/spf13/viper"
	"encoding/hex"
)

func FormDecrypt(textBase64 string) (string, error) {
	//log.Debugf("%s", textBase64)
	//text, err := base64.URLEncoding.DecodeString(textBase64)
	text, err := hex.DecodeString(textBase64)
	if err != nil {
		return "", err
	}

	de, err := decrypt(text, []byte(viper.GetString("form_ase_secret")))
	if err != nil {
		return "", err
	}

	return string(de), nil
}

func decrypt(ciphertext, key []byte) ([]byte, error) {
	pkey := paddingLeft(key, '0', 16)
	block, err := aes.NewCipher(pkey) //选择加密算法
	if err != nil {
		return nil, err
	}
	blockModel := cipher.NewCBCDecrypter(block, pkey)
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, []byte(ciphertext))
	plantText = pKCS7UnPadding(plantText, block.BlockSize())
	return plantText, nil
}

func pKCS7UnPadding(plantText []byte, blockSize int) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

func paddingLeft(ori []byte, pad byte, length int) []byte {
	if len(ori) >= length {
		return ori[:length]
	}
	pads := bytes.Repeat([]byte{pad}, length-len(ori))
	return append(pads, ori...)
}