package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var is_roman bool = false

var operators = [4]string{"+", "-", "*", "/"}
var romes = [7]string{"M", "C", "D", "X", "L", "V", "I"}

func extra(z_in string, devider string) int {
	var is_x_roman bool = false
	var is_y_roman bool = false
	out := 0
	x, y, _ := strings.Cut(z_in, devider)
	x = strings.TrimSpace(x)
	y = strings.TrimSpace(y)

	var x_num int
	var y_num int
	for _, i := range romes {
		if strings.Contains(x, i) {
			is_x_roman = true
		}
		if strings.Contains(y, i) {
			is_y_roman = true
		}
	}

	if is_x_roman == false && is_y_roman == false {
		is_roman = false
		x_num, _ = strconv.Atoi(x)
		y_num, _ = strconv.Atoi(y)
		if 1 > x_num || 1 > y_num || 10 < x_num || 10 < y_num {
			fmt.Println("Неть. Операнд должен быть от 1 до 10 включительно и быть целым числом")
			os.Exit(0)
		}
	} else if is_x_roman != is_y_roman {
		fmt.Println("Неть. Оба операнда должны быть одной системы счисления")
		os.Exit(0)
	} else {
		is_roman = true
		x_num = from_roman((x))
		y_num = from_roman(y)

	}

	switch devider {
	case "+":
		out = x_num + y_num
	case "-":
		out = x_num - y_num
	case "*":
		out = x_num * y_num
	case "/":
		out = x_num / y_num

	}
	if is_roman == true && out < 1 {
		fmt.Println("Неть. Римские числа могут быть только положительными")
		os.Exit(0)
	}
	return out
}

func in_to_roman(arab int) string {

	rom_dict := []struct {
		value int
		digit string
	}{

		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, rom_dict := range rom_dict {
		for arab >= rom_dict.value {
			roman.WriteString(rom_dict.digit)
			arab -= rom_dict.value
		}
	}

	return roman.String()
}

func from_roman(rom string) int {
	arab := 0
	rom_dict := []struct {
		value int
		digit string
	}{
		{10, "X"},
		{9, "IX"},
		{8, "VIII"},
		{7, "VII"},
		{6, "VI"},
		{5, "V"},
		{4, "IV"},
		{3, "III"},
		{2, "II"},
		{1, "I"},
	}

	for _, rom_dict := range rom_dict {
		if rom == rom_dict.digit {
			arab = rom_dict.value
		}
	}
	if arab == 0 {
		fmt.Println("Неть. Операнд должен быть от 1 до 10 включительно")
		os.Exit(0)

	}

	return arab
}

func main() {

	for {
		devider := "None"
		count := 0
		reader := bufio.NewReader(os.Stdin)
		z, _ := reader.ReadString('\n')
		z = strings.TrimSpace(z)
		for _, i := range operators {
			if strings.Contains(z, i) {
				devider = i
				count += 1
				if count > 1 || strings.Count(z, i) > 1 {
					fmt.Println("Неть, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
					os.Exit(0)
				}
			}

		}

		if devider == "None" {
			fmt.Println("Неть, вы не указали ни один из операторов (+, -, /, *)")
			os.Exit(0)
		}

		for_print := extra(z, devider)
		for_print_rom := in_to_roman(for_print)
		if is_roman == false {
			fmt.Println(for_print)
		} else {
			fmt.Println(for_print_rom)
		}

	}
}
