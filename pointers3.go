package main

import "fmt"

func main() {
	value := 20
	fmt.Println("Adresa: ", &value, ", vrednost: ", value)

	p := &value
	fmt.Println("Adresa: ", &p, ", vrednost: ", *p)
}
