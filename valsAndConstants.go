package main

import (
	"fmt"
	"reflect"
)

func valAndConst() {
	var message string
	message = "Hello world"
	fmt.Println(message)

	var name string
	var age int
	var weight float32
	var isAdult bool

	name = "Vasya"
	age = 25
	weight = 70.2
	isAdult = true

	fmt.Println("Person:", name)
	fmt.Println("Age:", age)
	fmt.Println("Weight:", weight)
	fmt.Println("Is adult:", isAdult)

	const goldenRatio = 1.618
	//goldenRatio = 1 // Is not assignable @CONST@
	fmt.Println(goldenRatio, reflect.TypeOf(goldenRatio))
}
