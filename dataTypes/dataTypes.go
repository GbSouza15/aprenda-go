package datatypes

import "fmt"

func DataTypes() {
	var number int = -10
	var nome string = "Gabriel"
	var boolean bool = true
	var number2 uint = 10

	fmt.Printf("Type %T, value %v", number, number)
	fmt.Printf("Type %T, value %v", nome, nome)
	fmt.Printf("Type %T, value %v", boolean, boolean)
	fmt.Printf("Type %T, value %v", number2, number2)
}
