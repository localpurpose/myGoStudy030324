package main

import "fmt"

func ptrHi() {
	name := "Artem"
	pname := &name

	fmt.Println(name, pname, *pname)

	setName(&name, "Vasya")
	fmt.Println("Новое имя:", name)
}

func setName(name *string, nameToSet string) {
	*name = nameToSet
}

func emptyPtr() {
	var empty *string
	fmt.Println(empty) // == nil
}
