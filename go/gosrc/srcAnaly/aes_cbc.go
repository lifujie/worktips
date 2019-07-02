package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"log"
	"runtime/debug"
)

func pkcspadding(src []byte, blocksize int) []byte {
	padnum := blocksize - len(src)%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
	return append(src, pad...)
}

func pkcsunpadding(src []byte) []byte {
	n := len(src)
	unpadnum := int(src[n-1])
	return src[:n-unpadnum]
}

func encryptAES(src []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	src = pkcspadding(src, block.BlockSize())
	blockmode := cipher.NewCBCEncrypter(block, key)
	blockmode.CryptBlocks(src, src)
	return src
}

func decryptAES(src []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockmode := cipher.NewCBCDecrypter(block, key)
	blockmode.CryptBlocks(src, src)
	src = pkcsunpadding(src)
	return src
}

// EncryptAESCBC aes-cbc加密
func EncryptAESCBC(password string) (result string, err error) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("[%s]EncryptAESCBC: \n%s\n%s", err, debug.Stack())
		}
	}()
	key, _ := base64.StdEncoding.DecodeString(string(key))
	src := []byte(password)
	block, _ := aes.NewCipher(key)
	src = pkcspadding(src, block.BlockSize())
	blockmode := cipher.NewCBCEncrypter(block, key)
	blockmode.CryptBlocks(src, src)

	result = base64.StdEncoding.EncodeToString(src)
	return result, err
}

// DecryptAESCBC aes-cbc解密
func DecryptAESCBC(password string) (result string, err error) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("[%s]DecryptAESCBC: \n%s\n%s", err, debug.Stack())
		}
	}()
	src, err := base64.StdEncoding.DecodeString(password)
	if err != nil {
		return "", err
	}
	key, _ := base64.StdEncoding.DecodeString(string(key))
	block, _ := aes.NewCipher(key)
	blockmode := cipher.NewCBCDecrypter(block, key)
	blockmode.CryptBlocks(src, src)
	src = pkcsunpadding(src)
	return string(src), nil
}

func tostring(input []byte) {
	encodeString := base64.StdEncoding.EncodeToString(input)
	fmt.Println("en", encodeString)

	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(decodeBytes))
}

var (
	key = []byte("CzV6QgcfVh5gKgJOM/ExJg==")
)

func main() {
	x := []byte("o2GpJzzCaCkjbs9cX9UERixQa91hOXSNB+68HyS/Awc=")
	// key := []byte("")
	//x := []byte("123")
	// key := []byte("EnLT3bAkj+4/nRc5")

	decodeBytes, _ := base64.StdEncoding.DecodeString(string(x))
	// // fmt.Println(string(decodeBytes))
	// // key = decodeBytes
	// fmt.Printf("len: %d, x: %d\n", len(key), len(x))
	// // x1 := encryptAES(x, key)
	// // tostring(x1)
	// //x2 := decryptAES(x1, key)
	// x1, _ := EncryptAESCBC(string(x))
	// x2, _ := DecryptAESCBC(x1)
	// fmt.Printf("%s, %s\n", x2, x1)
	fmt.Printf("%x\n", decodeBytes)
}
