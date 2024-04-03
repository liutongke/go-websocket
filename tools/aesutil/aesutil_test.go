package aesutil

import (
	"crypto/aes"
	"testing"
)

func TestAESCipher(t *testing.T) {
	key := []byte("1234567890123456") // 16字节密钥
	plaintext := []byte("hello, world!")

	// 测试CBC模式
	iv := make([]byte, aes.BlockSize)
	cbcCipher, err := NewAESCipher(key, CBCMode, PKCS7Padding, iv)
	if err != nil {
		t.Errorf("Error creating AESCipher in CBC mode: %v", err)
		return
	}

	ciphertext, err := cbcCipher.Encrypt(plaintext)
	if err != nil {
		t.Errorf("Error encrypting in CBC mode: %v", err)
		return
	}

	decrypted, err := cbcCipher.Decrypt(ciphertext)
	if err != nil {
		t.Errorf("Error decrypting in CBC mode: %v", err)
		return
	}

	if string(decrypted) != string(plaintext) {
		t.Errorf("Decrypted text does not match original text in CBC mode: expected %s, got %s", plaintext, decrypted)
	}

	// 测试ECB模式
	ecbCipher, err := NewAESCipher(key, ECBMode, PKCS7Padding, nil)
	if err != nil {
		t.Errorf("Error creating AESCipher in ECB mode: %v", err)
		return
	}

	ciphertext, err = ecbCipher.Encrypt(plaintext)
	if err != nil {
		t.Errorf("Error encrypting in ECB mode: %v", err)
		return
	}

	decrypted, err = ecbCipher.Decrypt(ciphertext)
	if err != nil {
		t.Errorf("Error decrypting in ECB mode: %v", err)
		return
	}

	if string(decrypted) != string(plaintext) {
		t.Errorf("Decrypted text does not match original text in ECB mode: expected %s, got %s", plaintext, decrypted)
	}
}
