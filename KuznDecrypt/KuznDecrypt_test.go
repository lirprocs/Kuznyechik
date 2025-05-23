package KuznDecrypt

import (
	"github.com/lirprocs/Kuznyechik/KuznEncrypt"
	"testing"
)

func TestSInv(t *testing.T) {
	tests := []struct {
		expected, data [16]byte
	}{
		{[16]byte{0xff, 0xee, 0xdd, 0xcc, 0xbb, 0xaa, 0x99, 0x88, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x00}, [16]byte{0xb6, 0x6c, 0xd8, 0x88, 0x7d, 0x38, 0xe8, 0xd7, 0x77, 0x65, 0xae, 0xea, 0x0c, 0x9a, 0x7e, 0xfc}},
		{[16]byte{0xb6, 0x6c, 0xd8, 0x88, 0x7d, 0x38, 0xe8, 0xd7, 0x77, 0x65, 0xae, 0xea, 0x0c, 0x9a, 0x7e, 0xfc}, [16]byte{0x55, 0x9d, 0x8d, 0xd7, 0xbd, 0x06, 0xcb, 0xfe, 0x7e, 0x7b, 0x26, 0x25, 0x23, 0x28, 0x0d, 0x39}},
		{[16]byte{0x55, 0x9d, 0x8d, 0xd7, 0xbd, 0x06, 0xcb, 0xfe, 0x7e, 0x7b, 0x26, 0x25, 0x23, 0x28, 0x0d, 0x39}, [16]byte{0x0c, 0x33, 0x22, 0xfe, 0xd5, 0x31, 0xe4, 0x63, 0x0d, 0x80, 0xef, 0x5c, 0x5a, 0x81, 0xc5, 0x0b}},
		{[16]byte{0x0c, 0x33, 0x22, 0xfe, 0xd5, 0x31, 0xe4, 0x63, 0x0d, 0x80, 0xef, 0x5c, 0x5a, 0x81, 0xc5, 0x0b}, [16]byte{0x23, 0xae, 0x65, 0x63, 0x3f, 0x84, 0x2d, 0x29, 0xc5, 0xdf, 0x52, 0x9c, 0x13, 0xf5, 0xac, 0xda}},
	}
	for _, tt := range tests {
		res := sInv(tt.data)
		if res != tt.expected {
			t.Errorf("sInv() = %X; want %X", res, tt.expected)
		}
	}
}

func TestRInv(t *testing.T) {
	tests := []struct {
		expected, data [16]byte
	}{
		{[16]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00}, [16]byte{0x94, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}},
		{[16]byte{0x94, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, [16]byte{0xa5, 0x94, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{[16]byte{0xa5, 0x94, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, [16]byte{0x64, 0xa5, 0x94, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{[16]byte{0x64, 0xa5, 0x94, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, [16]byte{0x0d, 0x64, 0xa5, 0x94, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
	}
	for _, tt := range tests {
		res := rInv(tt.data)
		if res != tt.expected {
			t.Errorf("rInv() = %X; want %X", res, tt.expected)
		}
	}
}

func TestLInv(t *testing.T) {
	tests := []struct {
		expected, data [16]byte
	}{
		{[16]byte{0x64, 0xa5, 0x94, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, [16]byte{0xd4, 0x56, 0x58, 0x4d, 0xd0, 0xe3, 0xe8, 0x4c, 0xc3, 0x16, 0x6e, 0x4b, 0x7f, 0xa2, 0x89, 0x0d}},
		{[16]byte{0xd4, 0x56, 0x58, 0x4d, 0xd0, 0xe3, 0xe8, 0x4c, 0xc3, 0x16, 0x6e, 0x4b, 0x7f, 0xa2, 0x89, 0x0d}, [16]byte{0x79, 0xd2, 0x62, 0x21, 0xb8, 0x7b, 0x58, 0x4c, 0xd4, 0x2f, 0xbc, 0x4f, 0xfe, 0xa5, 0xde, 0x9a}},
		{[16]byte{0x79, 0xd2, 0x62, 0x21, 0xb8, 0x7b, 0x58, 0x4c, 0xd4, 0x2f, 0xbc, 0x4f, 0xfe, 0xa5, 0xde, 0x9a}, [16]byte{0x0e, 0x93, 0x69, 0x1a, 0x0c, 0xfc, 0x60, 0x40, 0x8b, 0x7b, 0x68, 0xf6, 0x6b, 0x51, 0x3c, 0x13}},
		{[16]byte{0x0e, 0x93, 0x69, 0x1a, 0x0c, 0xfc, 0x60, 0x40, 0x8b, 0x7b, 0x68, 0xf6, 0x6b, 0x51, 0x3c, 0x13}, [16]byte{0xe6, 0xa8, 0x09, 0x4f, 0xee, 0x0a, 0xa2, 0x04, 0xfd, 0x97, 0xbc, 0xb0, 0xb4, 0x4b, 0x85, 0x80}},
	}
	for _, tt := range tests {
		res := lInv(tt.data)
		if res != tt.expected {
			t.Errorf("lInv() = %X; want %X", res, tt.expected)
		}
	}
}

func TestDecryptText(t *testing.T) {
	key := [32]byte{0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, 0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77,
		0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}
	plainText := [16]byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x00, 0xff, 0xee, 0xdd, 0xcc, 0xbb, 0xaa, 0x99, 0x88}

	ct := KuznEncrypt.Encrypt(plainText, key)

	res := Decrypt(ct, key)
	if res != plainText {
		t.Errorf("Decrypt() = %X; want %X", res, plainText)
	}
}
