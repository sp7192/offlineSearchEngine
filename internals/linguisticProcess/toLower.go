package linguisticprocess

import "strings"

type ToLower struct {
}

func (tl ToLower) Convert(input string) string {
	return strings.ToLower(input)
}
