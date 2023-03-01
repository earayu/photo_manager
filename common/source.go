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
	InputDir  string
	OutputDir string

	Files []os.DirEntry
	index atomic.Int32
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
	return fileName, f.index.Load() < int32(len(f.Files))
}
