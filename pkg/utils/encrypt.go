package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

var Encrypt *Encryption

// AES 对称加密
type Encryption struct {
	key string
}

func init() {
	Encrypt = NewEncryption()
}

func NewEncryption() *Encryption {
	return &Encryption{}
}

// PadPwd 填充密码长度
func PadPwd(srcByte []byte, blockSize int) []byte {
	padNum := blockSize - len(srcByte)%blockSize
	ret := bytes.Repeat([]byte{byte(padNum)}, padNum)
	srcByte = append(srcByte, ret...)
	return srcByte
}

// AesEncoding加密
func (k *Encryption) AesEncoding(src string) string {
	key := []byte(k.key)
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return ""
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}

	blockSize := block.BlockSize()
	origData := PadPwd([]byte(src), blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return base64.StdEncoding.EncodeToString(crypted)
}

// UnPadPwd
func UnPadPwd(dst []byte) ([]byte, error) {
	length := len(dst)
	if length == 0 {
		return nil, errors.New("invalid padding size")
	}
	unPadNum := int(dst[length-1])
	if unPadNum == 0 || unPadNum > length {
		return nil, errors.New("invalid padding size")
	}
	return dst[:(length - unPadNum)], nil
}

// AesDecoding
func (k *Encryption) AesDecoding(src string) string {
	key := []byte(k.key)
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return ""
	}

	crypted, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return ""
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}

	blockSize := block.BlockSize()
	if len(crypted)%blockSize != 0 {
		return ""
	}

	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData, err = UnPadPwd(origData)
	if err != nil {
		return ""
	}
	return string(origData)
}

func (k *Encryption) SetKey(key string) {
	k.key = key
}
