package main

import "errors"

const winePrice = 1000

func buyWine(age int, moneyAmount int) (int, error) {
	if age >= 18 {
		if moneyAmount >= winePrice {
			return moneyAmount - winePrice, nil
		}
		return 1, errors.New("not enough money")
	}
	return 1, errors.New("too young to buy wine")
}

func buyWine1(age int, moneyAmount int) (int, error) {
	if age >= 18 && moneyAmount >= winePrice {
		return moneyAmount - winePrice, nil
	}
	return 1, errors.New("some err")
}
