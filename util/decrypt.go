package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func Decrypt(data string) (string, error) {
	key := []byte("key-jti-techtest")

	r, _ := base64.StdEncoding.DecodeString(data)
	origData, err := AesDecrypt(r, key)
	if err != nil {
		return "", err
	}

	return string(origData), nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := []byte("1234567890123456")
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
