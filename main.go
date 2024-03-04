package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var is_roman bool = false

var operators = [10]string{"+", "-", "*", "/", ")", "(", "%", "^", "!", "="}
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
			panic(fmt.Sprintln("Неть. Операнд должен быть от 1 до 10 включительно и быть целым числом"))
		}
	} else if is_x_roman != is_y_roman {
		panic(fmt.Sprintln("Неть. Оба операнда должны быть одной системы счисления"))
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
	fmt.Println("Привет! Это скромный калькулятор, который умеет складывать, вычитать, умножать и делить целые числа от 1 до 10 в арабской и римской системах счисления.\n",
		"Чтобы начать, напиши здесь выражение в формате 'x+y' и нажми Enter. ")
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
					panic(fmt.Sprintf("Неть, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)"))
				}
			}

		}

		if devider == "None" {
			panic(fmt.Sprintf("Неть, вы не указали ни один из операторов (+, -, /, *)"))
		}

		for_print := extra(z, devider)
		for_print_rom := in_to_roman(for_print)
		if is_roman == false {
			fmt.Println(for_print)
		} else {
			if is_roman == true && for_print < 1 {
				panic(fmt.Sprintf("Неть. Римские числа могут быть только положительными"))
			} else {
				fmt.Println(for_print_rom)
			}
		}

	}
}
