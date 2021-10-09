package greetings

import (
	"errors"
	"fmt"
)

//hello Returns the name of the person with greeting
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empy value")
	}
	//var message string
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
