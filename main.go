package main

import (
	"fmt"
	"spam_filter/utils"
)

const filesDir string = "files/"
const hamDir string = filesDir + "ham-anlern/"
const spamDir string = filesDir + "spam-anlern/"

func main() {
	fmt.Println("let's get all the provided ham files from the ham directory")
	hamFilesNameSlice := utils.ListFilesInDir(hamDir)
	fmt.Printf("there are %d ham files in the directory", len(hamFilesNameSlice))
}

type word struct{
	probability float32
}
