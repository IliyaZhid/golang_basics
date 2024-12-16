package main

import (
	"fmt"
	"strings"
)

func main() {
	const delicatesMaxPrice = 500
	var builder strings.Builder
	builder.WriteString("Список деликатесов:\n")

	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	orderPrice := 0

	prices := map[string]int{
		"хлеб":     50,
		"молоко":   100,
		"масло ":   200,
		"колбаса":  500,
		"соль":     20,
		"огурцы":   200,
		"ветчина":  700,
		"буженина": 900,
		"помидоры": 250,
		"рыба":     300,
		"хамон":    1500,
		"сыр":      600,
	}

	for key, value := range prices {
		if value > delicatesMaxPrice {
			builder.WriteString(fmt.Sprintf("- %s\n", key))
		}
	}

	for _, value := range order {
		orderPrice += prices[value]
	}

	fmt.Println(builder.String())

	fmt.Println("Стоимость заказа:", orderPrice, "рублей")
}
