package main

import "fmt"

const (
	USD_EUR = 0.86
	USD_RUB = 81.06
	EUR_RUB = USD_RUB / USD_EUR
)

func main() {
	fmt.Print("Готово!")
}
