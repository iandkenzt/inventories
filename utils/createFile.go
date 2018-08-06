package utils

import (
	"fmt"
	"os"
)

// IsError ...
func IsError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

// CreateFile ...
func CreateFile(path string) (result string, err error) {
	// check path
	var _, e = os.Stat(path)

	// create new file
	if os.IsNotExist(e) {
		var file, err = os.Create(path)
		if IsError(err) {
			return "", err
		}
		defer file.Close()
	}

	return path, nil
}
