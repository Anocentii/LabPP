package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// сумма и разность двух float
func summaIRaznost(chislo1, chislo2 float64) (float64, float64) {
	return chislo1 + chislo2, chislo1 - chislo2
}

// среднее трёх чисел
func srednee(chislo1, chislo2, chislo3 float64) float64 {
	return (chislo1 + chislo2 + chislo3) / 3
}

func main() {
	// Задание 1
	tekusheeVremya := time.Now()
	fmt.Printf("Задание 1.\n")
	fmt.Println("Текущая дата и время:", tekusheeVremya.Format("02.01.2006 15:04:05."))

	// Задание 2
	var celoeChislo int = 25
	var drobnoeChislo float64 = 3.14
	var stroka string = "Привет, Anocentii"
	var logika bool = true
	fmt.Printf("\n Задание 2.\n")
	fmt.Println("int:", celoeChislo)
	fmt.Println("float64:", drobnoeChislo)
	fmt.Println("string:", stroka)
	fmt.Println("bool:", logika)

	// Задание 3 - через :=
	vozrast := 20
	rost := 172
	name := "Булат"
	student := true

	fmt.Printf("\n Задание 3.\n")
	fmt.Println("Возраст:", vozrast)
	fmt.Println("Рост:", rost)
	fmt.Println("Имя:", name)
	fmt.Println("Студент:", student)

	// Задание 4
	chisloA := 17
	chisloB := 5
	fmt.Printf("\n Задание 4.\n")
	fmt.Printf("%d + %d = %d\n", chisloA, chisloB, chisloA+chisloB)
	fmt.Printf("%d - %d = %d\n", chisloA, chisloB, chisloA-chisloB)
	fmt.Printf("%d * %d = %d\n", chisloA, chisloB, chisloA*chisloB)
	if chisloB != 0 {
		fmt.Printf("%d / %d = %d\n", chisloA, chisloB, chisloA/chisloB)
		fmt.Printf("%d %% %d = %d\n", chisloA, chisloB, chisloA%chisloB)
	}

	// Задание 5
	summa, raznost := summaIRaznost(7.5, 2.5)
	fmt.Printf("\n Задание 5.\n")
	fmt.Printf("Первое число = %.2f, Второе число = %.2f \n",summa, raznost)
	fmt.Printf("Сумма = %.2f, разность = %.2f\n \n", summa, raznost)

	// Задание 6
	fmt.Printf("\n Задание 6.\n")
	fmt.Println("Введите три числа через пробел:")
	reader := bufio.NewReader(os.Stdin)
	strokaVvoda, _ := reader.ReadString('\n')
	chasti := strings.Fields(strokaVvoda)

	if len(chasti) < 3 {
		fmt.Println("Нужно три числа")
		return
	}

	chisla := make([]float64, 3)
	for i := 0; i < 3; i++ {
		znachenie, oshibka := strconv.ParseFloat(chasti[i], 64)
		if oshibka != nil {
			fmt.Println("Не число:", chasti[i])
			return
		}
		chisla[i] = znachenie
	}

	fmt.Printf("Среднее: %.2f\n", srednee(chisla[0], chisla[1], chisla[2]))
}
