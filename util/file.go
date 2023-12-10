package util

import (
	"log"
	"os"
)

func RetrieveOrCreateFile(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)

	if err != nil {
		log.Fatal(err)
	}
	return file
}
