package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	defaultOperator = "+-/*"
)

type Operation struct {
	firstNum  int
	lastNum   int
	operator  string
	isRim     bool
	result    int
	resultRim string
}

func main() {
	p := Operation{}
	str := input()
	slice := split(str)
	checkNum(&p, slice)
	validator(&p, slice)
	operation(&p)
	if p.isRim {
		if p.result <= 0 {
			log.Fatal("В римских числах нет отрицательных")
		}
		fmt.Println(toRim(p.result))
	} else {
		fmt.Println(p.result)
	}

}

func input() string {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	str := s.Text()
	return str
}

func split(str string) []string {
	slice := strings.Split(str, " ")
	if len(slice) != 3 {
		log.Fatal("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
	return slice
}

func checkNum(o *Operation, slice []string) {
	rim := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	s, ok := rim[slice[0]]
	if !ok {
		o.isRim = false
		o.firstNum, _ = strconv.Atoi(slice[0])
		o.lastNum, _ = strconv.Atoi(slice[2])
		return
	}
	o.firstNum = s
	s, _ = rim[slice[2]]
	o.lastNum = s
	o.isRim = true
}

func validator(o *Operation, slice []string) {
	var err error
	if !strings.Contains(defaultOperator, slice[1]) && len(slice) != 1 {
		log.Fatal("Невалидный оператор")
	}
	o.operator = slice[1]
	if err != nil || o.firstNum > 10 || o.firstNum < 1 {
		log.Fatal("Неверное первое число")
	}
	if err != nil || o.lastNum > 10 || o.lastNum < 1 {
		log.Fatal("Неверное второе число")
	}

}

func operation(o *Operation) {
	switch o.operator {
	case "+":
		o.result = o.firstNum + o.lastNum
	case "-":
		o.result = o.firstNum - o.lastNum
	case "*":
		o.result = o.firstNum * o.lastNum
	case "/":
		o.result = o.firstNum / o.lastNum
	}
}

func toRim(result int) string {
	switch {
	case result < 0:
		return "в римской системе нет отрицательных чисел."
	case result > 100:
		return "Результат больше 100, не подходит под условия задачи"
	case result == 0:
		return ""
	case result < 4:
		return "I" + toRim(result-1)
	case result == 4:
		return "IV"
	case result < 9:
		return "V" + toRim(result-5)
	case result == 9:
		return "IX"
	case result < 40:
		return "X" + toRim(result-10)
	case result < 50:
		return "XL" + toRim(result-40)
	case result < 90:
		return "L" + toRim(result-50)
	case result < 100:
		return "XC" + toRim(result-90)
	case result == 100:
		return "C"
	}
	return ""
}
