package main

import (
	"fmt"
	"spam_filter/utils"
)

const (
	filesDir   string  = "files/"
	hamDir     string  = filesDir + "ham-anlern/"
	spamDir    string  = filesDir + "spam-anlern/"
	hamCalDir  string  = filesDir + "ham-kallibrierung/"
	spamCalDir string  = filesDir + "spam-kallibrierung/"
	alpha      float32 = 0.00001
)

type Word struct {
	spamOccurrences int16
	hamOccurrences  int16
}

type Probability struct {
	hamProbability  float32
	spamProbability float32
}

func main() {
	fmt.Println("let's get all the provided ham files from the ham directory")
	hamFiles := utils.ListFilesInDir(hamDir)
	numberOfHamFiles := len(hamFiles)
	fmt.Printf("there are %d ham files in the directory\n", numberOfHamFiles)
	wordOccurrences := make(map[string]Word)
	for i := range hamFiles {
		wordMap := utils.TurnFileIntoStringMap(hamDir + hamFiles[i])
		for word := range wordMap {
			curr := wordOccurrences[word]
			curr.hamOccurrences++
			wordOccurrences[word] = curr
		}
	}

	fmt.Println("let's get all the provided spam files from the spam directory")
	spamFiles := utils.ListFilesInDir(spamDir)
	numberOfSpamFiles := len(spamFiles)
	fmt.Printf("there are %d spam files in the directory\n", numberOfSpamFiles)
	for i := range spamFiles {
		wordMap := utils.TurnFileIntoStringMap(spamDir + spamFiles[i])
		for word := range wordMap {
			curr := wordOccurrences[word]
			curr.spamOccurrences++
			wordOccurrences[word] = curr
		}
	}

	totalNumberOfFiles := numberOfHamFiles + numberOfSpamFiles
	fmt.Printf("a total number of %d mail have been analysed\n", totalNumberOfFiles)
	wordProbabilities := make(map[string]Probability)
	for k, val := range wordOccurrences {
		i := Probability{}
		i.hamProbability = float32(val.hamOccurrences) / float32(numberOfHamFiles)
		if i.hamProbability == 0 {
			i.hamProbability = alpha
		}
		i.spamProbability = float32(val.spamOccurrences) / float32(numberOfSpamFiles)
		if i.spamProbability == 0 {
			i.spamProbability = alpha
		}
		wordProbabilities[k] = i
	}

	fmt.Println("start calibration")
	numberOfSpamClassifiedMails := 0
	numberOfHamClassifiedMails := 0

	fmt.Println("first the ham files")
	hamCalFiles := utils.ListFilesInDir(hamCalDir)
	numberOfHamCalFiles := len(hamCalFiles)
	fmt.Printf("there are %d ham files in the directory\n", numberOfHamCalFiles)

	for i := range hamCalFiles {
		wordMap := utils.TurnFileIntoStringMap(hamCalDir + hamCalFiles[i])
		b := 1.0
		c := 1.0
		wordSpamProbability := 0.0
		mailSpamProbability := 0.0
		for word := range wordMap {
			wordSpamProbability = float64(wordProbabilities[word].spamProbability) / (float64(wordProbabilities[word].spamProbability) + float64(wordProbabilities[word].hamProbability))
			b = b * wordSpamProbability
			c = c * (1.0 - wordSpamProbability)
		} //TODO: the combination of the probabilities is missing
		mailSpamProbability = b / (b + c)
		if mailSpamProbability > 0.5 {
			numberOfSpamClassifiedMails++
			fmt.Printf("SPAM: the mail with name %s has a spam probability of %f\n", hamCalFiles[i], mailSpamProbability)
		} else {
			numberOfHamClassifiedMails++
			fmt.Printf("HAM: the mail with name %s has a spam probability of %f\n", hamCalFiles[i], mailSpamProbability)
		}
	}
	fmt.Printf("there were %d SPAM mails and %d HAM mails. The total number of mails were %d\n", numberOfSpamClassifiedMails, numberOfHamClassifiedMails, numberOfHamCalFiles)

}
