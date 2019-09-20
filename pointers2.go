package main

import "fmt"

type user struct {
	name   string
	email  string
	logins int
}

func main() {

	bill := user{
		name:  "Bill",
		email: "bill@gmail.com",
	}

	display(&bill)
	increment(&bill.logins)
	display(&bill)

}

func increment(logins *int) {
	*logins++
	fmt.Println("Ispisuje adresu, ispisuje broj prijava, ispisuje adresu nakon prijave\n", &logins, *logins, logins)
}
func display(u *user) {
	fmt.Println("Ispisuje naziv usera i njegov email pre i posle inkrementacije\n", u, *u)
}
