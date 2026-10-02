package client

import (
	"log"
	"os"
)

func UploadFile(filePaths []string) {
	for _, path := range filePaths {
		fileData := getFileMetadata(path)
		if fileData == nil {
			continue
		}
		if !fileData.IsDir {
			sendFile(*fileData)
		}
	}
}

type FileMetadata struct {
	FileName string
	FileSize int64 // fileStats.Size() is int64
	Path     string
	IsDir    bool
}

func getFileMetadata(path string) *FileMetadata {
	fileStats, err := os.Stat(path)

	if os.IsNotExist(err) {
		log.Printf("%s File does not exist", path)
		return nil // Returns "nothing" (nil pointer)

	} else if err != nil {
		log.Fatal("INTERNAL ERROR")
	}

	return &FileMetadata{
		FileName: fileStats.Name(),
		FileSize: fileStats.Size(),
		Path:     path,
		IsDir:    fileStats.IsDir(),
	}
}
