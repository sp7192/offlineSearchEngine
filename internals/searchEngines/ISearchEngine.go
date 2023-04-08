package searchEngines

import (
	"OfflineSearchEngine/internals/scanners"
	searchEnginedata "OfflineSearchEngine/internals/searchEngines/models"
)

type ISearchEngine interface {
	AddData(scanners.IScanner, int)
	Search(string) (searchEnginedata.SearchResults, bool)
}
