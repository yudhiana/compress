package main

import (
	"io"
	"log"
	"os"

	"github.com/alexmullins/zip"
)

// compress files with password

func main() {
	password := "1234"
	fname, err := os.Create("testzip.zip")
	if err != nil {
		log.Fatal(err)
	}
	zipw := zip.NewWriter(fname)
	defer zipw.Flush()
	defer zipw.Close()

	w, err := zipw.Encrypt("hello.txt", password)
	if err != nil {
		log.Fatal(err)
	}

	fileToZip, err := os.Open("hello.txt")
	defer fileToZip.Close()
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(w, fileToZip)
	if err != nil {
		log.Fatal(err)
	}
}
