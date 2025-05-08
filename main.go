package main

import (
	"fmt"
	"mylearning/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	fmt.Println(greetings.Hello("Anamika"))
	fmt.Println(greetings.Bless())
	variables()
}
