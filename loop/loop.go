package loop

import "fmt"

func Loop() {
	for i := 0; i < 5; i++ {
		// if i == 3 {
		// 	break
		// }
		// fmt.Println(i)
		if i == 3 {
			panic("kkkkkkkkkkk")
		}
	}

	var numbers = []int{1, 2, 3, 4, 5}

	for i, v := range numbers {
		fmt.Println(i, v)
	}
}
