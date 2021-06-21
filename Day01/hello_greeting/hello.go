package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	// Set properties of the predefined logge

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// One name
	//message, err  := greetings.Hello("Gladys")
	// No name
	//message, err  := greetings.Hello("")
	// Many names
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	// No error
	fmt.Println(messages)
}
