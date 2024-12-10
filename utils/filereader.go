package utils

import (
	"log"
	"os"
	"strings"
)

func ListFilesInDir(directoy string) []string {
	var fileNames []string
	content, err := os.ReadDir(directoy)
	if err != nil {
		log.Fatal(err)
	}
	for _, val := range content {
		fileNames = append(fileNames, val.Name())
	}
	return fileNames
}

func TurnFileIntoStringSlice(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(content), " ")
}
