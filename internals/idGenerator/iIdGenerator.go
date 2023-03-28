package idgenerator

type IIdGenerator interface {
	AddFilename(fileName string)
	GetFilename(id int) (string, bool)
	RemoveFilename(id int)
}
