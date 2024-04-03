package aesutil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

// AESMode 表示AES加密模式
type AESMode int

const (
	// CBCMode 表示CBC模式
	CBCMode AESMode = iota
	// ECBMode 表示ECB模式
	ECBMode
)

// AESPadding 表示AES填充方式
type AESPadding int

const (
	// PKCS7Padding 表示PKCS7填充方式
	PKCS7Padding AESPadding = iota
	// ZeroPadding 表示Zero填充方式
	ZeroPadding
)

// AESCipher 表示AES加密器
type AESCipher struct {
	key       []byte
	block     cipher.Block
	blockSize int
	mode      AESMode
	padding   AESPadding
	iv        []byte // 初始化向量
}

// NewAESCipher 创建一个AES加密器
func NewAESCipher(key []byte, mode AESMode, padding AESPadding, iv []byte) (*AESCipher, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	return &AESCipher{
		key:       key,
		block:     block,
		blockSize: blockSize,
		mode:      mode,
		padding:   padding,
		iv:        iv,
	}, nil
}

// Encrypt 使用AES加密器对数据进行加密
func (a *AESCipher) Encrypt(plaintext []byte) ([]byte, error) {
	// 对明文进行填充
	plaintext = a.pad(plaintext)

	var iv []byte
	if a.mode == CBCMode {
		if len(a.iv) != a.blockSize {
			return nil, errors.New("IV length must be equal to block size")
		}
		iv = a.iv
	}

	// 选择加密模式
	var encrypter cipher.BlockMode
	switch a.mode {
	case CBCMode:
		encrypter = cipher.NewCBCEncrypter(a.block, iv)
	case ECBMode:
		encrypter = NewECBEncrypter(a.block)
	default:
		return nil, errors.New("unsupported AES mode")
	}

	ciphertext := make([]byte, len(plaintext))
	encrypter.CryptBlocks(ciphertext, plaintext)
	return ciphertext, nil
}

// Decrypt 使用AES加密器对数据进行解密
func (a *AESCipher) Decrypt(ciphertext []byte) ([]byte, error) {
	// 选择解密模式
	var decrypter cipher.BlockMode
	switch a.mode {
	case CBCMode:
		if len(a.iv) != a.blockSize {
			return nil, errors.New("IV length must be equal to block size")
		}
		iv := a.iv
		decrypter = cipher.NewCBCDecrypter(a.block, iv)
	case ECBMode:
		decrypter = NewECBDecrypter(a.block)
	default:
		return nil, errors.New("unsupported AES mode")
	}

	plaintext := make([]byte, len(ciphertext))
	decrypter.CryptBlocks(plaintext, ciphertext)

	// 对解密后的数据进行去填充
	plaintext = a.unpad(plaintext)

	return plaintext, nil
}

// pad 对明文进行填充
func (a *AESCipher) pad(plaintext []byte) []byte {
	padding := a.blockSize - len(plaintext)%a.blockSize
	var padText []byte
	switch a.padding {
	case PKCS7Padding:
		padText = bytes.Repeat([]byte{byte(padding)}, padding)
	case ZeroPadding:
		padText = bytes.Repeat([]byte{byte(0)}, padding)
	default:
		padText = nil
	}
	return append(plaintext, padText...)
}

// unpad 对解密后的数据进行去填充
func (a *AESCipher) unpad(plaintext []byte) []byte {
	padding := int(plaintext[len(plaintext)-1])
	return plaintext[:len(plaintext)-padding]
}

// NewECBEncrypter 创建一个ECB加密器
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return &ecbEncrypter{b}
}

type ecbEncrypter struct{ b cipher.Block }

func (x *ecbEncrypter) BlockSize() int { return x.b.BlockSize() }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.b.BlockSize() != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.b.BlockSize()])
		src = src[x.b.BlockSize():]
		dst = dst[x.b.BlockSize():]
	}
}

// NewECBDecrypter 创建一个ECB解密器
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return &ecbDecrypter{b}
}

type ecbDecrypter struct{ b cipher.Block }

func (x *ecbDecrypter) BlockSize() int { return x.b.BlockSize() }

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.b.BlockSize() != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.b.BlockSize()])
		src = src[x.b.BlockSize():]
		dst = dst[x.b.BlockSize():]
	}
}
