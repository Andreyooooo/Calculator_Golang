package main

//Если удалить строки с номерами 172-175 и 195-198,
//то все будет работать и с числами больше 10, но не меньше 0
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

//из арабских в римские
func arabicToRomain(arabic_num int) string {
	var e_num, d_num, s_num, t_num int
	e_num = arabic_num % 10
	d_num = arabic_num / 10 % 10
	s_num = arabic_num / 100 % 10
	t_num = arabic_num / 1000
	edi := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	des := []string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	sot := []string{"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tis := []string{"M", "MM", "MMM", "MMMM"}
	var ans string
	if t_num != 0 {
		ans += tis[t_num-1]
	}
	if s_num != 0 {
		ans += sot[s_num-1]
	}
	if d_num != 0 {
		ans += des[d_num-1]
	}
	if e_num != 0 {
		ans += edi[e_num-1]
	}
	return ans
}

//из римских в арабские
func romainToArabic(romain_num string) int {
	years_map := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	ln := utf8.RuneCountInString(romain_num)

	//Отсекаем то, что не должно обрабатываться в цикле
	if ln == 1 {
		return (years_map[string(romain_num[0])])
	}

	//положим в ответ крайнее значение, потому что оно однозначно не может вычитаться
	var ans = (years_map[string(romain_num[ln-1])])

	//Логический параметр ласт_оперэйшон, чтобы понимать вычитать или добавлять при равенстве
	//соседних символов. если ранее вычли, то продолжим, аналогично со сложением
	var last_op bool
	if (years_map[string(romain_num[ln-1])]) >= (years_map[string(romain_num[ln-2])]) {
		last_op = true
	} else {
		last_op = false
	}

	for i := ln - 1; i != 0; i-- {
		if (years_map[string(romain_num[i-1])]) > (years_map[string(romain_num[i])]) {
			ans += (years_map[string(romain_num[i-1])])
			last_op = true
		} else if (years_map[string(romain_num[i-1])]) < (years_map[string(romain_num[i])]) {
			ans -= (years_map[string(romain_num[i-1])])
			last_op = false
		} else if (years_map[string(romain_num[i-1])]) == (years_map[string(romain_num[i])]) {
			if last_op == true {
				ans += (years_map[string(romain_num[i-1])])
			} else {
				ans -= (years_map[string(romain_num[i-1])])
			}
		}
	}
	return ans
}

func main() {
	// список разрешенных символов, которые могут встретиться в строке
	allowed_symbols := map[string]int{
		"I": 1, "V": 1, "X": 1, "L": 1, "C": 1, "D": 1,
		"M": 1,
		"0": 3, "1": 3, "2": 3, "3": 3, "4": 3, "5": 3,
		"6": 3, "7": 3, "8": 3, "9": 3, " ": 3,
		"+": 2, "-": 2, "*": 2, "/": 2}

	// счетчик всех встретившихся операторов
	//должен быть = 1, иначе ошибка
	var operation_counter int
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	for _, value := range text {
		key := allowed_symbols[string(value)]
		if key == 0 {
			fmt.Println("Ошибка. В строке присутствует неизвестный символ =", string(value))
			os.Exit(1)
		} else if key == 2 {
			operation_counter++
		}
	}
	if operation_counter > 1 {
		fmt.Println("В строке не может быть более одного оператора")
		os.Exit(1)
	} else if operation_counter == 0 {
		fmt.Println("Строка не является математической операцией")
		os.Exit(1)
	}
	//Так как строка будет разбиваться по пробелам, проверим случаи, когда
	//пользователь вводит 2+ 2, или 2  + 2. В общем, любое кривое использование пробела - ошибка
	if strings.Contains(text, "  ") {
		fmt.Println("В строке не должно содеражиться двойного пробела")
		os.Exit(1)
	}

	//Сплитим.
	//Если длина списка < 3, значит забыт пробел
	spisok := strings.Split(text, " ")
	if len(spisok) < 3 {
		fmt.Println("Между операторами и операндами необходимо ставить пробел")
		os.Exit(1)
	}
	var val_1 = spisok[0]
	type_val_1 := allowed_symbols[string(val_1[0])]

	var val_2 = spisok[2]
	type_val_2 := allowed_symbols[string(val_2[0])]

	var operation = spisok[1]

	//Если в одном отрезки встретились значения 1 и 3
	//(при обращении в allowed_symbols),
	//значит пользователь передает что-то вроде 2IX23 -> выдаем ошибку
	for _, v := range val_1 {
		if allowed_symbols[string(v)] != type_val_1 {
			fmt.Println("В одном операнде не могут присутствовать арабские и римские цифры")
			os.Exit(1)
		}
	}
	for _, v := range val_2 {
		if allowed_symbols[string(v)] != type_val_2 {
			fmt.Println("В одном операнде не могут присутствовать арабские и римские цифры")
			os.Exit(1)
		}
	}
	//Фиксируем тип левого и правого операнда - это параметры type_val_1 и type_val_2 соответственно
	//Если они расходятся, значит вводится что-то типо 2 - I -> ошибка
	if type_val_1 != type_val_2 {
		fmt.Println("Оба операнда должны быть из одной системы счисления")
		os.Exit(1)
	}
	//Сразу отловим деление на 0, чтобы потом не забыть
	//с этого момента type_val_1 = type_val_2, поэтому оперируем только одним
	if string(operation) == "/" {
		if type_val_1 == 3 {
			if utf8.RuneCountInString(val_2) == 1 {
				if string(val_2[0]) == "0" {
					fmt.Println("На ноль делить нельзя!")
					os.Exit(1)
				}
			}
		}
	}

	//обрабатываем римские числа
	if type_val_1 == 1 {
		var res int
		var res_1 = romainToArabic(val_1)
		var res_2 = romainToArabic(val_2)
		if res_1 < 1 || res_1 > 10 || res_2 < 1 || res_2 > 10 {
			fmt.Println("Числа больше 10 и меньше 0 не принимаются")
			os.Exit(1)
		}
		if string(operation) == "+" {
			res = res_1 + res_2
		} else if string(operation) == "-" {
			res = res_1 - res_2
			if res <= 0 {
				fmt.Println("Результат разности римских чисел не может быть меньше <= 0")
				os.Exit(1)
			}
		} else if string(operation) == "*" {
			res = res_1 * res_2
		} else if string(operation) == "/" {
			res = res_1 / res_2
		}
		var final_res = arabicToRomain(res)
		fmt.Println("Результат вычисления", final_res)
		//Обрабатываем арабские числа
	} else if type_val_2 == 3 {
		var res_1, _ = strconv.Atoi(val_1)
		var res_2, _ = strconv.Atoi(val_2)
		if res_1 < 1 || res_1 > 10 || res_2 < 1 || res_2 > 10 {
			fmt.Println("Числа больше 10 и меньше 0 не принимаются")
			os.Exit(1)
		}
		if string(operation) == "+" {
			fmt.Println("Результат вычисления", res_1+res_2)
		} else if string(operation) == "-" {
			fmt.Println("Результат вычисления", res_1-res_2)
		} else if string(operation) == "*" {
			fmt.Println("Результат вычисления", res_1*res_2)
		} else if string(operation) == "/" {
			fmt.Println("Результат вычисления", res_1/res_2)
		}

	}
}
