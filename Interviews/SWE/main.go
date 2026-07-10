package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

type file struct {
	ID string
	Name string
	Size int64
	UploadedAt time.Time
}

type FileStorage struct {
	uploadedFiles []file
}


func NewFileStorage(
	uploadedFiles []file,
) *FileStorage {
	return &FileStorage{
		uploadedFiles: uploadedFiles,
	}
}

func main() {
	var ctx context.Context

	storage := NewFileStorage([]file{})

	fileID, err := storage.Upload(ctx, "./Files/first-file.txt")
	if err != nil {
		log.Fatal(err)
	}

	meta, _ :
	
}

func (fs *FileStorage) Upload(ctx context.Context, filePath string) (string, error) {
	uploadedFile, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("Could not read file from filepath")
	}

	info, err := uploadedFile.Stat()
	if err != nil {
		return "", fmt.Errorf("Uploaded File has no metadata")
	}

	newFile := file{
		ID: uuid.NewString(),
		Name: info.Name(),
		Size: info.Size(),
		UploadedAt: info.ModTime(),
	}

	fs.uploadedFiles = append(fs.uploadedFiles, newFile)

	return newFile.ID, nil
}

func (fs *FileStorage) GetMetadata(ctx context.Context, fileID string) (file error) {

}


func

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
