package conditionals

import "fmt"

func Conditionals() {
	a := 10
	b := 20

	if a > b {
		fmt.Println(a, " é maior que ", b)
	} else {
		fmt.Println(b, " é maior que ", a)
	}

	num := 1

	switch num {
	case 1:
		fmt.Println("O número escolhido foi 1")
	case 2:
		fmt.Println("O número escolhido foi 2")
	}
}
