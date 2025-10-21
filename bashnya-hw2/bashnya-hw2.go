// Takes a number, from -max to max, does some strange calculations, checking the number for certain features
// Translates the resulting number into correct words and expressions in Russian.
// Prints words up to 999K

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	)
	
var max_int = 112307
	
// Checks the input for validity (a single int in range without trailing garbage)
func checker(inp string) (int, bool) {

	fmt.Println(" ")
	check := false
	inp = strings.TrimSpace(inp)
	num, err := strconv.Atoi(inp)
	if err != nil {
		fmt.Println(" Неверный ввод")
		check = true
	} else {
		fmt.Printf(" Целое число: %d\n", num)
	}
	
	if (num < -max_int || num > max_int) {
		fmt.Println(" Должно быть число от -12307 до 12307")
		check = true
	} else {
		fmt.Printf(" Попадает в диапазон: -12307 <= %d >= 12307\n", num)
	}
	return num, check
}

// Do calculations according to the task
func calc(num int) int {
	fmt.Printf("\n——————————————————————————————————————————————————\n\n   Число: %d. Расчёт: \n--------------------------------------------------\n", num)
	
	for num < max_int {
		
		if num < 0 {
			num *= -1 
			fmt.Printf(" Отрицательное, умножить на -1: %d\n", num)
		} else if num % 7 == 0 { 
			fmt.Printf(" %d делится без остатка на 7, умножить на 39: %d * 39 = %d\n", num, num, num * 39)
			num *= 39 
		} else if num % 9 == 0 { 
			fmt.Printf(" %d делится без остатка на 9, умножить на 13, прибавить 1: (%d * 13) + 1 = %d\n", num, num, num * 13 + 1)
			num = num * 13 + 1 
			continue
		} else {
			fmt.Printf(" %d не соответствует другим условиям, прибавить 2, умножить на 3: (%d + 2) * 3 = %d\n", num, num, (num + 2) * 3)
			num = (num + 2) * 3
		}
		
		if num % 9 == 0 && num % 13 == 0 {
			fmt.Printf(" Второй цикл, %d делится без остатка на 9 И на 13, error\n", num)
			fmt.Println(" service error ")
			num = 0
			break
		} else { 
			num += 1
			fmt.Printf(" Второй цикл, прибавить 1: %d\n", num)			
		}
	}
	if num != 0 { fmt.Printf(" %d >= %d, Обработка всё!", num, max_int) }
	return num
}	

//————————————————————————————————————————\Перевод в слова\————————————————————————————————————————\

// Translate 1 - 9 into words
func trans_one_nine(num_one_nine int) string {
	one_nine := ""
	switch num_one_nine {
		case 1:
		one_nine += "один "
		case 2:
		one_nine += "два"
		case 3:
		one_nine += "три"
		case 4:
		one_nine += "четыре"
		case 5:
		one_nine += "пять"
		case 6:
		one_nine += "шесть"
		case 7:
		one_nine += "семь"
		case 8:
		one_nine += "восемь"
		case 9:
		one_nine += "девять"
	}
	return one_nine
}

// Translate 10 - 19 into words
func trans_teens(num_teen int) string {
	teen := ""
	switch num_teen {
		case 10:
		teen = "десять"
		case 11:
		teen = "одиннадцать"
		case 12:
		teen = "двенадцать"
		case 13:
		teen = "тринадцать"
		case 14:
		teen = "четырнадцать"
		case 15:
		teen = "пятнадцать"
		case 16:
		teen = "шестнадцать"
		case 17:
		teen = "семнадцать"
		case 18:
		teen = "восемнадцать"
		case 19:
		teen = "девятнадцать"
	}
	return teen
	
}

// Translate 20 - 90 into words
func trans_tens(num_tens int) string {
	tens := ""
	switch num_tens / 10 {
		case 2:
			tens = "двадцать"
		case 3:
			tens = "тридцать"
		case 4:
			tens = "сорок"
		case 5:
			tens = "пятьдесят"	
		case 6:
			tens = "шестьдесят"
		case 7:
			tens = "семьдесят"
		case 8:
			tens = "восемьдесят"
		case 9:
			tens = "девяносто"
	}
	//if num_tens % 10 != 0 {
	//	tens = tens + " " + trans_one_nine(num_tens % 10)
	//}
	return tens
}

// Translate 100 - 900 into words
func trans_hund(num int) string {
	hund := ""
	switch num /100 {
		case 1:
			hund += "сто"
		case 2:
			hund += "двести"
		case 3:
			hund += "триста"
		case 4:
			hund += "четыреста"
		case 5:
			hund += "пятьсот"
		case 6:
			hund += "шестьсот"
		case 7:
			hund += "семьсот"
		case 8:
			hund += "восемьсот"
		case 9:
			hund += "девятьсот"
	}
	return hund
}

