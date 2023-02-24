package linguisticprocess

import "strings"

type PunctuationRemover struct {
}

func (st PunctuationRemover) Convert(input string) string {
	// TODO : to add these in a config file
	punctuations := []string{".", ",", ";", ":", "!", "?", `"`, `'`, ")", "]", "}", ">", "(", "[", "{", "<"}

	for _, v := range punctuations {
		input = strings.Trim(input, v)
	}
	return input
}
