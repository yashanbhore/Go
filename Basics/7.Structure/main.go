package main

import "fmt"

type Person struct {
	Name    string
	Age     int64
	Gender  string
	Address Address
}

type Address struct {
	pincode int
	city    string
}

func main() {

	// Zero Value
	var po Person
	fmt.Println(po)

	var p Person

	p.Name = "Yash"
	p.Age = 21
	p.Gender = "Male"
	p.Address = Address{
		425001,
		"Jalgaon",
	}

	fmt.Println(p)


	//Using Posotional Syntax
	p1 := Person{"Aryan", 21, "Male", Address{425001, "Jalgaon"}}
	fmt.Println(p1)

	//Using Field Name
	p2 := Person{Name: "Swapnil", Age: 21, Gender: "Male", Address: Address{425001, "Jalgaon"}}
	fmt.Println(p2)


	p3 := Person{Name: p.Name, Age: p.Age, Gender: p.Gender, Address: Address{425001, "Jalgaon"}}
	fmt.Println(p3)

}
