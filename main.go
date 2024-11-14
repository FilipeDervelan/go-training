package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	name string
	number1, number2 int
	cara, coroa int
)

func hello(name string) {
	fmt.Println("Hello", name, "!")
}

func sum(a, b int) int {
	return a + b 
}

func convertAndSum(a int, b string) (total int, err error) {
	i, err := strconv.Atoi(b)

	if err != nil {
		return
	}

	total = a + i

	return
}

func sumAndGreeting(num1, num2 int, name string) string {
	var message string

	mySum := sum(num1, num2)

	numConverted := strconv.Itoa(mySum)

	message = fmt.Sprintf("Hello %s! The result of the sum is %s", name, numConverted)

	return message
}

func conditionals() {
	a, b := 20, 15

	if a > b {
		fmt.Println("a é maior que b")
	} else if a < b {
		fmt.Println("a é menor que b")
	} else {
		fmt.Println("a e b são iguais")
	}

	switch {
	case a > b:
		fmt.Println("a é maior que b")

	case a < b:
		fmt.Println("a é menor que b")

	default:
		fmt.Println("a e b são iguais")
	}

	file, err := os.Open("hello.txt")
	if err != nil {
		log.Panic(err)
	}

	data := make([]byte, 100)
	if _, erro := file.Read(data); erro != nil {
		log.Panic(erro)
	}

	fmt.Println(string(data))
}

func lancarMoeda(lado string) {
	switch lado {
	case "cara":
		cara++

	case "coroa":
		coroa++

	default:
		fmt.Println("caiu em pé")
	}
}

func looping() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	nomes := []string{"Filipe", "Mirela", "André", "Sara"}

	for i := 0; i < len(nomes); i++ {
		fmt.Println(nomes[i])
	}

	fmt.Println("")

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	fmt.Println("")

	var i int
	for i < len(nomes) {
		fmt.Println(nomes[i])
		i++
	}
}

func slices() {
	// nomes := []string{"Filipe", "Mirela", "André", "Sara"}
	// fmt.Println(nomes, len(nomes), cap(nomes))

	// nomes = append(nomes, "Ana")
	// fmt.Println(nomes, len(nomes), cap(nomes))

	// nomes = append(nomes, "Sergio")
	// fmt.Println(nomes, len(nomes), cap(nomes))

	// nomes := make([]string, 10, 20)
}

func maps() {
	idades := make(map[string]uint8)
	idades["Filipe"] = 19
	idades["Mirela"] = 20
	idades["André"] = 31

	fmt.Println(idades)
	fmt.Println(idades["Filipe"])
	fmt.Println(idades["Sara"])

	val, ok := idades["Filipe"]
	fmt.Println(val, ok)
}

type Pessoa struct {
	Nome string
	Sobrenome string
	Idade uint8
	Status bool
}

func structs() {
	// var p Pessoa
	// p.Nome = "Filipe"
	// p.Sobrenome = "Dervelan"
	// p.Idade = 19
	// p.Status = true

	// p := Pessoa {
	// 	Nome: "Filipe",
	// 	Sobrenome: "Dervelan",
	// 	Idade: 19,
	// 	Status: true,
	// }

	// Não recomendado, porque se a struct for alterada não funcionará.
	p := Pessoa {"Filipe", "Dervelan", 16, true}

	fmt.Println(p)
	fmt.Println(p.Nome)
	fmt.Println(p.Idade)
}

func main() {
	// hello("Filipe Dervelan")

	// fmt.Println("Total:", sum(10,10))

	// total, err := convertAndSum(10, "40")
	// fmt.Println("Total:", total, err)

	// name := "Filipe Dervelan"

	// fmt.Println(sumAndGreeting(12, 12, name))

	// conditionals()
	// lancarMoeda("a")

	// looping()

	// slices()
	// maps()
	structs()
}