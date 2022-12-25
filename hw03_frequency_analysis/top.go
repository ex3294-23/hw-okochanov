package hw03frequencyanalysis

import (
	"fmt"
	"regexp"
	"sort"
)

var r = regexp.MustCompile(`\s+`)

func Top10(text string) []string {
	if len(text) == 0 {
		return nil
	}

	wordsMap := make(map[string]int)

	words := r.Split(text, -1)
	for _, word := range words {
		wordsMap[word]++
	}

	result := make([]string, 0)
	for name := range wordsMap {
		result = append(result, name)
	}
	if len(wordsMap) <= len(result) {
		sort.Slice(result, func(i, j int) bool {
			first := result[i]
			second := result[j]

			if wordsMap[first] == wordsMap[second] {
				return first < second
			}

			return wordsMap[first] > wordsMap[second]
		})
	} else {
		fmt.Print("Need more words")
	}
	return result[:10]
}
