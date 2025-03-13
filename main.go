package main

import "fmt"

var Pi = [256]byte{
	252, 238, 221, 17, 207, 110, 49, 22, 251, 196, 250, 218, 35, 197, 4, 77,
	233, 119, 240, 219, 147, 46, 153, 186, 23, 54, 241, 187, 20, 205, 95, 193,
	249, 24, 101, 90, 226, 92, 239, 33, 129, 28, 60, 66, 139, 1, 142, 79,
	5, 132, 2, 174, 227, 106, 143, 160, 6, 11, 237, 152, 127, 212, 211, 31,
	235, 52, 44, 81, 234, 200, 72, 171, 242, 42, 104, 162, 253, 58, 206, 204,
	181, 112, 14, 86, 8, 12, 118, 18, 191, 114, 19, 71, 156, 183, 93, 135,
	21, 161, 150, 41, 16, 123, 154, 199, 243, 145, 120, 111, 157, 158, 178, 177,
	50, 117, 25, 61, 255, 53, 138, 126, 109, 84, 198, 128, 195, 189, 13, 87,
	223, 245, 36, 169, 62, 168, 67, 201, 215, 121, 214, 246, 124, 34, 185, 3,
	224, 15, 236, 222, 122, 148, 176, 188, 220, 232, 40, 80, 78, 51, 10, 74,
	167, 151, 96, 115, 30, 0, 98, 68, 26, 184, 56, 130, 100, 159, 38, 65,
	173, 69, 70, 146, 39, 94, 85, 47, 140, 163, 165, 125, 105, 213, 149, 59,
	7, 88, 179, 64, 134, 172, 29, 247, 48, 55, 107, 228, 136, 217, 231, 137,
	225, 27, 131, 73, 76, 63, 248, 254, 141, 83, 170, 144, 202, 216, 133, 97,
	32, 113, 103, 164, 45, 43, 9, 91, 203, 155, 37, 208, 190, 229, 108, 82,
	89, 166, 116, 210, 230, 244, 180, 192, 209, 102, 175, 194, 57, 75, 99, 182}

var LFactors = [16]byte{148, 32, 133, 16, 194, 192, 1, 251, 1, 192, 194, 16, 133, 32, 148, 1}

func S(a [16]byte) [16]byte {
	for i := 0; i < 16; i++ {
		a[i] = Pi[a[i]]
	}
	return a
}

func R(a [16]byte) [16]byte {
	var z byte
	for i := 0; i < 16; i++ {
		z ^= GF256Mul(a[i], LFactors[i])
	}
	for i := 15; i > 0; i-- {
		a[i] = a[i-1]
	}
	a[0] = z
	return a
}

func L(a [16]byte) [16]byte {
	for i := 0; i < 16; i++ {
		a = R(a)
	}
	return a
}

func F(k [16]byte, a1, a0 [16]byte) ([16]byte, [16]byte) {
	tmp := XORBytes(L(S(XORBytes(a1, k))), a0)
	return tmp, a1
}

func XORBytes(a, b [16]byte) [16]byte {
	var c [16]byte
	for i := range a {
		c[i] = a[i] ^ b[i]
	}
	return c
}

func GF256Mul(a, b byte) byte {
	var p byte
	for i := 0; i < 8; i++ {
		if (b & 1) != 0 {
			p ^= a
		}
		hiBitSet := (a & 0x80) != 0
		a <<= 1
		if hiBitSet {
			a ^= 0xC3
		}
		b >>= 1
	}
	return p
}

func Vec128(i int) [16]byte {
	var v [16]byte
	v[15] = byte(i)
	return v
}

func Encrypt(plainText [16]byte, keys [10][16]byte) [16]byte {
	state := XORBytes(plainText, keys[0])
	for i := 1; i < 10; i++ {
		state = XORBytes(L(S(state)), keys[i])
	}
	return state
}

func KeySchedule(masterKey [32]byte) { // [10][16]byte {
	var keys [10][16]byte
	var C [32][16]byte

	copy(keys[0][:], masterKey[:16])
	copy(keys[1][:], masterKey[16:])

	fmt.Printf("K1: %x\n", keys[0])
	fmt.Printf("K2: %x\n", keys[1])

	for i := 0; i < 32; i++ {
		C[i] = L(Vec128(i + 1))
		//fmt.Printf("С%d: %x\n", i+1, C[i])
	}

	//Проверка
	//fmt.Printf(": %X\n", XORBytes(keys[0], C[0]))
	//fmt.Printf(": %X\n", S(XORBytes(keys[0], C[0])))
	//fmt.Printf(": %X\n", L(S(XORBytes(keys[0], C[0]))))

	for i := 0; i < 4; i++ {
		var left [9][16]byte
		var right [9][16]byte
		left[0] = keys[2*i]
		right[0] = keys[2*i+1]

		//fmt.Printf("----%X\n", left[i])
		//fmt.Printf("----%X\n", right[i])

		for j := 0; j < 8; j++ {
			left[j+1] = XORBytes(L(S(XORBytes(left[j], C[8*i+j]))), right[j])
			right[j+1] = left[j]
		}
		//fmt.Printf("%d: %X\n", left[8])
		keys[2*i+2] = left[8]
		keys[2*i+3] = right[8]
		fmt.Printf("%d: %X\n", 2*i+3, keys[2*i+2])
		fmt.Printf("%d: %X\n", 2*i+4, keys[2*i+3])
	}

	//var left [8][16]byte
	//var right [8][16]byte
	//
	//left[0] = XORBytes(L(S(XORBytes(keys[0], C[0]))), keys[1])
	//right[0] = keys[0]
	//fmt.Printf("%X\n", left[0])
	//fmt.Printf("%X\n", right[0])
	//
	//left[1] = XORBytes(L(S(XORBytes(left[0], C[1]))), right[0])
	//right[1] = left[0]
	//fmt.Printf("%X\n", left[1])
	//fmt.Printf("%X\n", right[1])
	//
	//left[2] = XORBytes(L(S(XORBytes(left[1], C[2]))), right[1])
	//right[2] = left[1]
	//fmt.Printf("%X\n", left[2])
	//fmt.Printf("%X\n", right[2])

}

