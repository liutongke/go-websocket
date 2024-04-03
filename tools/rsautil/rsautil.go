package rsautil

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// PaddingScheme 填充方案
type PaddingScheme int

const (
	// PKCS1v15 表示PKCS#1 v1.5填充
	PKCS1v15 PaddingScheme = iota
	// OAEP 表示RSA-OAEP填充
	OAEP
)

// GenerateRSAKeyPair 生成RSA密钥对
func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, bits)
}

// RSASign 使用RSA私钥对消息进行签名
func RSASign(privateKey *rsa.PrivateKey, message []byte) ([]byte, error) {
	hashed := sha256.Sum256(message)
	return rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
}

// RSAVerify 使用RSA公钥验证签名
func RSAVerify(publicKey *rsa.PublicKey, message, signature []byte) error {
	hashed := sha256.Sum256(message)
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
}

// RSAEncrypt 使用RSA公钥加密消息
func RSAEncrypt(publicKey *rsa.PublicKey, plaintext []byte, scheme PaddingScheme) ([]byte, error) {
	switch scheme {
	case PKCS1v15:
		return rsa.EncryptPKCS1v15(rand.Reader, publicKey, plaintext)
	case OAEP:
		return rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, plaintext, nil)
	default:
		return nil, errors.New("unsupported padding scheme")
	}
}

// RSADecrypt 使用RSA私钥解密消息
func RSADecrypt(privateKey *rsa.PrivateKey, ciphertext []byte, scheme PaddingScheme) ([]byte, error) {
	switch scheme {
	case PKCS1v15:
		return rsa.DecryptPKCS1v15(nil, privateKey, ciphertext)
	case OAEP:
		return rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
	default:
		return nil, errors.New("unsupported padding scheme")
	}
}

// PrivateKeyToPEMString 将RSA私钥转换为PEM格式的字符串
func PrivateKeyToPEMString(privateKey *rsa.PrivateKey) (string, error) {
	derBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derBytes,
	}
	return string(pem.EncodeToMemory(block)), nil
}

// PublicKeyToPEMString 将RSA公钥转换为PEM格式的字符串
func PublicKeyToPEMString(publicKey *rsa.PublicKey) (string, error) {
	derBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", err
	}
	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derBytes,
	}
	return string(pem.EncodeToMemory(block)), nil
}

// ParsePrivateKeyFromPEMString 从PEM格式的字符串解析RSA私钥
func ParsePrivateKeyFromPEMString(pemStr string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the private key")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// ParsePublicKeyFromPEMString 从PEM格式的字符串解析RSA公钥
func ParsePublicKeyFromPEMString(pemStr string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("failed to parse public key: not an RSA public key")
	}
	return pub, nil
}
