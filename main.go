package main

import "fmt"

const (
	USD_EUR = 0.86
	USD_RUB = 81.06
	EUR_RUB = USD_RUB / USD_EUR
)

func main() {
	readInput()
}

func readInput() {

	var original, target string
	var quantity int

	currency := []string{"RUB", "USD", "EUR"}

	for {
		fmt.Printf("Ввод исходной валюты (%s, %s, %s):", currency[0], currency[1], currency[2])
		fmt.Scan(&original)

		bre := false

		for _, val := range currency {
			if val == original {
				bre = true
			}
		}

		if bre {
			break
		}

	}

	for {
		fmt.Print("Ввод числа: ")
		_, err := fmt.Scan(&quantity)

		if err == nil {
			break
		}
	}

	for {
		fmt.Printf("Ввод целевой валюты (%s, %s, %s):", currency[0], currency[1], currency[2])
		fmt.Scan(&target)

		bre := false

		for _, val := range currency {
			if val == target && target != original {
				bre = true
			}
		}

		if bre {
			break
		}
	}

	num := calculation(quantity, original, target)
	fmt.Print(num)

}

func calculation(quantity int, original string, target string) int {

	switch {
	case (original == "RUB" && target == "USD") || (original == "USD" && target == "RUB"):
		return int(USD_RUB * float64(quantity))
	case (original == "RUB" && target == "EUR") || (original == "EUR" && target == "RUB"):
		return int(EUR_RUB * float64(quantity))
	case (original == "EUR" && target == "USD") || (original == "USD" && target == "EUR"):
		return int(USD_EUR * float64(quantity))		
	}

	return 0

}
