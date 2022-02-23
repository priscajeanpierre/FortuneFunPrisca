package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	go fortune()
	myChan := make(chan string)

}

func fortune(c chan string) {
	for true {

		//open file
		file, err := ioutil.ReadFile("Fortunes.txt")
		if err != nil {
			log.Fatalln("Error reading file: ", err)
		}

		fileContents := string(file)
		lines := strings.Split(fileContents, "%%")

	}

}