//func KeySchedule(masterKey [32]byte) [10][16]byte {
//	var keys [10][16]byte
//	var C [32][16]byte
//
//	copy(keys[0][:], masterKey[:16])
//	copy(keys[1][:], masterKey[16:])
//
//	fmt.Printf("K1:%x\n", keys[0])
//	fmt.Printf("K2:%x\n", keys[1])
//
//	for i := 0; i < 32; i++ {
//		C[i] = L(Vec128(i))
//	}
//
//	for i := 1; i < 5; i++ {
//		for j := 0; j < 8; j++ {
//			keys[2*i], keys[2*i+1] = F(C[j], keys[2*i-2], keys[2*i-1])
//			fmt.Printf("С%d: %x\n", j, C[j])
//		}
//		fmt.Printf("K%d: %x\n", 2*i+1, keys[2*i])
//		fmt.Printf("K%d: %x\n", 2*i+2, keys[2*i+1])
//	}
//	return keys
//}

func main() {
	//// А.1.1 Преобразование S
	//S1 := [16]byte{0xff, 0xee, 0xdd, 0xcc, 0xbb, 0xaa, 0x99, 0x88, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x00}
	//S2 := [16]byte{0xb6, 0x6c, 0xd8, 0x88, 0x7d, 0x38, 0xe8, 0xd7, 0x77, 0x65, 0xae, 0xea, 0x0c, 0x9a, 0x7e, 0xfc}
	//S3 := [16]byte{0x55, 0x9d, 0x8d, 0xd7, 0xbd, 0x06, 0xcb, 0xfe, 0x7e, 0x7b, 0x26, 0x25, 0x23, 0x28, 0x0d, 0x39}
	//S4 := [16]byte{0x0c, 0x33, 0x22, 0xfe, 0xd5, 0x31, 0xe4, 0x63, 0x0d, 0x80, 0xef, 0x5c, 0x5a, 0x81, 0xc5, 0x0b}
	//fmt.Printf("S(input): %X\n", S(S1))
	//fmt.Printf("S(input): %X\n", S(S2))
	//fmt.Printf("S(input): %X\n", S(S3))
	//fmt.Printf("S(input): %X\n", S(S4))
	////S5 := [16]byte{0x23, 0xae, 0x65, 0x63, 0x3f, 0x84, 0x2d, 0x29, 0xc5, 0xdf, 0x52, 0x9c, 0x13, 0xf5, 0xac, 0xda}
	//
	//// А.1.2 Преобразование R
	//R1 := [16]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00}
	//R2 := [16]byte{0x94, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
	//R3 := [16]byte{0xa5, 0x94, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	//R4 := [16]byte{0x64, 0xa5, 0x94, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	//fmt.Printf("R(input): %X\n", R(R1))
	//fmt.Printf("R(input): %X\n", R(R2))
	//fmt.Printf("R(input): %X\n", R(R3))
	//fmt.Printf("R(input): %X\n", R(R4))
	////R5 := [16]byte{0x0d, 0x64, 0xa5, 0x94, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	//
	//// А.1.3 Преобразование L
	//L1 := [16]byte{0x64, 0xa5, 0x94, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	//L2 := [16]byte{0xd4, 0x56, 0x58, 0x4d, 0xd0, 0xe3, 0xe8, 0x4c, 0xc3, 0x16, 0x6e, 0x4b, 0x7f, 0xa2, 0x89, 0x0d}
	//L3 := [16]byte{0x79, 0xd2, 0x62, 0x21, 0xb8, 0x7b, 0x58, 0x4c, 0xd4, 0x2f, 0xbc, 0x4f, 0xfe, 0xa5, 0xde, 0x9a}
	//L4 := [16]byte{0x0e, 0x93, 0x69, 0x1a, 0x0c, 0xfc, 0x60, 0x40, 0x8b, 0x7b, 0x68, 0xf6, 0x6b, 0x51, 0x3c, 0x13}
	//fmt.Printf("L(input): %X\n", L(L1))
	//fmt.Printf("L(input): %X\n", L(L2))
	//fmt.Printf("L(input): %X\n", L(L3))
	//fmt.Printf("L(input): %X\n", L(L4))

	key := [32]byte{0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, 0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77,
		0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}
	//plainText := [16]byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x00, 0xff, 0xee, 0xdd, 0xcc, 0xbb, 0xaa, 0x99, 0x88}
	//keys := KeySchedule(key)
	//cipherText := Encrypt(plainText, keys)
	//fmt.Printf("cipherText: %X\n", cipherText)

	KeySchedule(key)
}
