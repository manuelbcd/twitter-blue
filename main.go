package main

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	message.SetString(language.Spanish, "hello world", "hola mundo")
	p := message.NewPrinter(language.Spanish)
	fmt.Println(p.Sprintf("hello world"))
}
