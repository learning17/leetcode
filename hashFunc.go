package main
import (
	"unsafe"
)

func RSHash(str []byte) uint32 {
	var b uint32 = 378551
	var a uint32 = 63689
	var hash uint32 = 0

	for i := 0; i < len(str); i++ {
		hash = hash*a + uint32(rune(str[i]))
		a = a * b
	}

	return hash
}

/* End Of RS Hash Function */

func JSHash(str []byte) uint32 {
	var hash uint32 = 1315423911

	for i := 0; i < len(str); i++ {
		hash ^= ((hash << 5) + uint32((rune(str[i]))) + (hash >> 2))
	}

	return hash
}

/* End Of JS Hash Function */

func PJWHash(str []byte) uint32 {
	uint32len := uint32(unsafe.Sizeof(*(*uint32)(unsafe.Pointer(nil))))
	var BitsInUnsignedInt uint32 = uint32(uint32len * 8)
	var ThreeQuarters uint32 = uint32((BitsInUnsignedInt * 3) / 4)
	var OneEighth uint32 = uint32(BitsInUnsignedInt / 8)
	var HighBits uint32 = uint32((0xFFFFFFFF) << (BitsInUnsignedInt - OneEighth))
	var hash uint32 = 0

	for i := 0; i < len(str); i++ {
		hash = (hash << OneEighth) + uint32(rune(str[i]))

		if test := hash & HighBits; test != 0 {
			hash = ((hash ^ (test >> ThreeQuarters)) & (^HighBits))
		}
	}

	return hash
}

/* End Of  P. J. Weinberger Hash Function */

func ELFHash(str []byte) uint32 {
	var hash uint32 = 0
	var x uint32 = 0

	for i := 0; i < len(str); i++ {
		hash = (hash << 4) + uint32(rune(str[i]))
		if x = hash & 0xF0000000; x != 0 {
			hash ^= (x >> 24)
		}
		hash &= ^x
	}

	return hash
}

/* End Of ELF Hash Function */

func BKDRHash(str []byte) uint32 {
	var seed uint32 = 131 /* 31 131 1313 13131 131313 etc.. */
	var hash uint32 = 0

	for i := 0; i < len(str); i++ {
		hash = (hash * seed) + uint32(rune(str[i]))
	}

	return hash
}

/* End Of BKDR Hash Function */

func SDBMHash(str []byte) uint32 {
	var hash uint32 = 0

	for i := 0; i < len(str); i++ {
		hash = uint32(rune(str[i])) + (hash << 6) + (hash << 16) - hash
	}

	return hash
}

/* End Of SDBM Hash Function */

func DJBHash(str []byte) uint32 {
	var hash uint32 = 5381

	for i := 0; i < len(str); i++ {
		hash = ((hash << 5) + hash) + uint32(rune(str[i]))
	}

	return hash
}

/* End Of DJB Hash Function */

func DEKHash(str []byte) uint32 {
	var hash uint32 = uint32(len(str))

	for i := 0; i < len(str); i++ {
		hash = ((hash << 5) ^ (hash >> 27)) ^ uint32(rune(str[i]))
	}
	return hash
}

/* End Of DEK Hash Function */

func BPHash(str []byte) uint32 {
	var hash uint32 = 0
	for i := 0; i < len(str); i++ {
		hash = (hash << 7) ^ uint32(rune(str[i]))
	}

	return hash
}

/* End Of BP Hash Function */

func FNVHash(str []byte) uint32 {
	var fnv_prime uint32 = 0x811C9DC5
	var hash uint32 = 0

	for i := 0; i < len(str); i++ {
		hash *= fnv_prime
		hash ^= uint32(rune(str[i]))
	}

	return hash
}

/* End Of FNV Hash Function */
