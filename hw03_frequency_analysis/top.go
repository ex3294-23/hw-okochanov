package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type wordsAndCountStruct struct {
	word  string
	count int
}

type WordsAndCountSlice []wordsAndCountStruct

func (s WordsAndCountSlice) Len() int {
	return len(s)
}

func (s WordsAndCountSlice) Less(i, j int) bool {
	if s[i].count == s[j].count {
		return s[i].word < s[j].word
	}
	return s[i].count > s[j].count
}

func (s WordsAndCountSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func Top10(str string) []string {
	words := getSortedSliceFromStr(str)

	wcSlice := getWordsAndCountSlice(words)

	sort.Sort(wcSlice)

	return formatResult(wcSlice)
}

func getWordsAndCountSlice(words []string) WordsAndCountSlice {
	wcSlice := WordsAndCountSlice{}

	currentWord := ""
	currentCount := 0
	for _, value := range words {
		if value != currentWord {
			if currentWord != "" {
				wcSlice = append(wcSlice, wordsAndCountStruct{currentWord, currentCount})
			}

			currentWord = value
			currentCount = 1
		} else {
			currentCount++
		}
	}
	if currentWord != "" {
		wcSlice = append(wcSlice, wordsAndCountStruct{currentWord, currentCount})
	}

	return wcSlice
}

func getSortedSliceFromStr(str string) []string {
	words := strings.Fields(str)

	sort.Slice(words, func(i, j int) bool {
		return words[i] < words[j]
	})
	return words
}

func formatResult(wcSlice WordsAndCountSlice) []string {
	res := []string{}

	for _, wordsStruct := range wcSlice {
		res = append(res, wordsStruct.word)
	}

	limit := 10
	if len(res) < 10 {
		limit = len(res)
	}

	return res[0:limit]
}
