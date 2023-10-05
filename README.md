# Aprendendo Golang
## Documentação para aprender Go de maneira simples e direta!

Bem-vindo ao seu primeiro passo na jornada de aprendizado de Go!

## Índice

1. [Capítulo 1: Introdução a Sintaxe do Go](https://github.com/GbSouza15/aprenda-go#documenta%C3%A7%C3%A3o-para-aprender-go-de-maneira-simples-e-direta)
2. [Capítulo 2: Variáveis e declarações](https://github.com/GbSouza15/aprenda-go#vari%C3%A1veis-e-declara%C3%A7%C3%B5es-capitulo-2-variaveis-e-declara%C3%A7%C3%B5es)
2. [Capítulo 3: Tipos de dados](#capitulo-3-tipos-de-dados)

##

### Entendendo a Sintaxe do Go {#capitulo-1-introducao-a-sintaxe}

**Pacotes e a Função Principal:**

No Go, cada programa faz parte de um pacote. O pacote principal é chamado "main" e é onde a execução do programa começa. Aqui está um exemplo de um programa Go simples:

```go
package main

import "fmt"

func main() {
  fmt.Println("Olá, mundo!")
}
```

Neste código, observamos os seguintes elementos:

- `package main`: Declara que este programa faz parte do pacote "main".
- `import "fmt"`: Importa o pacote "fmt", que nos permite realizar operações de entrada e saída.
- `func main() { ... }`: Define a função principal do programa. Todo o código que está entre as chaves `{}` desta função será executado.

**Primeiro Passo: "Hello World!"**

A linha `fmt.Println("Olá, mundo!")` imprime a mensagem "Olá, mundo!" no terminal. Este é o famoso "Hello World" em Go.

Isso é apenas o começo! Ao longo desta documentação, exploraremos com mais detalhes a sintaxe do Go. Até breve!

**Próximos Passos**

Agora que você teve uma breve introdução ao Go, você pode continuar aprendendo com os próximos capítulos. Cada capítulo irá abordar tópicos específicos e avançar gradualmente em direção ao domínio do Go. Sinta-se à vontade para explorar os capítulos seguintes quando estiver pronto. Feliz aprendizado!

## Variáveis e declarações {#capitulo-2-variaveis-e-declarações}

**Tipos de variáveis**

No Go exitem vários tipos de variáveis:

`Int`

`string`

`bool`

Para saber sobre outros tipos acesse: [este recurso](https://www.w3schools.com/go/go_data_types.php)

**Sintaxe**

```go
var nomeDaVariavel tipo = valor
```

Você pode declarar variáveis usando essa sintaxe dentro ou fora de funções.

Agora veja essa sintaxe:

```go
nomeDaVariavel := valor
```

Nesse caso o valor da variável é inferido, ou seja, o tipo da variável será com base no valor que você irá passar.

**Mais algumas observações**

Você pode declarar uma variável sem passar um valor para ela. Ex:

```go
//Variável declarada sem valor
var valorDaVariavel tipo;

//Atribuindo um valor a variável
valorDaVariavel = valor
```

Mas, isso não funciona se você usar o: 
```go
:=
```

Não é possível declarar uma variável e não atribuir um valor para ela usando essa forma mais curta.

**Diferença entre var e :=**

Além do que já falamos sobre elas na parte de **Mais algumas observações**, temos outra diferença. Usando **var**, você pode usar tando dentro como fora de funções. Ex:

```go
package main

import "fmt"

var nome string = "Gabriel"

func main() {

  var sobrenome string = "Souza"

  fmt.Println(nome + sobrenome)

}
```

Já usando o **:=**, não é possível fazer isso. Ex:

```go
package main

import "fmt"

nome := "Gabriel" //Não pode ser usado fora da função

func main() {
  sobrenome := "Souza"

  fmt.Println(nome + sobrenome)
}
```

Ao fazer desse jeito, teremos um erro ao tentar executar nosso código:

```
syntax error: non-declaration statement outside function body
```

**Declaração de múltiplas variáveis**

No Go é possível declarar múltiplas variáveis na mesma linha. Ex:

```go
package main

import "fmt"

func main() {
  var a, b, c, d int = 1, 2, 3, 4

  fmt.Println(a, b, c, d)
}
```

Se todas as variáveis que você deseja declarar são do mesmo tipo, você pode fazer como o exemplo acima. Mas, e se eu quiser declarar múltiplas variáveis de tipos diferentes? Veja: 

```go
package main

import "fmt"

func main() {
  a, b, c := 1, "Hello", true

  fmt.Println(a, b, c)
}
```

É só ocultar o tipo e usar a forma reduzida para declarar variáveis.

## Tipos de dados {#capitulo-3-tipos-de-dados}

**Tipo booleano**

Um tipo booleano é declarado com a palavra chave bool e só pode ter o valor **true** ou **false**

```go
package main

import "fmt"

func main() {
  var b1 bool = true
  var b2 bool = false

  fmt.Println(b1) //true
  fmt.Println(b2) // false
}
```

**Tipos inteiros**

Tipos de dados inteiros são usados para armazenar números inteiros sem decimais. Ex:

```go
package main

import "fmt"

func main() {
  var n1 int = -30
  var n2 int = -9
  var n3 int = 50
  var n4 int = 300

  fmt.Println(n1) 
  fmt.Println(n2) 
  fmt.Println(n3) 
  fmt.Println(n4) 
}
```

Tipos de dados inteiros possui duas categorias:

- Inteiros assinados
- Inteiros não assinados

Os assinados podem armazenar valores positivos e negativos. Já os não assinados só armazenam valores positivos.

Para usar os não assinados, é usado essa sintaxe:

```go
package main

import "fmt"

func main() {
  var n1 uint = 10
  var n2 uint = 3
  var n3 uint = 90
  var n4 uint = -40 //Está errado, por ser um inteiro não assinado, não é possível atribuir um valor negativo a ele.
}
```

Para saber mais sobre tipos inteiros, acesse [este recurso](https://www.w3schools.com/go/go_integer_data_type.php)

**Tipos de ponto flutuante**

Usados para armazenar números com ponto decimal positivos e negativos. Ex:

```go
package main

import "fmt"

func main() {
  var n1 float64 = 12.80
  var n2 float64 = -50.99
}
```

Para saber mais sobre esse tipo, acesse [esse recurso](https://www.w3schools.com/go/go_float_data_type.php)

**Tipos string**

Servem para armazenar texto e devem usar **aspas duplas**. Ex:

```go
package main

func main() {
  var nome string = "Gabriel"
  var sobrenome string = "Souza"
}
```