// Word selection for "thousands"
func thous_print(num_th int) string {
	
	thous := ""
	if num_th != 11 && num_th % 10 == 1 {
		thous += "тысяча"
	} else if (num_th % 10 >= 2 && num_th % 10 <= 4) && !(num_th >= 10 && num_th < 20) {
		thous += "тысячи"
	} else {
		thous += "тысяч"
	}
	return thous
}

// Thousands processing (capable of 1-999K)
func trans_thous(num_th int) string {
	
	thous := ""
	if num_th >= 100 {
		thous = thous + trans_hund(num_th) + " "
		fmt.Printf (" Извлекаем сотни тысяч: %d ==> %s\n", num_th, trans_hund(num_th))		
		fmt.Printf (" Остаток от деления на 100 (1-99, тыс.): %d\n", num_th % 100)
		num_th = num_th % 100
	}
		
	if num_th >= 20 && num_th < 100 {
		num_th_tnth := num_th % 10
		thous = thous + trans_tens(num_th) + " "
		fmt.Printf (" Извлекаем десятки тысяч (20-99): %d ==> %s\n", num_th, trans_tens(num_th))
		if num_th_tnth > 0 { fmt.Printf (" Остаток от деления на 10 (1 - 9 тыс.): %d", num_th_tnth) }
		num_th = num_th_tnth
		
	} else if num_th >= 10 && num_th <= 19 {
		thous = thous + trans_teens(num_th)
		fmt.Printf (" Если > 10, < 20 тысяч: %d ==> %s\n", num_th, trans_teens(num_th))
	} 
	
	if num_th == 1 {
		fmt.Println ("... одна тысяча")
		thous += "одна"
	} else if num_th == 2 {
		fmt.Println ("... две тысячи")
		thous += "две"				
	} else if num_th > 2 && num_th < 10 {
		thous += trans_one_nine(num_th)
		fmt.Printf ("... %s тыс.\n", trans_one_nine(num_th)) 
		
	}
	return thous
}

func translator(num int) string {
	fmt.Printf ("\n——————————————————————————————————————————————————\n   Результат: %d. Перевод в слова:\n--------------------------------------------------\n", num)
	
	
	result := ""
	if num >= 1000 {
		num_th := num/1000
		result += trans_thous(num_th)
		result = result + " " + thous_print(num_th)
		fmt.Printf(" Остаток от деления на 1000 (1-999): %d ==> %d\n", num, num % 1000)
		num = num % 1000
	} 
	
	if num >= 100 {
		result = result + " " + trans_hund(num)
		fmt.Printf(" Сотни: %d %s\n", num, trans_hund(num) )
		num = num % 100
	}
		
	if num >= 20 && num < 100 {
		result = result + " " + trans_tens(num)
		fmt.Printf(" Десятки (20 <= n < 100): %d %s\n", num, trans_tens(num))
		num = num % 10
	} else if num >= 10 && num < 20 {
		result = result + " " + trans_teens(num)
		fmt.Printf(" Десятки, (10 <= n < 20): %d %s\n", num, trans_teens(num))
	}
	
	if num >= 1 && num < 10 {
		result = result + " " + trans_one_nine(num)
		fmt.Printf(" Единицы (1 <= n < 10): %d %s\n", num, trans_one_nine(num))
	}

	return result
}

func main() {
	num, check := 0, false
	input := bufio.NewReader(os.Stdin)
	const head_lngth = 30
	header := "\nW" + strings.Repeat("=", head_lngth) + "~~~START~~~" + strings.Repeat("=", head_lngth) + "W\n"
	fmt.Print(header + "\n   Пожалуйста, введите число от -12307 до 12307: ")
	inp, err := input.ReadString('\n')
	if err != nil {
		fmt.Println(" Error: ", err)
	} else { 
		num, check = checker(inp) 
	}
	
	if !check {
		crazy_shit := calc(num)
		
// Uncomment raws below to run the MULTIPLE INPUT
// It will process every n-th i. You can set it in the "for" loop.
		
		/*for i := 1; i <= max_int; i+=256 {
			fmt.Println(" ")
			num = i
			crazy_shit := calc(num) */
			
			if crazy_shit == 0 {
				fmt.Printf(" Выход\n\nM=============================== END ===============================M\n")
				
		//		continue   // uncomment this 'continue' for the loop
		
			} else {
			
				fmt.Printf("\n--------------------------------------------------\n  Число ==> результат: %d ==> %d\n", num, crazy_shit)
				translation := translator(crazy_shit)
				deconstruct := strings.Fields(translation)
				repair := strings.Join(deconstruct, " ")
				fmt.Printf("\n--------------------------------------------------\n  Результат: %d; Перевод: %s \n\nM=============================== END ===============================M\n", crazy_shit, repair)
			}
			
		// '}' close brackets here
	}
}