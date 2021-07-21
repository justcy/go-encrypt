package RSA

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	GenerateRSAKey(2048,".")
	GenerateRSAKey(2048,".")
	message:=[]byte("hello world")
	//加密
	cipherText:=Encrypt(message,"public.pem")
	fmt.Println("加密后为：",string(cipherText))
	//解密
	plainText := Decrypt(cipherText, "private.pem")
	fmt.Println("解密后为：",string(plainText))
	//type args struct {
	//	plainText []byte
	//	path      string
	//}
	//tests := []struct {
	//	name string
	//	args args
	//	want []byte
	//}{
	//	// TODO: Add test cases.
	//	{name: "encrypt", args: args{plainText: []byte("1111"), path: "public.pem"},want:[]byte("111")},
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := Encrypt(tt.args.plainText, tt.args.path); !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("Encrypt() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}
