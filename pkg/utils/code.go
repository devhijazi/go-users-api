package utils

import gonanoid "github.com/matoous/go-nanoid/v2"

func GenerateCode(size int) string {
	numberValues := "1234567890"

	code, _ := gonanoid.Generate(numberValues, size)

	return code
}
