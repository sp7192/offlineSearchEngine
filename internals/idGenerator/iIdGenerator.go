package idgenerator

type IIdGenerator interface {
	AddFilename(fileName string) int
	GetFilename(id int) (string, bool)
	RemoveFilename(id int)
}
