package texthandler

import (
	idgenerator "OfflineSearchEngine/internals/idGenerator"
	linguisticprocess "OfflineSearchEngine/internals/linguisticProcess"
	"OfflineSearchEngine/internals/scanners"
)

type TextHandler struct {
	StringConverter linguisticprocess.IStringConverter
	Scanner         scanners.IScanner
	IdGenerator     idgenerator.IIdGenerator
}

func NewTextHandler(converter linguisticprocess.IStringConverter, scanner scanners.IScanner, idGenerator idgenerator.IIdGenerator) TextHandler {
	return TextHandler{
		StringConverter: converter,
		Scanner:         scanner,
		IdGenerator:     idGenerator,
	}
}
