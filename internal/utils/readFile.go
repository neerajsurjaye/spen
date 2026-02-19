package utils

import (
	"os"
)

func ReadShader(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil{
		return "", err
	}
	
	shader := string(bytes) + "\x00"
	return shader, err
}