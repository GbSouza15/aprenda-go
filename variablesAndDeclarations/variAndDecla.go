package variablesanddeclarations

import "fmt"

func User() {
	var nome string = "Gabriel"
	idade := 19

	var a, b, c, d, e int = 1, 2, 3, 4, 5

	x, y := 5, "Oi"

	fmt.Println("Olá, meu nome é", nome, "e eu tenho", idade, "anos")
	fmt.Println(a, b, c, d, e)
	fmt.Println(x, y)
}
