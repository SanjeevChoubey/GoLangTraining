package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// diff way of declaring struct
	//sanjeev := person{"Sanjeev", "Choubey"}

	// sanjeev := person{firstName: "Sanjeev", lastName: "Choubey"}

	// var sanjeev person
	// sanjeev.firstName = "Sanjeev"
	// sanjeev.lastName = "Choubey"
	// fmt.Println(sanjeev)
	// fmt.Printf("%+v", sanjeev)

	sanjeev := person{
		firstName: "Sanjeev",
		lastName:  "Choubey",
		contact: contactInfo{
			email: "schoubey80@gmail.com",
			zip:   560037,
		},
	}
	// pass by value
	//sanjeev.updateFirstname("SANJEEV")

	// Pass by refrence
	sanjeevaddr := &sanjeev
	sanjeevaddr.updateFirstname("SANJEEV")
	sanjeev.print()
}

func (pointerToperson *person) updateFirstname(newfirstname string) {
	(*pointerToperson).firstName = newfirstname
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
