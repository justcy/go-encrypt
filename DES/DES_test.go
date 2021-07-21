package DES

import (
	"testing"
)

//
//func TestEncrypt(t *testing.T) {
//	type args struct {
//		orig string
//		key  string
//	}
//	tests := []struct {
//		name string
//		args args
//		want string
//	}{
//		// TODO: Add test cases.
//		{name:"ass",args:args{orig:"1111",key:"12345678"},want:"0pk7MUdE9T8="},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := Encrypt(tt.args.orig, tt.args.key); got != tt.want {
//				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestDecrypt(t *testing.T) {
//	type args struct {
//		data string
//		key  string
//	}
//	tests := []struct {
//		name string
//		args args
//		want string
//	}{
//		// TODO: Add test cases.
//		{name:"ass",args:args{data:"0pk7MUdE9T8=",key:"12345678"},want:"1111"},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := Decrypt(tt.args.data, tt.args.key); got != tt.want {
//				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestEncrypt(t *testing.T) {
	type args struct {
		text string
		key  []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "ass", args: args{text: "1111", key: []byte("12345678")}, want: "851cfe05a90ecd29"},
		{name: "ass", args: args{text: "1111", key: []byte("123456")}, wantErr:true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encrypt(tt.args.text, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type args struct {
		decrypted string
		key       []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "ass", args: args{decrypted: "851cfe05a90ecd29", key: []byte("12345678")}, want: "1111"},
		{name: "ass", args: args{decrypted: "851cfe05a90ecd", key: []byte("12345678")},wantErr:true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decrypt(tt.args.decrypted, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
