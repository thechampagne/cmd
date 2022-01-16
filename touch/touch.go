package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("missing file name")
		os.Exit(1)
	}
	for i, v := range args {
		if i == 0 {
			continue
		}
		create(v)
	}
}

func create(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
}