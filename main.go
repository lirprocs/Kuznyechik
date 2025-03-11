package Kuznyechik

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

func KeySchedule(masterKey [32]byte) [10][16]byte {
	var keys [10][16]byte
	var C [32][16]byte

	copy(keys[0][:], masterKey[:16])
	copy(keys[1][:], masterKey[16:])

	fmt.Printf("K1: %x\n", keys[0])
	fmt.Printf("K2: %x\n", keys[1])

	for i := 0; i < 32; i++ {
		C[i] = L(Vec128(i + 1))
	}

	// Генерация ключей
	for i := 1; i < 5; i++ {
		for j := 0; j < 8; j++ {
			index := 8*(i-1) + j
			keys[2*i], keys[2*i+1] = F(C[index], keys[2*i-2], keys[2*i-1])
			fmt.Printf("C[%d]: %x\n", index+1, C[index])
		}
		fmt.Printf("K%d: %x\n", 2*i+1, keys[2*i])
		fmt.Printf("K%d: %x\n", 2*i+2, keys[2*i+1])
	}
	return keys
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
	inputS := [16]byte{0xff, 0xee, 0xdd, 0xcc, 0xbb, 0xaa, 0x99, 0x88, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x00}
	fmt.Printf("S(input): %X\n", S(inputS))

	// Проверка R-функции
	inputR := [16]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00}
	fmt.Printf("R(input): %X\n", R(inputR))

	// Проверка L-функции
	inputL := [16]byte{0x64, 0xa5, 0x94, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	fmt.Printf("L(input): %X\n", L(inputL))

	key := [32]byte{0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, 0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77,
		0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}
	plainText := [16]byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x00, 0xff, 0xee, 0xdd, 0xcc, 0xbb, 0xaa, 0x99, 0x88}
	keys := KeySchedule(key)
	cipherText := Encrypt(plainText, keys)
	fmt.Printf("cipherText: %X\n", cipherText)
}
