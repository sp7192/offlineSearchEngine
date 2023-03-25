package main

import (
	"OfflineSearchEngine/api"
	linguisticprocess "OfflineSearchEngine/internals/linguisticProcess"
	engineBuilder "OfflineSearchEngine/internals/searchEngines/builder"
	"bufio"
	"strings"
)

func main() {
	lm := linguisticprocess.CreateLinguisticModule(&linguisticprocess.CheckStopWord{}, &linguisticprocess.PunctuationRemover{}, &linguisticprocess.ToLower{})
	se := engineBuilder.NewSearchEngine("InvertedIndex", 500, lm)
	server := api.NewServer(se)
	sc := bufio.NewScanner(strings.NewReader("test query"))
	sc.Split(bufio.ScanWords)
	se.AddData(sc, 1)
	server.Run(":8080")
}	
