package file

import (
	"io/ioutil"
	"os"
)

func CreateFile(name, content string) error {
	return ioutil.WriteFile(name, []byte(content), os.ModePerm)
}
