package main

import "fmt"

type Contact struct {
	fName string
	sName string
	age   int
	phone string
}

func doSm() {
	vasya := Contact{
		fName: "Vasya",
		sName: "Pupkin",
		age:   25,
		phone: "79998887667",
	}
	fmt.Println(vasya)
}
