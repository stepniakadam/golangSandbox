package main

import "fmt"
import "github.com/adamstepniak41/golangSandbox/filesHandling"

func main() {
	var persons = exercise1.ReadPersons()

	fmt.Print("Person name is = \n", persons)
}