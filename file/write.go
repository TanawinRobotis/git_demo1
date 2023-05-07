package main

import (
	"os"
)

func main() {
	data1 := []byte("Hello\n AI-Center")
	err := os.WriteFile("data.txt", data1, 0644)

	if err != nil {
		panic(err)
	}

	f, err := os.Create("employeeName")
	panic(err)

	defer f.Close()

	// data2 := []byte("Hello\n AI-Center-2")
	// os.WriteFile("employeeName.txt", data2, 0644)
}
