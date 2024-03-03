package main

import "fmt"

func printPersonInfo(name string, age int, weight float32) {
	fmt.Printf("Name: %s\nAge:%d\nWeight:%f\n", name, age, weight)
}

func unlimitedArgs(nums ...int) {
	fmt.Println("Nums is a slice with integers: ", nums)
}

func isAdult(age int) bool {
	if age >= 18 {
		return true
	}
	return false
}

func multipleReturnedValues() (string, error) {
	return "", nil
}
