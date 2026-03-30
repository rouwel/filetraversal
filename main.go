package main

import (
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]

	info, err := os.Stat(filename)
	if err != nil {
		fmt.Printf("Path does not exist:%v\n", err)
		return
	}

	if info.IsDir() {
		fmt.Println("Directory found")
	} else {
		fmt.Println("File found")
	}
}
