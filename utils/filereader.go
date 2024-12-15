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

func TurnFileIntoStringMap(filename string) map[string]struct{} {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	m := make(map[string]struct{})
	words := strings.Split(string(content), " ")
	for _, val := range words {
		m[val] = struct{}{}
	}
	return m
}
