package common

import (
	"os"
	"sync/atomic"
)

type Source interface {
	Open() error
	Next() (string, bool)
}

type FileSystemSource struct {
	inputDir  string
	outputDir string

	files []os.DirEntry
	index atomic.Int32
}

func (f *FileSystemSource) Open() error {
	// Check input directory
	inputDirInfo, err := os.Stat(f.inputDir)
	if err != nil {
		return err
	}
	if !inputDirInfo.IsDir() {
		return os.ErrInvalid
	}

	// Check output directory
	outputDirInfo, err := os.Stat(f.outputDir)
	if err != nil {
		return err
	}
	if !outputDirInfo.IsDir() {
		return os.ErrInvalid
	}

	//read input directory
	files, err := os.ReadDir(f.inputDir)
	if err != nil {
		return err
	}
	f.files = files
	return nil
}

func (f *FileSystemSource) Next() (string, bool) {
	if f.index.Load() >= int32(len(f.files)) {
		return "", false
	}
	fileName := f.files[f.index.Load()].Name()
	f.index.Add(1)
	return fileName, f.index.Load() < int32(len(f.files))
}
