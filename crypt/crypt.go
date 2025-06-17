package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"errors"
)

type Crypt struct {
	key []byte
	iv  []byte
}

// NewCrypt constrói uma instância do Crypt com chave e IV
func NewCrypt(keyHex, iniVector string) (*Crypt, error) {
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return nil, err
	}

	// IV (vetor de inicialização), deve ter exatamente 16 bytes
	iv := []byte(iniVector)
	return &Crypt{key: key, iv: iv}, nil
}

// Decrypt executa a descriptografia
func (c *Crypt) Decrypt(encrypted string) (string, error) {
	// descriptografa direto do base64
	ciphertext := []byte(encrypted)

	block, err := aes.NewCipher(c.key)
	if err != nil {
		return "", err
	}

	// modo CBC
	mode := cipher.NewCBCDecrypter(block, c.iv)

	// base64 decode (como openssl_encrypt gera base64)
	ciphertextRaw, err := decodeBase64(ciphertext)
	if err != nil {
		return "", err
	}

	if len(ciphertextRaw)%aes.BlockSize != 0 {
		return "", errors.New("tamanho inválido do texto criptografado")
	}

	decrypted := make([]byte, len(ciphertextRaw))
	mode.CryptBlocks(decrypted, ciphertextRaw)

	// remove padding (PKCS#7)
	decrypted = pkcs7Unpad(decrypted)

	// hex -> string
	decoded, err := hex.DecodeString(string(decrypted))
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}

func decodeBase64(data []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(data))
}

func pkcs7Unpad(data []byte) []byte {
	length := len(data)
	if length == 0 {
		return data
	}
	padding := int(data[length-1])
	if padding > length {
		return data
	}
	return data[:length-padding]
}
