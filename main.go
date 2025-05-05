package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	ExtractionAndSorting("24z6 1x23 y369 89a 900b")

}

func ExtractionAndSorting(text string) {

	type Pair struct {
		Key string
		Num float64
	}
	Pairs := make([]Pair, 0)
	words := strings.Fields(text)

	// fmt.Println(text)
	// fmt.Println(words)

	for _, word := range words {
		// fmt.Printf("Word %d is: %s\n", idx, word)
		var number float64
		letter := ""
		strFloat64 := ""

		for _, char := range word {
			if !unicode.IsDigit(char) {
				letter += string(char)
			} else if unicode.IsDigit(char) {
				strFloat64 += string(char)
			}
		}

		num, err := strconv.ParseFloat(strFloat64, 64)
		if err != nil {
			fmt.Println("Ошибка преобразования строки в число", err)
		} else {
			number = num
		}
		Pairs = append(Pairs, Pair{letter, number})
	}

	sort.SliceStable(Pairs, func(i, j int) bool {
		return Pairs[i].Key < Pairs[j].Key
	})
	fmt.Println(Pairs)

	NewSlice := make([]float64, 0)
	for _, num := range Pairs {
		NewSlice = append(NewSlice, num.Num)
	}

	fmt.Println(NewSlice) // Записал в новый слайс числа (второе значение структур) для дальнейших операций
	// fmt.Println(NewSlice[0])
	// fmt.Println(len(NewSlice))

	result := NewSlice[0]
	if len(NewSlice) == 1 {
		fmt.Println(result)
	} else {

		// i := 0
		var j int

		for i := 0; i < len(NewSlice)-1; i++ {
			j = i % 4
			if j == 0 {
				result += NewSlice[i+1]
			} else if j == 1 {
				result -= NewSlice[i+1]
			} else if j == 2 {
				result *= NewSlice[i+1]
			} else if j == 3 {
				result /= NewSlice[i+1]
			}
		}

		fmt.Printf("\"%s\"\nРезультат:\n%d\n", text, int(math.Round(result)))
	}
}
