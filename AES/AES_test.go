package AES

import (
	"fmt"
	"testing"
)

//测试padding  unpadding
func TestPadding(t *testing.T) {
	tool := NewAesTool([]byte{}, 16)
	src := []byte{1, 2, 3, 4, 5}
	src = tool.padding(src)
	fmt.Println(src)
	src = tool.unPadding(src)
	fmt.Println(src)
}

//测试AES ECB 加密解密
func TestEncryptDecrypt(t *testing.T) {
	key := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F}
	blickSize := 16
	tool := NewAesTool(key, blickSize)
	encryptData, _ := tool.Encrypt([]byte("32334erew32"))
	fmt.Println(encryptData)
	decryptData, _ := tool.Decrypt(encryptData)
	fmt.Println(string(decryptData))
}