package main

import (
	"os"
	"strconv"
)

func main() {
	for a := 1; a <= 1000000; a++ {
		createFile(a)
	}
}

func createFile(a int) {

	f, err := os.Create("osCreate/file-" + strconv.Itoa(a))
	if err != nil {
	}
	defer f.Close()
}
