package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// positive / negative / zero
func identifyingmark(chislo int) string {
	if chislo > 0 {
		return "Positive"
	} else if chislo < 0 {
		return "Negative"
	}
	return "Zero"
}

// длина строки
func stringlenght(line string) int {
	return len(line)
}

type Rectangle struct {
	width float64
	height  float64
}

// площадь прямоугольника
func (p Rectangle) S() float64 {
	return p.width * p.height
}

// среднее двух целых
func sredneeDvuh(chislo1, chislo2 int) float64 {
	return float64(chislo1+chislo2) / 2
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Задание 1
	fmt.Println("Задание 1.")
	fmt.Print("Введите число: ")
	line1, _ := reader.ReadString('\n')
	number1, error1 := strconv.Atoi(strings.TrimSpace(line1))
	if error1 != nil {
		fmt.Println("Нужно целое число")
	} else if number1%2 == 0 {
		fmt.Printf("Число %d чётное\n", number1)
	} else {
		fmt.Printf("Число %d нечётное\n", number1)
	}

	// Задание 2
	fmt.Println("\n Задание 2.")
	fmt.Print("Введите число: ")
	line2, _ := reader.ReadString('\n')
	number2, error2 := strconv.Atoi(strings.TrimSpace(line2))
	if error2 != nil {
		fmt.Println("Нужно целое число")
	} else {
		fmt.Println("Результат:", identifyingmark(number2))
	}

	// Задание 3
	fmt.Println("\n Задание 3.")
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Задание 4
	fmt.Println("\n Задание 4.")
	fmt.Print("Введите строку: ")
	line4, _ := reader.ReadString('\n')
	line4 = strings.TrimSpace(line4)
	fmt.Println("Длина строки:", stringlenght(line4))

	// Задание 5
	fmt.Println("\n Задание 5.")
	pryamoygolnik := Rectangle{width: 5.5, height: 3.2}
	fmt.Printf("Ширина = %.1f, высота = %.1f\n", pryamoygolnik.width, pryamoygolnik.height)
	fmt.Printf("Площадь = %.2f\n", pryamoygolnik.S())

	// Задание 6
	fmt.Println("\n Задание 6.")
	fmt.Print("Введите два целых числа через пробел: ")
	line6, _ := reader.ReadString('\n')
	part6 := strings.Fields(line6)

	if len(part6) < 2 {
		fmt.Println("Нужно два числа")
		return
	}

	number6a, error6a := strconv.Atoi(part6[0])
	nubmer6b, error6b := strconv.Atoi(part6[1])

	if error6a != nil || error6b != nil {
		fmt.Println("Оба числа должны быть целыми")
		return
	}

	fmt.Printf("Среднее значение = %.2f\n", sredneeDvuh(number6a, nubmer6b))
}
