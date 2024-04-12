package main

import (
	"fmt"
	"tutorials/learn-go-modules/greetings"
)

func main() {
	// get a greeting message and print it
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
