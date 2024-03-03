package main

import "fmt"

func doArr() {
	contactList := [3]string{"Vasya", "Petya", "Sanya"}

	fmt.Println(contactList)
	fmt.Println(contactList[0])

	contactList[0] = "Andrew"
	fmt.Println(contactList[0])

	fmt.Println(len(contactList), cap(contactList))

	dstArr := contactList
	fmt.Println("Copy of contactList", dstArr)

}

func doSlc() {
	var length int
	length = 3

	var capp int
	capp = 5

	slc := []string{"string", "string", "string"}
	slc1 := make([]string, length, capp)

	fmt.Println(slc)
	fmt.Println(slc1)

	for i, v := range slc {
		fmt.Printf("index:%d\tvalue:%s\n", i, v)
	}

	slc1 = append(slc1, "hello", "world")
}

func iter(slc []string) {
	for _, v := range slc {
		fmt.Println(v)
	}

	for i := 0; i < len(slc); i++ {
		fmt.Println(slc[i])
	}
}
