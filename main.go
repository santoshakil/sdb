package main

import (
	"fmt"

	"github.com/santoshakil/sdb/functions/file"
)

func main() {
	file.CreateFile("test", "Hello, world!")
	watch, _ := file.WatchFile("test")
	fmt.Println(watch)
}
