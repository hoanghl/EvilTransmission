package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"io"
	"os"

	"github.com/zenazn/pkcs7pad"
)

const EncryptKeySize = 32 // Use AES-256

func GenSecureKey(n int) []byte {
	ret := make([]byte, n)
	nOut, err := rand.Read(ret)
	if err != nil || nOut != n {
		logger.Error("Cannot create random key: ", err)
	}

	return ret
}

func EncryptAES246(plainText []byte, key []byte) []byte {
	// Pad
	plainText = pkcs7pad.Pad(plainText, aes.BlockSize)

	// Prepare things
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	_, err := io.ReadFull(rand.Reader, iv)
	if err != nil {
		logger.Error("Create iv err: ", err)
		panic(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		logger.Error("Creating AES encrypter error: ", err)
		panic(err)
	}
	encrypter := cipher.NewCBCEncrypter(block, iv)

	// Encrypt
	encrypter.CryptBlocks(cipherText[aes.BlockSize:], plainText)

	return cipherText

}

func DecryptAES256(cipherText []byte, key []byte) []byte {
	// Create res needed for decrypting
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	outText := make([]byte, len(cipherText))
	block, err := aes.NewCipher(key)
	if err != nil {
		logger.Error("Creating AES decrypter error: ", err)
		panic(err)
	}

	// Decrypt
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypter.CryptBlocks(outText, cipherText)

	// Unpad
	plainText, err := pkcs7pad.Unpad(outText)
	if err != nil {
		logger.Errorf("Unpad error: %s", err)
	}

	return plainText
}

func GenRSAKey(bitSize int) *rsa.PrivateKey {
	privKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		logger.Errorf("Creating RSA private err: %s", err)
		panic(err)
	}

	return privKey

}

func LoadRSAKey(pathPemPriKey string, pathPemPubKey string) (*rsa.PrivateKey, *rsa.PublicKey) {
	bytesPriKey, _ := os.ReadFile(pathPemPriKey)
	bytesPubKey, _ := os.ReadFile(pathPemPubKey)

	priKey, _ := x509.ParsePKCS1PrivateKey(bytesPriKey)
	pubKey, _ := x509.ParsePKCS1PublicKey(bytesPubKey)

	return priKey, pubKey
}

func EncryptRSA4096(plainText []byte, priKey *rsa.PrivateKey, pubKey *rsa.PublicKey) []byte {
	cipher, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, plainText)
	if err != nil {
		logger.Errorln("Encrypting RSA-4096 error: ", err)
		panic(err)
	}

	return cipher
}

func DecryptRSA4096(cipherText []byte, priKey *rsa.PrivateKey) []byte {
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, cipherText)

	if err != nil {
		logger.Errorln("Decrypting RSA-4096 error: ", err)
		panic(err)
	}

	return plainText
}
