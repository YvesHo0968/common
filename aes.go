package common

import (
	"crypto/aes"
	"crypto/cipher"
)

// AesEncryptECB 模式加密
func AesEncryptECB(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	data = Pkcs7Padding(data, block.BlockSize())
	b := make([]byte, len(data))
	for i := 0; i < len(data); i += block.BlockSize() {
		block.Encrypt(b[i:i+block.BlockSize()], data[i:i+block.BlockSize()])
	}
	return b, nil
}

// AesDecryptECB 模式解密
func AesDecryptECB(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := make([]byte, len(data))
	for i := 0; i < len(data); i += block.BlockSize() {
		block.Decrypt(b[i:i+block.BlockSize()], data[i:i+block.BlockSize()])
	}
	return UnPkcs7Padding(b), nil
}

// AesEncryptCBC 模式加密
func AesEncryptCBC(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	data = Pkcs7Padding(data, block.BlockSize())
	b := make([]byte, len(data))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(b, data)
	return b, nil
}

// AesDecryptCBC 模式解密
func AesDecryptCBC(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := make([]byte, len(data))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(b, data)
	return UnPkcs7Padding(b), nil
}

// AesEncryptCTR 模式加密
func AesEncryptCTR(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := make([]byte, len(data))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(b, data)
	return b, nil
}

// AesDecryptCTR 模式解密
func AesDecryptCTR(data, key, iv []byte) ([]byte, error) {
	return AesEncryptCTR(data, key, iv)
}

// AesEncryptCFB 加密器
func AesEncryptCFB(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	data = Pkcs7Padding(data, block.BlockSize())
	cfb := cipher.NewCFBEncrypter(block, iv[:block.BlockSize()])
	ciphertext := make([]byte, len(data))
	cfb.XORKeyStream(ciphertext, data)

	return ciphertext, nil
}

// AesDecryptCFB 模式解密
func AesDecryptCFB(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBDecrypter(block, iv[:block.BlockSize()])
	plaintext := make([]byte, len(data))
	cfb.XORKeyStream(plaintext, data)

	return UnPkcs7Padding(plaintext), nil
}

// AesEncryptOFB 加密器
func AesEncryptOFB(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	data = Pkcs7Padding(data, block.BlockSize())
	ofb := cipher.NewOFB(block, iv[:block.BlockSize()])
	ciphertext := make([]byte, len(data))
	ofb.XORKeyStream(ciphertext, data)

	return ciphertext, nil
}

// AesDecryptOFB 模式解密
func AesDecryptOFB(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ofb := cipher.NewOFB(block, iv[:block.BlockSize()])
	plaintext := make([]byte, len(data))
	ofb.XORKeyStream(plaintext, data)

	return UnPkcs7Padding(plaintext), nil
}
