package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}

func main() {
	alex := person{
		firstname: "Alex",
		lastname:  "Anderson",
		contact: contactInfo{
			email: "Alex@gmail.com",
			zip:   10110},
	}
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
