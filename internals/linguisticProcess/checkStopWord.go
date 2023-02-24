package linguisticprocess

type CheckStopWord struct {
}

func (ch CheckStopWord) Convert(input string) string {
	// TODO : get from config in future
	stopWords := []string{"of", "the", "in", "on", "a", "an"}
	for _, v := range stopWords {
		if input == v {
			return ""
		}
	}
	return input
}
