package utils

import (
	"log"
	"os"
)

func CheckFile(fileUri string) (*os.File, error) {
	f, err := os.Open(fileUri)

	if err != nil {
		log.Fatal("Failed to access file")
		return nil, err
	}

	return f, nil
}
