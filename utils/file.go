package utils

type File struct {
	Name     string
	Contents string
}

func NewFile(name, contents string) *File {
	return &File{
		Name:     name,
		Contents: contents,
	}
}
