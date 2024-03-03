package main

import "fmt"

func doMap() {
	contactList := map[string]string{"Vasya": "123123123", "Petya": "321321321", "Sanya": "000000000"}
	fmt.Println(contactList)

	for key, val := range contactList {
		fmt.Printf("Key: %s\tValue: %s\n", key, val)
	}

	delete(contactList, "Vasya")

	fmt.Println(contactList)
}
