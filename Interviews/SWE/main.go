package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

// file struct holds the object properties.
type file struct {
	ID string
	Name string
	Size int64
	UploadedAt time.Time
	Contents []byte
}

// FileStorage struct containing an array of files.
type FileStorage struct {
	uploadedFiles []file
}


// FileStorage constructor to allow for initialization of a FileStorage element independent of FileStorage itself.
func NewFileStorage(
	uploadedFiles []file,
) *FileStorage {
	return &FileStorage{
		uploadedFiles: uploadedFiles,
	}
}

func main() {
	var ctx context.Context

	// Initialize an instance of the FileStorage that can have methods run on it without affecting the FileStorage struct.
	storage := NewFileStorage([]file{})

	// Running the Upload method on the FileStorage instance.
	fileID, err := storage.Upload(ctx, "./home/frostbyte/workspace/first-file.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("New File ID: %v", fileID)

	meta, err := storage.GetMetadata(ctx, "124898343f75934")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File Metadata %v", meta)

	if err := storage.Download(ctx, fileID, "./home/frostbyte/workspace/profile.jpg"); err != nil {
		log.Fatal(err)
	}

	storage.Delete(ctx, fileID)
	
}

// Upload is the first method of the FileStorage object.
// It takes a filepath, collects the file from there and
// returns a generated uuid as the File ID.
func (fs *FileStorage) Upload(ctx context.Context, filePath string) (string, error) {
	// os.Open extracts the file.
	uploadedFile, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("Could not read file from filepath")
	}

	// os.File.Stat() returns the metadata of the file.
	info, err := uploadedFile.Stat()
	if err != nil {
		return "", fmt.Errorf("Uploaded File has no metadata")
	}

	// io.ReadAll collects the content of the file.
	contents, err := io.ReadAll(uploadedFile)
	if err != nil {
		return "", fmt.Errorf("Uploaded file cannot be read")
	}

	newFile := file{
		ID: uuid.NewString(),
		Name: info.Name(),
		Size: info.Size(),
		UploadedAt: info.ModTime(),
		Contents: contents,
	}

	fs.uploadedFiles = append(fs.uploadedFiles, newFile)

	return newFile.ID, nil
}

// GetMetaData is the second method of the FileStorage object.
// It takes the FileID, searches for the file that matches that ID and
// returns that File along with all the metadata
func (fs *FileStorage) GetMetadata(ctx context.Context, fileID string) (file, error) {
	var metadata file
	for _, file := range fs.uploadedFiles {
		if file.ID == fileID {
			metadata = file
		}
	}

	return metadata, nil
}

// Download is the third method of the FileStorage object.
// It takes the FileID & destinationPath.
// First it finds the file among the array, extracts its contents
// It then creates the destination file in the destination path
// and writes the contents of the original file into the destination file.
func (fs *FileStorage) Download(ctx context.Context, fileID, destinationPath string) error {
	var storedFile file
	for _, file := range fs.uploadedFiles {
		if file.ID == fileID {
			storedFile = file
		}
	}

	//os.Create creates a an empty file in the destinationPath
	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return fmt.Errorf("Could not create file in destination Path")
	}

	// Different ways to write into the file.
	// Method 1, write directly
	destinationFile.Write(storedFile.Contents)

	// Method 2, wrap the bytes in a reader and then copy into the destinationfile
	// copyFile := bytes.NewReader(storedFile.Contents)
	// io.Copy(destinationFile, copyFile)

	return nil
}

func (fs *FileStorage) Delete(ctx context.Context, fileID string) error {
	var ErrNotFound = errors.New("file not found")
	for i, file := range fs.uploadedFiles {
		if file.ID == fileID {
			fs.uploadedFiles = append(
				fs.uploadedFiles[:i],
				fs.uploadedFiles[i+1:]...,
			)
			
			return nil
		}
	}

	return ErrNotFound
}

/*

func (fs *FileStorage) Upload(filePath string) (string, error)          // store the file, return a unique fileID
func (fs *FileStorage) Download(fileID, destinationPath string) error   // reconstruct the file at destination
func (fs *FileStorage) Delete(fileID string) error                      // remove the file
func (fs *FileStorage) GetMetadata(fileID string) (Metadata, error)     // { ID, FileName, Size, UploadedAt }

func main() {
	storage := filestorage.New()

	fileID, err := storage.Upload("./photos/profile.jpg")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileID) // e.g. "961dada0d04a40afa8e779c705d25408"

	meta, _ := storage.GetMetadata(fileID)
	fmt.Printf("%+v\n", meta)
	// {ID:961dada0... FileName:profile.jpg Size:102400 Checksum:9be16830... UploadedAt:2026-07-10 09:27:16 +0000 UTC}

	if err := storage.Download(fileID, "./downloads/profile.jpg"); err != nil {
		log.Fatal(err)
	}

	storage.Delete(fileID)
	_, err = storage.GetMetadata(fileID)
	fmt.Println(errors.Is(err, filestorage.ErrNotFound)) // true
}

func NewFileStorage(....) *FileStorage {
   return &FileStrorage{
     ....
}
}

*/
