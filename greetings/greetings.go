package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v Welcome!", name) // create variable and initialise it also
	return message
}
