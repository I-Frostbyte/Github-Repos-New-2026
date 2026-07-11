# 🚀 File Exercise #1

-----------------------------------------------------------------------------------------------------------------------------
File Exercise #1
- Create a simple program that allows for the creation, deletion, rename, editing and copying (single & bulk) of files in between two folders (source & destination).
- For simplicity, files so far should be .txt

- CreateFile(filename, filetype string) error {}
- RenameFile(filename, newFilename string) error {}
- EditFile(filename string, content []bytes) error {}
- DeleteFile(filename string) error {}
- CopyFile(filename, source, destination string) (string, error) {}
-----------------------------------------------------------------------------------------------------------------------------