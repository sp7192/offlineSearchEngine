package builder

import (
	idgenerator "OfflineSearchEngine/internals/idGenerator"
	linguisticprocess "OfflineSearchEngine/internals/linguisticProcess"
	"OfflineSearchEngine/internals/scanners"
	"OfflineSearchEngine/internals/searchEngines/interfaces"
	"OfflineSearchEngine/internals/searchEngines/invertedIndexEngine"
	"OfflineSearchEngine/internals/searchEngines/linearFastAddEngine"
	"OfflineSearchEngine/internals/searchEngines/linearFastSearchEngine"
	linearsortedengine "OfflineSearchEngine/internals/searchEngines/linearSortedEngine"
	linearsortedenginewithposting "OfflineSearchEngine/internals/searchEngines/linearSortedEngineWithPosting"
	texthandler "OfflineSearchEngine/internals/textHandler"
)

func NewSearchEngine(name string, capacity int, converter linguisticprocess.IStringConverter, scanner scanners.IScanner, idGenerator idgenerator.IIdGenerator) interfaces.ISearchEngine {
	textHandler := texthandler.NewTextHandler(converter, scanner, idGenerator)
	switch name {
	case "LinearFastAddEngine":
		return linearFastAddEngine.NewLinearFastAddEngine(capacity, textHandler)
	case "LinearFastSearchEngine":
		return linearFastSearchEngine.NewLinearFastSearchEngine(capacity, textHandler)
	case "LinearSortedEngine":
		return linearsortedengine.NewLinearSortedEngine(capacity, textHandler)
	case "LinearSortedEngineWithPosting":
		return linearsortedenginewithposting.NewLinearSortedEngineWithPosting(capacity, textHandler)
	case "InvertedIndex":
		return invertedIndexEngine.NewInvertedIndexEngine(capacity, textHandler)
	}
	return nil
}
