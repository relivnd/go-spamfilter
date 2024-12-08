package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const hamDirectory string = "files/ham-anlern/"

func main() {
	var fileArray = readDirectory(hamDirectory)
	var stringArray = readFile(hamDirectory + fileArray[0].Name())
	fmt.Println(stringArray)
}

func readFile(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(content), " ")
}

func readDirectory(directoy string) []os.DirEntry {
	content, err := os.ReadDir(directoy)
	if err != nil {
		log.Fatal(err)
	}
	return content
}
