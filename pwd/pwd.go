package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dir := filepath.Base(path)
	fmt.Println(dir)
}