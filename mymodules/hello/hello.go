package main

import (
	"fmt"
	"log"

	"github.com/yogeshmalkoti/go/mymodules/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// A slice of names.
	names := []string{"Yogesh", "Amit", "Anil", "Sandeep"}
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
