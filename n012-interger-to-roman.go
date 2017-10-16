package main

import "fmt"

func intToRoman(num int) string {
	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	var result string
	for i := 0; num != 0; i++ {
		for num >= value[i] {
			num -= value[i]
			result += symbol[i]
		}
		fmt.Println(result)
		fmt.Println(num)
	}

	return result
}
