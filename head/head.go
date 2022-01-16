package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("no file specified")
		os.Exit(1)
	}
	file, err := os.Open(args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()
	for i, line := range lines {
		if i == 10 {
			os.Exit(0)
		}
		fmt.Println(line)
	}
}