package main

import (
	"fmt"
	"log"

	"com.zy/greeting"
)

func main() {
	log.SetPrefix("greeting: ")
	log.SetFlags(0)

	message, err := greeting.Hello("asd")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
