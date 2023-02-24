package builder

import (
	linguisticprocess "OfflineSearchEngine/internals/linguisticProcess"
	"OfflineSearchEngine/internals/searchEngines/interfaces"
	"OfflineSearchEngine/internals/searchEngines/linearFastAddEngine"
	"OfflineSearchEngine/internals/searchEngines/linearFastSearchEngine"
)

func NewSearchEngine(name string, capacity int, converter linguisticprocess.IStringConverter) interfaces.ISearchEngine {
	switch name {
	case "LinearFastAddEngine":
		return linearFastAddEngine.NewLinearFastAddEngine(capacity, converter)
	case "LinearFastSearchEngine":
		return linearFastSearchEngine.NewLinearFastSearchEngine(capacity, converter)
	}
	return nil
}
