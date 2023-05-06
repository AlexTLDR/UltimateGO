package main

import (
	"compress/gzip"
	"crypto/sha1"
	"io"
	"os"
)

func main() {
	sha1Sum("http.log.gz")
}

// $ cat http.log.gz| gunzip | sha1sum
func sha1Sum(fileName string) (string, error) {
	// idion: aquire resource, check for error, defer release
	file, err := os.Open(fileName)
	if err != nil {
		return "", nil
	}
	defer file.Close() //defered are called in LIFO order

	r, err := gzip.NewReader(file)
	if err != nil {
		return "", err
	}
	//io.CopyN(os.Stdout, r, 100)
	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	return "", nil
}
