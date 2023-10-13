package array

import "fmt"

func ThisArray() {
	arr1 := []int{1, 2, 3, 4, 5}

	fmt.Println(arr1)

	for i, v := range arr1 {
		fmt.Printf("Índice no Array: %d, Valor: %d\n", i, v)
	}
}
