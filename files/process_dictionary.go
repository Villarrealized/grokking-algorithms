package files

import (
	"os"
	"strings"
)

func ProcessDictionary(filename string) (dictionary []string) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	for line := range strings.SplitSeq(string(bytes), "\n") {
		before, word, found := strings.Cut(line, "\t")
		if found {
			dictionary = append(dictionary, word)
		} else {
			dictionary = append(dictionary, strings.TrimSpace(before))
		}
	}

	return dictionary
}
