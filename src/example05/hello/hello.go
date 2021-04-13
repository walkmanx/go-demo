package main

import (
	"fmt"
	"log"

	"com.zy/greeting"
)

func main() {
	log.SetPrefix("greeting: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	message, err := greeting.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
