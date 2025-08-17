package utils

import (
	"os"
	"strings"
)

func MustReadFile(filename string) []byte {
	contents, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return contents
}

func ReadFileToLines(filename string) []string {
	contents, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(contents), "\n")
}
