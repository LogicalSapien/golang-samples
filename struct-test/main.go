package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	address
}

type contactInfo struct {
	email  string
	number int
}

type address struct {
	doorNo   string
	postcode string
}

// call by value only
func (p person) print() {
	fmt.Printf("%+v", p)
}

// call by reference
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func main() {
	// sample way of assinging
	// tony := person{"Tony", "Stark"}
	// other way
	tony := person{
		firstName: "Tony",
		lastName:  "Stark",
		contact: contactInfo{
			email:  "tony@stark.com",
			number: 1234567890,
		},
		address: address{
			doorNo:   "300",
			postcode: "POSTCODE",
		},
	}
	fmt.Println(tony)
	// with var names
	fmt.Printf("%+v", tony)
	fmt.Println()
	tony.firstName = "Antony"
	fmt.Printf("%+v", tony)
	fmt.Println()

	// for call by reference
	pointerToPerson := &tony
	pointerToPerson.updateName("Tony")
	tony.print()
	fmt.Println()

	// shortcut for go by refernece
	// go takes it as pass bby reference because receiver has *
	tony.updateName("Anthony")
	tony.print()
}
