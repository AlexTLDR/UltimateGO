package main

import "os"

// $ cat http.log.gz| gunzip | sha1sum
func sha1Sum(fileName string) (string, error) {
	// idion: aquire resource, check for error, defer release
	file, err := os.Open(fileName)
	if err != nil {
		return "", nil
	}
	defer file.Close()

	return "", nil
}
