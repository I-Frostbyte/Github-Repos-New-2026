package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	savePath = "source"

	copyPath = "destination"
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
	
	// Create file with filename.
	newFile, err := manager.CreateFile("test-7") 
	if err != nil {
		log.Fatal(err)
	}
	
	// Write content to file. In this case, text.
	if err := manager.WriteToFile(newFile, []byte("Copy file test.")); err != nil {
		log.Fatal(err)
	}

	// Read content from file.
	fileContent, err := manager.ReadFromFile(newFile)
	if err != nil {
		log.Fatal(err)
	}

	// Read metadata from file to get filename.
	metadata, err := newFile.Stat()
	if err != nil {
		log.Fatal(err)
	}

	fileName := metadata.Name()

	// Print content read from file.
	fmt.Printf("Content read from %s: %s\n", fileName, fileContent)

	// Extracting the absolute path from file and renaming it.
	oldPath, newPath, err := RenameFilePath(newFile)
	if err != nil {
		log.Fatal(err)
	}

	// Rename file
	if err := manager.RenameFile(oldPath, newPath); err != nil {
		log.Fatal(err)
	}

	// Setting up source and destination paths.
	sourcePath := savePath + "/" + "test-7-renamed" + ".txt"
	destinationPath := copyPath + "/" + "test-7-renamed" + ".txt"

	// copying file from source to destination
	if err := manager.CopyFile(sourcePath, destinationPath); err != nil {
		log.Fatal(err)
	}
}

func (fm *FileManager) CreateFile(filename string) (*os.File, error) {
	// CreatedFile is stored in the savePath and saved with the filename passed.
	createdFile, err := os.Create(savePath + "/" + filename + ".txt")
	if err != nil {
		return nil, fmt.Errorf("File could not be created in source directory")
	}

	return createdFile, nil
}

func (fm *FileManager) RenameFile(oldfilePath, newFilePath string) error {
	if err := os.Rename(oldfilePath, newFilePath); err != nil {
		return fmt.Errorf("Could not rename")
	}

	return nil
}

func (fm *FileManager) WriteToFile(file *os.File, content []byte) error {
	// Writing content to file in bytes.
	_, err := file.Write(content)
	if err != nil {
		return fmt.Errorf("Content length doesn't match number of bytes written.")
	}

	return nil
}

func (fm *FileManager) ReadFromFile(file *os.File) (string, error) {
	// SeekStart resets the cursor on the file to avoid EOD reading.
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		return "", fmt.Errorf("Could not reset cursor")
	}

	// Reading from file in content bytes.
	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("Could not read content from file")
	}

	contentString := string(content)

	if contentString == "" {
		return "", fmt.Errorf("Bytes not properly handled.")
	}

	return contentString, nil
}

func (fm *FileManager) DeleteFile(filename string) error {
	panic("unimplemented")
}

func (fm *FileManager) CopyFile(sourcePath, destinationPath  string) error {
	// Open the source file
	source, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Could not open source file")
	}
	defer source.Close()

	// Create the destination file
	destination, err := os.Create(destinationPath)
	if err != nil {
		return fmt.Errorf("Could not create the destination file")
	}
	defer destination.Close()

	// Copy the contents
	_, err = io.Copy(destination, source)
	if err != nil {
		return fmt.Errorf("Could not copy contents of file from source to destination")
	}

	return nil
}

func RenameFilePath (file *os.File) (string, string, error) {
	// Extracting the absolute path of the file.
	oldPath, err := filepath.Abs(file.Name())
	if err != nil {
		return "", "", fmt.Errorf("Could not parse absolute path of file")
	}

	// Split the path into directory and filename
	dir := filepath.Dir(oldPath)
	filename := filepath.Base(oldPath)

	// Separate the extension from the filename
	ext := filepath.Ext(filename)

	// Remove the extension from the filename
	name := strings.TrimSuffix(filename, ext)

	// Renaming the filename to the new name
	newFilename := name + "-renamed" + ext

	// Build the new absolute path
	newPath := filepath.Join(dir, newFilename)

	return oldPath, newPath, nil
}