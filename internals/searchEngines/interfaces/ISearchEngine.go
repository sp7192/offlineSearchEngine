package interfaces

import (
	searchEnginedata "OfflineSearchEngine/internals/searchEngines/models"
	"bufio"
)

type ISearchEngine interface {
	AddData(*bufio.Scanner, int)
	Search(string) (searchEnginedata.SearchResults, bool)
}
