package main

import (
	"os"

	"github.com/eventials/go-tus"
)

func main() {
	f, err := os.Open("my-file.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	// create the tus client.
	client, _ := tus.NewClient("http://localhost:8080/file/upload", nil)

	// create an upload from a file.
	upload, _ := tus.NewUploadFromFile(f)
	upload.Metadata["token"] = "4e1610bdab129d2143652093de01e15200000000000000000000000000000000"
	// create the uploader.
	uploader, err := client.CreateUpload(upload)
	if err != nil {
		panic(err)
	}

	// start the uploading process.

	err = uploader.Upload()
	if err != nil {
		panic(err)
	}
}
