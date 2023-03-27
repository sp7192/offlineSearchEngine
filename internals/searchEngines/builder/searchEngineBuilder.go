package builder

import (
	linguisticprocess "OfflineSearchEngine/internals/linguisticProcess"
	"OfflineSearchEngine/internals/scanners"
	"OfflineSearchEngine/internals/searchEngines/interfaces"
	"OfflineSearchEngine/internals/searchEngines/invertedIndexEngine"
	"OfflineSearchEngine/internals/searchEngines/linearFastAddEngine"
	"OfflineSearchEngine/internals/searchEngines/linearFastSearchEngine"
	linearsortedengine "OfflineSearchEngine/internals/searchEngines/linearSortedEngine"
	linearsortedenginewithposting "OfflineSearchEngine/internals/searchEngines/linearSortedEngineWithPosting"
)

func NewSearchEngine(name string, capacity int, converter linguisticprocess.IStringConverter, scanner scanners.IScanner) interfaces.ISearchEngine {
	switch name {
	case "LinearFastAddEngine":
		return linearFastAddEngine.NewLinearFastAddEngine(capacity, converter, scanner)
	case "LinearFastSearchEngine":
		return linearFastSearchEngine.NewLinearFastSearchEngine(capacity, converter, scanner)
	case "LinearSortedEngine":
		return linearsortedengine.NewLinearSortedEngine(capacity, converter, scanner)
	case "LinearSortedEngineWithPosting":
		return linearsortedenginewithposting.NewLinearSortedEngineWithPosting(capacity, converter, scanner)
	case "InvertedIndex":
		return invertedIndexEngine.NewInvertedIndexEngine(capacity, converter, scanner)
	}
	return nil
}
