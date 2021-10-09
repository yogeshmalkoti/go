package greetings

import "fmt"

//hello Returns the name of the person with greeting
func Hello(name string) string {
	//var message string
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
