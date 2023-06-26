package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// Pkcs7Pad pks7文本补充
func Pkcs7Pad(plaintext []byte, blockSize int) []byte {
	padding := blockSize - (len(plaintext) % blockSize)
	paddedText := append(plaintext, bytes.Repeat([]byte{byte(padding)}, padding)...)
	return paddedText
}

// Pkcs7UnPad pks7文本去除填充
func Pkcs7UnPad(plaintext []byte) []byte {
	padding := int(plaintext[len(plaintext)-1])
	unPaddedText := plaintext[:len(plaintext)-padding]
	return unPaddedText
}

type ecbEncrypt struct {
	b         cipher.Block
	blockSize int
}

func NewECBEncrypt(b cipher.Block) cipher.BlockMode {
	return &ecbEncrypt{b, b.BlockSize()}
}

func (x *ecbEncrypt) BlockSize() int { return x.blockSize }

func (x *ecbEncrypt) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypt struct {
	b         cipher.Block
	blockSize int
}

func NewECBDecrypt(b cipher.Block) cipher.BlockMode {
	return &ecbDecrypt{b, b.BlockSize()}
}

func (x *ecbDecrypt) BlockSize() int { return x.blockSize }

func (x *ecbDecrypt) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

func AesEncrypt(plaintext string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	paddedPlaintext := Pkcs7Pad([]byte(plaintext), aes.BlockSize)

	ciphertext := make([]byte, len(paddedPlaintext))

	ecb := NewECBEncrypt(block)
	ecb.CryptBlocks(ciphertext, paddedPlaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func AesDecrypt(ciphertext string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", err
	}

	if len(decodedCiphertext)%aes.BlockSize != 0 {
		return "", err
	}

	decryptedText := make([]byte, len(decodedCiphertext))

	ecb := NewECBDecrypt(block)
	ecb.CryptBlocks(decryptedText, decodedCiphertext)

	unPaddedText := Pkcs7UnPad(decryptedText)

	return string(unPaddedText), nil
}
