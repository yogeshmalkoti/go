package main

import (
	"fmt"

	"github.com/yogeshmalkoti/go/mymodules/greetings"
)

func main() {
	message := greetings.Hello("Yogesh")
	fmt.Println(message)
}
