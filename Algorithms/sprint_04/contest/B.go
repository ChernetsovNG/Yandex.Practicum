package main

import (
	"fmt"
)

type Collision struct {
	str1, str2   string
	hash1, hash2 int
}

func newCollision(str1, str2 string, hash1, hash2 int) Collision {
	return Collision{str1, str2, hash1, hash2}
}

// основание, по которому вычисляется хеш
const a = 1000

// модуль m
const m = 123_987_123

func main() {
	n := 16
	// генерируем строки длины n, пока не найдём строки, вызывающие коллизии
	hashesStrings := make(map[int]string)
	generateStrings(n, "", hashesStrings)
}

// генерация строк длины n
func generateStrings(n int, prefix string, hashes map[int]string) {
	if n == 0 {
		hash := polynomialHash(a, m, prefix)
		wasString, ok := hashes[hash]
		if ok { // нашли коллизию
			fmt.Printf("collision: str1 = %s, str2 = %s\n", prefix, wasString)
			hashes[hash] = prefix
		}
		hashes[hash] = prefix
		return
	}
	generateStrings(n-1, prefix+"a", hashes)
	generateStrings(n-1, prefix+"b", hashes)
	generateStrings(n-1, prefix+"c", hashes)
	generateStrings(n-1, prefix+"d", hashes)
	generateStrings(n-1, prefix+"e", hashes)
	generateStrings(n-1, prefix+"f", hashes)
	generateStrings(n-1, prefix+"g", hashes)
	generateStrings(n-1, prefix+"h", hashes)
	generateStrings(n-1, prefix+"i", hashes)
	generateStrings(n-1, prefix+"j", hashes)
	generateStrings(n-1, prefix+"k", hashes)
	generateStrings(n-1, prefix+"l", hashes)
	generateStrings(n-1, prefix+"m", hashes)
	generateStrings(n-1, prefix+"n", hashes)
	generateStrings(n-1, prefix+"o", hashes)
	generateStrings(n-1, prefix+"p", hashes)
	generateStrings(n-1, prefix+"q", hashes)
	generateStrings(n-1, prefix+"r", hashes)
	generateStrings(n-1, prefix+"s", hashes)
	generateStrings(n-1, prefix+"t", hashes)
	generateStrings(n-1, prefix+"u", hashes)
	generateStrings(n-1, prefix+"v", hashes)
	generateStrings(n-1, prefix+"w", hashes)
	generateStrings(n-1, prefix+"x", hashes)
	generateStrings(n-1, prefix+"y", hashes)
	generateStrings(n-1, prefix+"z", hashes)
}

func polynomialHash(a, m int, str string) int {
	var hash, symbolCode int
	var symbol rune
	n := len(str)
	symbols := []rune(str)
	hash = 0
	for i := 0; i < n; i++ {
		symbol = symbols[i]
		symbolCode = int(symbol)
		hash = (hash + symbolCode*powerByModule(a, n-1-i, m)) % m
	}
	return hash
}

func powerByModule(b, e, m int) int {
	result := 1
	if 1&e == 1 {
		result = b
	}
	for true {
		if e == 0 {
			break
		}
		e >>= 1
		b = (b * b) % m
		if e&1 == 1 {
			result = (result * b) % m
		}
	}
	return result
}
