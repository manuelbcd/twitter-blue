package main

import (
	"fmt"

	"github.com/containerd/containerd"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	message.SetString(language.Spanish, "hello world", "hola mundo")
	p := message.NewPrinter(language.Spanish)
	fmt.Println(p.Sprintf("hello world"))
	fmt.Println(containerd.DefaultSnapshotter)
}
