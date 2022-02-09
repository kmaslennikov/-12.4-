package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileName := "text.txt"
	var bb bytes.Buffer
	bb.WriteString("Hello")
	bb.WriteString("\n")
	bb.WriteString("Good day!")
	if err := ioutil.WriteFile(fileName, bb.Bytes(), 0666); err != nil {
		fmt.Println(err)
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	result, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(string(result))
}
