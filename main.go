package main

import (
	"fmt"
	"spam_filter/utils"
)

const (
	filesDir string = "files/"
	hamDir   string = filesDir + "ham-anlern/"
	spamDir  string = filesDir + "spam-anlern/"
)

type Word struct {
	spamOccurrences int16
	hamOccurrences  int16
}

func main() {
	fmt.Println("let's get all the provided ham files from the ham directory")
	hamFiles := utils.ListFilesInDir(hamDir)
	fmt.Printf("there are %d ham files in the directory", len(hamFiles))
	wordOccurrences := make(map[string]Word)
	for i := range hamFiles {
		wordSlice := utils.TurnFileIntoStringSlice(hamDir + hamFiles[i])
		for _, word := range wordSlice {
			curr := wordOccurrences[word]
			curr.hamOccurrences++
			wordOccurrences[word] = curr
		}
	}

	fmt.Println("let's get all the provided spam files from the spam directory")
	spamFiles := utils.ListFilesInDir(spamDir)
	fmt.Printf("there are %d spam files in the directory", len(spamFiles))
	for i := range spamFiles {
		wordSlice := utils.TurnFileIntoStringSlice(spamDir + spamFiles[i])
		for _, word := range wordSlice {
			curr := wordOccurrences[word]
			curr.spamOccurrences++
			wordOccurrences[word] = curr
		}
	}
	fmt.Println(wordOccurrences)
}
