package main

import (
	"fmt"

	"github.com/yogeshmalkoti/mymodules/greetings"
)

func main() {
	message := greetings.Hello("Yogesh")
	fmt.Println(message)
}
