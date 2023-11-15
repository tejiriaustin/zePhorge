package os

import "os"

type (
	FileSystem struct{}

	FileSystemInterface interface {
		Create(name string) (*os.File, error)
		CreateTemp(dir string, pattern string) (*os.File, error)
	}
)

var _ FileSystemInterface = (*FileSystem)(nil)

func (f FileSystem) Create(name string) (*os.File, error) {
	return os.Create(name)
}

func (f FileSystem) CreateTemp(dir string, pattern string) (*os.File, error) {
	return os.CreateTemp(dir, pattern)
}
