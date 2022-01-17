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
		remove(v)
	}
}

func remove(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}