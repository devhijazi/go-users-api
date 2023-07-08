package utils

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"
)

func StringToNumber(str string) int {
	num, _ := strconv.Atoi(str)

	return num
}

func StringRandomHex(size int) string {
	bytes := make([]byte, size)

	rand.Read(bytes)

	return hex.EncodeToString(bytes)
}
