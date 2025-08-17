package utils

import (
	"os"
	"strings"
)

func ReadFileToLines(filename string) []string {
	contents, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(contents), "\n")
}
