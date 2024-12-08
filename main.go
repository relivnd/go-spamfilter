package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var file = "files/ham-anlern/0126.d002ec3f8a9aff31258bf03d62abdafa"

func main() {

	array := readFile(file)
	fmt.Println(array[2])
}

func readFile(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(content), " ")
}
