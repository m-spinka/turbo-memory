package main

import (
	"log"
)

func main() {
	log.Println("test")
	Greet("Hello world!")
}

func Greet(str string) error {
	log.Println(str)

	return nil
}
