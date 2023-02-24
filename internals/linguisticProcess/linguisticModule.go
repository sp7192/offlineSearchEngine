package linguisticprocess

type LinguisticModule struct {
	converters []IStringConverter
}

func CreateLinguisticModule(converters ...IStringConverter) LinguisticModule {
	return LinguisticModule{converters: converters}
}

func (lm LinguisticModule) Convert(input string) string {
	ret := input
	for _, converter := range lm.converters {
		ret = converter.Convert(ret)
	}
	return ret
}
