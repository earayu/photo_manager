package common

import (
	"image"
	"os"
	"sync/atomic"
)

type Source interface {
	Open() error
	Next() (string, bool)
}

type FileSystemSource struct {
	InputDir  string
	OutputDir string

	Files []os.DirEntry
	index atomic.Int32

	//filter out some files based on file name, e.g. filter out all files that end with ".mov" or ".mp4"
	SourceFilter func(fileName string) bool

	//processed images
	ProcessedImages []*image.Image
}

func (f *FileSystemSource) Open() error {
	// Check input directory
	inputDirInfo, err := os.Stat(f.InputDir)
	if err != nil {
		return err
	}
	if !inputDirInfo.IsDir() {
		return os.ErrInvalid
	}

	// Check output directory
	outputDirInfo, err := os.Stat(f.OutputDir)
	if err != nil {
		return err
	}
	if !outputDirInfo.IsDir() {
		return os.ErrInvalid
	}

	//read input directory
	files, err := os.ReadDir(f.InputDir)
	if err != nil {
		return err
	}
	f.Files = files
	return nil
}

func (f *FileSystemSource) Next() (string, bool) {
	if f.index.Load() >= int32(len(f.Files)) {
		return "", false
	}
	fileName := f.Files[f.index.Load()].Name()
	f.index.Add(1)
	if f.SourceFilter != nil && !f.SourceFilter(fileName) {
		return f.Next()
	}
	return fileName, f.index.Load() < int32(len(f.Files))
}
