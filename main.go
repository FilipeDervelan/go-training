package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
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
	nomes := []string{"Filipe", "Mirela", "André", "Sara"}
	fmt.Println(nomes, len(nomes), cap(nomes))

	nomes = append(nomes, "Ana")
	fmt.Println(nomes, len(nomes), cap(nomes))

	nomes = append(nomes, "Sergio")
	fmt.Println(nomes, len(nomes), cap(nomes))

	nomes = make([]string, 10, 20)
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

type Document interface {
	Doc() string
}

type Pessoa struct {
	Nome string
	Idade uint8
	Status bool
}

type PessoaFisica struct {
	Pessoa
	Sobrenome string
	cpf string
}

func (pf PessoaFisica) Doc() string {
	return fmt.Sprintf("meu cpf é: %s", pf.cpf)
}

type PessoaJuridica struct {
	Pessoa
	RazaoSocial string
	cnpj string
}

func (pj PessoaJuridica) Doc() string {
	return fmt.Sprintf("meu cnpj é: %s", pj.cnpj)
}

func (p Pessoa) String() string  {
	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos.", p.Nome, p.Idade)
}

func show(d Document) {
	switch d.(type) {
	case PessoaFisica:
		fmt.Println(d.(PessoaFisica).Sobrenome)

	case PessoaJuridica:
		fmt.Println(d.(PessoaJuridica).RazaoSocial)

	default:
		fmt.Println("tipo desconhecido")
	}

	fmt.Println(d)
	fmt.Println(d.Doc())
}

func structs() {
	var p Pessoa
	p.Nome = "Filipe"
	// p.Sobrenome = "Dervelan"
	p.Idade = 19
	p.Status = true

	// p := Pessoa {
	// 	Nome: "Filipe",
	// 	Sobrenome: "Dervelan",
	// 	Idade: 19,
	// 	Status: true,
	// }

	// Não recomendado, porque se a struct for alterada não funcionará.
	// p := Pessoa {"Filipe", "Dervelan", 16, true}

	fmt.Println(p)
}

// Quando o channel fica tipado da maneira abaixo, significa que a função apenas
// faz a escrita no channel e não a leitura
func numeros(done chan <- bool) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 150)
	}

	// Escrevendo no channel: true
	done <- true
}

// Quando o channel fica tipado da maneira abaixo, significa que a função apenas
// faz a escrita no channel e não a leitura
func letras(done chan <- bool) {
	for l := 'a'; l < 'j'; l++ {
		fmt.Printf("%c ", l)
		time.Sleep(time.Millisecond * 230)
	}

	// Escrevendo no channel: true
	done <- true
}

func goRoutinesAndChannels() {
	// Quando estamos apontando a flecha para o Channel, significa que estamos
	// escrevendo algo nele
	// E quando estamos colocando a flecha e depois o channel, significa que
	// estamos lendo ele

	cn := make(chan bool)
	cl := make(chan bool)

	go numeros(cn)
	go letras(cl)
	
	// Lendo o channel
	<-cn
	<-cl

	fmt.Println("Fim da execução")
}

func numerosBuffer(n chan <- int) {
	for i := 0; i < 10; i++ {
		n <- i
		fmt.Printf("Escrito no channel: %d\n", i)
		time.Sleep(time.Millisecond * 90)
	}
	fmt.Println("Fim da escrita")
	close(n)
}

func goRoutinesAndChannelsBuffer() {
	cn := make(chan int, 3)
	go numerosBuffer(cn)
	
	for v := range cn {
		fmt.Printf("lido do channel: %d\n", v)
		time.Sleep(time.Millisecond * 180)
	}

	// Lendo o channel
	<-cn

	fmt.Println("Fim da execução")
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
	// structs()

	// pf := PessoaFisica{
	// 	Pessoa{Nome: "Filipe", Idade: 19, Status: true},
	// 	"Dervelan",
	// 	"000.000.000-00",
	// }	

	// show(pf)
	
	// fmt.Println()

	// pj := PessoaJuridica{
	// 	Pessoa{Nome: "Filipe", Idade: 19, Status: true},
	// 	"Dervelan LTDA",
	// 	"00.000.000/0000-00",
	// }

	// show(pj)

	// below i will use go routines
	// goRoutinesAndChannels()

	goRoutinesAndChannelsBuffer()
}