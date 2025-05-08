package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("write ur name")
	var name string
	fmt.Scan(&name)
	fmt.Println("name", name)

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println("You said:", text)

}

// go run userInput/main.go
