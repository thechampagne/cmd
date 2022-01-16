package main

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() {
			_, err := color.New(color.FgYellow).Print(file.Name() + " ")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			continue
		}
		_, err := color.New(color.FgGreen).Print(file.Name() + " ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}