package algorithms

import (
	"spam_filter/models"
	"spam_filter/utils"
)

func ASDF(filenames []string, wordOccurrences *map[string]models.Word, directory string) {
	for _, val := range filenames {
		wordMap := utils.TurnFileIntoStringMap(directory + val)
		for word := range wordMap {
			curr := (*wordOccurrences)[word]
			curr.HamOccurrences++
			(*wordOccurrences)[word] = curr
		}
	}
}
