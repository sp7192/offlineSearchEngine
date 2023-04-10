package main

import (
	"OfflineSearchEngine/api"
	"OfflineSearchEngine/configs"
	idgenerator "OfflineSearchEngine/internals/idGenerator"
	linguisticprocess "OfflineSearchEngine/internals/linguisticProcess"
	"OfflineSearchEngine/internals/searchEngines/builder"
	"log"
)

func main() {
	config, err := configs.LoadConfigs("../configs")
	if err != nil {
		log.Fatalf("could not load config file, err is : %s\n", err.Error())
	}

	dbConfig, err := configs.LoadDbConfigs("../configs")

	lm := linguisticprocess.NewLinguisticModule(&linguisticprocess.CheckStopWord{},
		&linguisticprocess.PunctuationRemover{},
		&linguisticprocess.ToLower{})

	idGenerator := idgenerator.NewIdGenerator()
	se := builder.NewSearchEngine(config.EngineType, 500, lm)

	server := api.NewServer(se, &idGenerator, &config, &dbConfig)
	err = server.LoadDirectoryFiles("../data")
	if err != nil {
		log.Fatalf("could not load files, err is : %s\n", err.Error())
	}
	server.Run(":8080")
}
