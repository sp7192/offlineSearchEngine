package idgenerator

type IdGenerator struct {
	filenamesId map[int]string
	lastId      int
}

func NewIdGenerator() IdGenerator {
	return IdGenerator{
		filenamesId: make(map[int]string),
		lastId:      0,
	}
}

func (idg *IdGenerator) AddFilename(fileName string) {
	idg.filenamesId[idg.lastId] = fileName
	idg.lastId++
}

func (idg *IdGenerator) GetFilename(id int) (string, bool) {
	ret, ok := idg.filenamesId[id]
	return ret, ok
}

func (idg *IdGenerator) RemoveFilename(id int) {
	delete(idg.filenamesId, id)
}
