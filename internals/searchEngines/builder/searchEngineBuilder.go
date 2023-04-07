package builder

import (
	linguisticprocess "OfflineSearchEngine/internals/linguisticProcess"
	"OfflineSearchEngine/internals/searchEngines"
	"OfflineSearchEngine/internals/searchEngines/invertedIndexEngine"
	"OfflineSearchEngine/internals/searchEngines/linearFastAddEngine"
	"OfflineSearchEngine/internals/searchEngines/linearFastSearchEngine"
	linearsortedengine "OfflineSearchEngine/internals/searchEngines/linearSortedEngine"
	linearsortedenginewithposting "OfflineSearchEngine/internals/searchEngines/linearSortedEngineWithPosting"
)

func NewSearchEngine(name string, capacity int, converter linguisticprocess.IStringConverter) searchEngines.ISearchEngine {
	switch name {
	case "LinearFastAddEngine":
		return linearFastAddEngine.NewLinearFastAddEngine(capacity, converter)
	case "LinearFastSearchEngine":
		return linearFastSearchEngine.NewLinearFastSearchEngine(capacity, converter)
	case "LinearSortedEngine":
		return linearsortedengine.NewLinearSortedEngine(capacity, converter)
	case "LinearSortedEngineWithPosting":
		return linearsortedenginewithposting.NewLinearSortedEngineWithPosting(capacity, converter)
	case "InvertedIndex":
		return invertedIndexEngine.NewInvertedIndexEngine(capacity, converter)
	}
	return nil
}
