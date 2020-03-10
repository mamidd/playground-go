package main

import "fmt"

func main() {
	name := "bill"

	namePointer := &name

	fmt.Println(namePointer) // questa è una variabile che come valore ha un indirizzo di memoria
	fmt.Println(&name)

	fmt.Println(&namePointer) // il suo indirizzo di memoria è ovviamente diverso dal valore che contiene
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}
