package scanners

type IFolderScanner interface {
	GetFileNames(string) ([]string, error)
}
