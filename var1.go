package main

import "fmt"

func main() {
	var num int
	fmt.Println("Введите целое число:")
	fmt.Scanf("%d\n", &num)

	if num > 12307 {
		fmt.Println("Число слишком большое, попробуйте меньшее значение.")
		return
	}

	for num < 12307 {
		if num < 0 {
			num = num * -1
		} else if num%7 == 0 {
			num = num * 39
		} else if num%9 == 0 {
			num = num*13 + 1
			continue
		} else {
			num = (num + 2) * 3
		}
		if num%13 == 0 && num%9 == 0 {
			break
			fmt.Println("service error")
		} else {
			num += 1
		}
	}

	fmt.Printf("Результат обработки числа: %d\n", num)
}
