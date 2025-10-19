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

		bre := inputValidation(currency, original)

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

		if target == original {
			continue			
		}

		bre := inputValidation(currency, target)

		if bre {
			break
		}
	}

	num := calculation(quantity, original, target)
	fmt.Print(num)

}

func calculation(quantity int, original string, target string) float64 {

	switch {
	case (original == "RUB" && target == "USD"):
		return  float64(quantity) / USD_RUB
	case (original == "RUB" && target == "EUR"):
		return  float64(quantity) / EUR_RUB
	case (original == "EUR" && target == "USD"):
		return float64(quantity) / USD_EUR
	case (original == "USD" && target == "EUR"):
		return float64(quantity) * USD_EUR	
	case (original == "USD" && target == "RUB"):
		return  float64(quantity) * USD_RUB	
	case (original == "EUR" && target == "RUB"):
		return  float64(quantity) * EUR_RUB				
	}

	return 0

}

func inputValidation(currency []string, curren string) bool {

	for _, val := range currency {
		if val == curren {
			return  true
		}
	}

	return false

}
