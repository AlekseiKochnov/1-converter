package main

import (
	"fmt"
	"strconv"
	"strings"
)

var oper = map[string]func([]int){
	"AVG": AVG,
	"SUM": SUM,
	"MED": MED,
}

func main() {

	dataEntry()

}

func dataEntry() {

	var operation string
	var num string

	for {
		fmt.Print("Ввод операции (AVG - среднее, SUM - сумму, MED - медиану): ")
		fmt.Scan(&operation)

		if _, ok := oper[operation]; ok {
			break
		}

	}

	for {

		fmt.Print("Ввод чисел через запятую (2, 10, 9): ")
		_, er := fmt.Scan(&num)

		if er == nil {
			break
		}

	}

	slNum := splitLine(num)
	oper[operation](slNum)
}

func splitLine(str string) []int {

	numStrs := strings.Split(str, ",")

	numbers := make([]int, len(numStrs), cap(numStrs))

	for _, val := range numStrs {

		val = strings.TrimSpace(val)

		num, err := strconv.Atoi(val)
		if err != nil {
			fmt.Printf("Ошибка преобразования: %v\n", err)
			continue
		}
		numbers = append(numbers, num)

	}

	return numbers

}

func SUM(num []int) {
	var sum int
	for _, num := range num {
		sum += num
	}
	fmt.Println(sum)
}

func AVG(num []int) {
	var sum int
	for _, num := range num {
		sum += num
	}
	fmt.Println(float64(sum) / float64(len(num)))
}

func MED(num []int) {
	var median float64
	sliceLen := len(num)
	center := sliceLen / 2

	if sliceLen%2 == 0 {
		median = float64((num[center-1] + num[center]) / 2)
	} else {
		median = float64(num[center])
	}

	fmt.Println(median)
}
