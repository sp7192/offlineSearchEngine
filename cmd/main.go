package main

import (
	"OfflineSearchEngine/api"
	idgenerator "OfflineSearchEngine/internals/idGenerator"
	linguisticprocess "OfflineSearchEngine/internals/linguisticProcess"
	"OfflineSearchEngine/internals/scanners"
	engineBuilder "OfflineSearchEngine/internals/searchEngines/builder"
	"bufio"
	"log"
	"strings"
)

func main() {
	lm := linguisticprocess.NewLinguisticModule(&linguisticprocess.CheckStopWord{},
		&linguisticprocess.PunctuationRemover{},
		&linguisticprocess.ToLower{})

	reader, err := scanners.NewFileReaderClosers("../data")
	if err != nil {
		log.Fatalf("Error in reading : %s\n", err.Error())
	}
	scanner, err := scanners.NewFolderScanner(reader)
	if err != nil {
		log.Fatalf("Error in Scanner : %s\n", err.Error())
	}

	idGenerator := idgenerator.NewIdGenerator()

	se := engineBuilder.NewSearchEngine("InvertedIndex", 500, lm, &scanner, &idGenerator)
	sc := bufio.NewScanner(strings.NewReader("test query"))
	sc.Split(bufio.ScanWords)
	se.AddData(sc, 1)
	server := api.NewServer(se)
	server.Run(":8080")
}
