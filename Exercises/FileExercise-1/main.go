package main

import (
	"fmt"
	"log"
	"os"
)

const (
	savePath = "source"
)

type FileManager struct {
	file *os.File
}

func NewFileManager(
	file *os.File,
) *FileManager {
	return &FileManager{
		file: file,
	}
}

func main () {
	var file *os.File

	// Initialize an instance of FileManager
	manager := NewFileManager(file)
	
	newFile, err := manager.CreateFile("test-1") 
	if err != nil {
		log.Fatal(err)
	}
	newFile.Stat()
}

func (fm *FileManager) CreateFile(filename string) (*os.File, error) {
	createdFile, err := os.Create(savePath + "/" + filename + ".txt")
	if err != nil {
		return nil, fmt.Errorf("File could not be created in source directory")
	}

	return createdFile, nil
}

func (fm *FileManager) RenameFile(filename, newFilename string) error {
	panic("unimplemented")
}

func (fm *FileManager) WriteToFile(file os.File, content []byte) {
	panic("unimplemented")
}

func (fm *FileManager) ReadToFile(file os.File, content []byte) {
	panic("unimplemented")
}

func (fm *FileManager) DeleteFile(filename string) error {
	panic("unimplemented")
}

func (fm *FileManager) CopyFile(filename, sourcePath, destinationPath  string) (string, error) {
	panic("unimplemented")
}