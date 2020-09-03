package main

import (
	"fmt" // Package Alias
	str "strings"

	"github.com/adicasa/packer.git/numeros"
	"github.com/adicasa/packer.git/strings"
	"github.com/adicasa/packer.git/strings/greeting"
	// Importing a nested package
)

func main() {
	fmt.Println(numeros.IsPrime(19))

	fmt.Println(greeting.WelcomeText)

	fmt.Println(strings.Reverse("adicasa"))

	fmt.Println(str.Count("Go es Amigable y agradable. Me gusta trabajar en Go", "Go"))
}
