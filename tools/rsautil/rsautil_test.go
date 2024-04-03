package rsautil

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestRsaPEM(t *testing.T) {
	privateKey, _ := GenerateRSAKeyPair(2048)
	publicKey := &privateKey.PublicKey

	// 转换私钥为字符串
	privateKeyStr, _ := PrivateKeyToPEMString(privateKey)
	fmt.Println("Private Key (PEM):")
	fmt.Println(privateKeyStr)

	// 转换公钥为字符串
	publicKeyStr, _ := PublicKeyToPEMString(publicKey)
	fmt.Println("Public Key (PEM):")
	fmt.Println(publicKeyStr)

	// 从字符串解析私钥
	parsedPrivateKey, _ := ParsePrivateKeyFromPEMString(privateKeyStr)
	fmt.Println("Parsed Private Key:")
	fmt.Println(parsedPrivateKey)

	// 从字符串解析公钥
	parsedPublicKey, _ := ParsePublicKeyFromPEMString(publicKeyStr)
	fmt.Println("Parsed Public Key:")
	fmt.Println(parsedPublicKey)
}

func TestRsa(t *testing.T) {
	// 生成RSA密钥对
	privateKey, err := GenerateRSAKeyPair(2048)
	if err != nil {
		fmt.Println("Error generating RSA private key:", err)
		return
	}

	// 使用RSA私钥进行签名
	message := []byte("hello, world")
	signature, err := RSASign(privateKey, message)
	if err != nil {
		fmt.Println("Error signing message:", err)
		return
	}

	// 使用RSA公钥验证签名
	publicKey := &privateKey.PublicKey
	err = RSAVerify(publicKey, message, signature)
	if err != nil {
		fmt.Println("Signature verification failed:", err)
		return
	}
	fmt.Println("Signature verified successfully!")

	// 使用RSA公钥加密消息
	ciphertext, err := RSAEncrypt(publicKey, message, PKCS1v15)
	if err != nil {
		fmt.Println("Error encrypting message:", err)
		return
	}

	// 对字符串进行Base64编码
	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Println("加密后数据:", encoded)

	// 对Base64编码的字符串进行解码
	//decoded, err := base64.StdEncoding.DecodeString(encoded)
	//if err != nil {
	//	fmt.Println("Error decoding:", err)
	//	return
	//}
	//fmt.Println("Decoded:", string(decoded))
	// 使用RSA私钥解密消息
	plaintext, err := RSADecrypt(privateKey, ciphertext, PKCS1v15)
	if err != nil {
		fmt.Println("Error decrypting message:", err)
		return
	}
	fmt.Println("Decrypted message:", string(plaintext))
}
