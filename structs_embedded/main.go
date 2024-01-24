package main

import "fmt"

func main() {
	type Contact struct {
		email   string
		address string
		phone   int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	john := Employee{
		name:   "John Keller",
		salary: 50000,
		contactInfo: Contact{
			email:   "johnk@gmail.com",
			address: "street address",
			phone:   9870123456,
		},
	}

	fmt.Printf("%+v\n", john)
	fmt.Printf("Employees's email: %s\n", john.contactInfo.email)

	john.contactInfo.email = "abc@gmail.com"

	fmt.Printf("Employees's email: %s\n", john.contactInfo.email)

	myContact := Contact{email: "abcd@gmail.com", phone: 9876543219, address: "Street address"}
	fmt.Println(myContact)

}
