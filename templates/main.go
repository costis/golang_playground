package main

import (
	"log"
	//"os"
	//"strings"
	//"text/template"
	"io/ioutil"
	"fmt"
	"os"
	//"io"
)

type person struct {
	Name string
	Age  int
}

func directoryListing(path string) {
	fInfo, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range fInfo {
		fmt.Println(file.Name())
	}
}

func main() {
	filePath := os.Args[1]
	directoryListing(filePath)
}
