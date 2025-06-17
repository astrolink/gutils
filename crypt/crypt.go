package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

type Crypt struct {
	key []byte
	iv  []byte
}

// NewCrypt constrói uma instância do Crypt com chave e IV
func NewCrypt(keyHex, iniVector string) (*Crypt, error) {
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar chave hex: %s", err)
	}

	// IV (vetor de inicialização), deve ter exatamente 16 bytes
	iv := []byte(iniVector)
	return &Crypt{key: key, iv: iv}, nil
}

// Decrypt executa a descriptografia
func (c *Crypt) Decrypt(encrypted string) (string, error) {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return "", fmt.Errorf("erro ao criar cipher AES: %s", err)
	}

	// modo CBC
	mode := cipher.NewCBCDecrypter(block, c.iv)

	// base64 decode (como openssl_encrypt gera base64)
	ciphertextRaw, err := decodeBase64(encrypted)
	if err != nil {
		return "", fmt.Errorf("erro ao decodificar base64: %s", err)
	}

	if len(ciphertextRaw)%aes.BlockSize != 0 {
		return "", errors.New("tamanho inválido do texto criptografado")
	}

	decrypted := make([]byte, len(ciphertextRaw))
	mode.CryptBlocks(decrypted, ciphertextRaw)

	// remove padding (PKCS#7)
	decrypted, err = pkcs7Unpad(decrypted)
	if err != nil {
		return "", fmt.Errorf("erro ao remover padding: %s", err)
	}

	// hex -> string
	decoded, err := hex.DecodeString(string(decrypted))
	if err != nil {
		return "", fmt.Errorf("erro ao decodificar hex: %s", err)
	}

	return string(decoded), nil
}

func decodeBase64(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

func pkcs7Unpad(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("dados vazios")
	}

	padding := int(data[length-1])
	if padding == 0 || padding > length {
		return nil, errors.New("padding inválido")
	}

	// Valida todos os bytes de padding
	for i := 0; i < padding; i++ {
		if data[length-1-i] != byte(padding) {
			return nil, errors.New("conteúdo do padding inválido")
		}
	}

	return data[:length-padding], nil
}
