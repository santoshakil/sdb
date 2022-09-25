package file

import (
	"io/ioutil"
)

func WatchFile(path string) (string, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(file), nil
}
