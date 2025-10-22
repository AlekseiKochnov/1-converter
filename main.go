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
		bre := readQuantity(&quantity)

		if bre {
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

	wellDivision := map[string]float64{
		"RUB_USD": USD_RUB,
		"RUB_EUR": EUR_RUB,
		"EUR_USD": USD_EUR,
	}

	wellMultiplication := map[string]float64{
		"USD_EUR": USD_EUR,
		"USD_RUB": USD_RUB,
		"EUR_RUB": EUR_RUB,
	}	

	Division, ok := wellDivision[original + "_" + target]

	if ok {
		return float64(quantity) / Division	
	}

	Multiplication, ok := wellMultiplication[original + "_" + target]

	if ok {
		return float64(quantity) * Multiplication
	}	

	return 0

}

func inputValidation(currency []string, curren string) bool {

	for _, val := range currency {
		if val == curren {
			return true
		}
	}

	return false

}

func readQuantity(quantity *int) bool {
	fmt.Print("Ввод числа: ")
	_, err := fmt.Scan(quantity)

	if err == nil {
		return true
	}

	return false
}
