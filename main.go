package main

import (
	"fmt"

	"github.com/santoshakil/sdb/functions/file"
)

func main() {
	file.CreateFile("test", "Hello, world!")
	content, _ := file.ReadFile("test")
	fmt.Println(content)
}
