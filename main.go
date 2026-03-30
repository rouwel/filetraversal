package main

import (
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]

	info, err := os.Stat(filename)
	if err != nil {
		fmt.Println("Path does not exist")
	}

	if info.IsDir() {
		fmt.Println("Directory found")
	} else {
		fmt.Println("File found")
	}
}
