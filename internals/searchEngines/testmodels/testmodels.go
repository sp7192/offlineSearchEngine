package testmodels

import "OfflineSearchEngine/internals/searchEngines/models"

type DocData struct {
	Text  string
	DocId int
}

type SearchInputData struct {
	Inputs []DocData
	Query  string
}

type AddDataInput struct {
	TestName  string
	InputData []DocData
}

type SearchOutput struct {
	Output models.SearchResults
	Ok     bool
}
