package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	var input string
	var number1, number2, znak1, result int
	var numbers []string
	var rome_check_1 bool = false
	var rome_check_2 bool = false
	var arabic [10]string = [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	var rome = [14]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XL", "L", "XC", "C"}
	var znaki [4]string = [4]string{"+", "-", "*", "/"}

	fmt.Scanln(&input)
	/*if (len(input)) > 5 {
		fmt.Println("error in lenght")
	}*/
	for i := 0; i < len(znaki); i++ {
		if strings.Contains(input, znaki[i]) {
			numbers = strings.Split(input, znaki[i])
			if len(numbers) > 2 {
				//fmt.Println("error too many arguments")
				panic(errors.New("error too many arguments"))
			}
			znak1 = i
			// fmt.Println(numbers, znak1)
			break
		}
		if (strings.Contains(input, arabic[i]) == false) && (i == (len(znaki) - 1)) {
			//fmt.Println("error  in sign")
			panic(errors.New("error  in sign"))
		}

	}
	for j := 0; j < 10; j++ {
		if numbers[0] == rome[j] {
			rome_check_1 = true
		}
	}
	for j := 0; j < 10; j++ {
		if numbers[1] == rome[j] {
			rome_check_2 = true
		}
	}
	if rome_check_1 != rome_check_2 {
		// fmt.Println("error different systems")
		panic(errors.New("error different systems"))

	}

	for j := 0; j < 10; j++ {
		switch {
		case numbers[0] == arabic[j]:
			number1 = j + 1
		case numbers[0] == rome[j]:
			number1 = j + 1
		}
		if number1 == 0 {
			// fmt.Println("")
			panic(errors.New("out of range number was entered"))
		}

	}
	for j := 0; j < 10; j++ {
		switch {
		case numbers[1] == arabic[j]:
			number2 = j + 1
		case numbers[1] == rome[j]:
			number2 = j + 1
		}
		if number2 == 0 {
			// fmt.Println("out of range number was entered")
			panic(errors.New("out of range number was entered"))
		}
	}
	switch znak1 {
	case 0:
		result = number1 + number2
	case 1:
		if rome_check_1 && number2 > number1 {
			panic(errors.New("error negatives not possible in rome numbers"))
		}
		result = number1 - number2
	case 2:
		result = number1 * number2
	case 3:
		result = number1 / number2
	}
	if rome_check_1 == false {
		fmt.Println(result)
		return
	}
	var rome_result []string = []string{}
	if result == 100 {
		fmt.Println(rome[13])
		return
	}
	desyatki := result / 10
	//fmt.Println(result, desyatki, rome_result)
	switch {
	case desyatki >= 1 && desyatki < 4:
		rome_result = append(rome_result, strings.Repeat(rome[9], desyatki))
		//fmt.Println("here")
	case desyatki == 4:
		rome_result = append(rome_result, rome[10])
		//fmt.Println("here1")
	case desyatki >= 5 && desyatki < 9:
		rome_result = append(rome_result, rome[11])
		rome_result = append(rome_result, strings.Repeat(rome[9], desyatki-5))
		//fmt.Println(rome_result)
		//fmt.Println("here2")
	case desyatki == 9:
		rome_result = append(rome_result, rome[12])
		//fmt.Println("here3")
	}
	desyatki = result % 10
	switch {
	case desyatki >= 1 && desyatki < 4:
		rome_result = append(rome_result, strings.Repeat(rome[0], desyatki))
		//fmt.Println("here")
	case desyatki == 4:
		rome_result = append(rome_result, rome[3])
		//fmt.Println("here1")
	case desyatki >= 5 && desyatki < 9:
		rome_result = append(rome_result, rome[4])
		rome_result = append(rome_result, strings.Repeat(rome[0], desyatki-5))
	//	fmt.Println(rome_result)
	//fmt.Println("here2")
	case desyatki == 9:
		rome_result = append(rome_result, rome[8])
		//fmt.Println("here3")
	}
	fmt.Println(strings.Join(rome_result, ""))
}
