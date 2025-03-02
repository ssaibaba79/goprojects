package main

import (
	"composition/identities"
	"fmt"
)

func main() {

	p := identities.NewPerson("Shashti", "t@mail.com", "123-456-7890")
	fmt.Println(*p)
	s := identities.NewStaff("Sarvan", "s@mail.com", "123-567-8888", "admin")
	fmt.Println(*s)

}
