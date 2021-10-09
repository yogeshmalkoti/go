package main

import (
	"fmt"
	"log"

	"github.com/yogeshmalkoti/go/mymodules/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
