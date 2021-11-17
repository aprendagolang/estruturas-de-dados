package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int8
}

func main() {
	list := List{}

	tiago := Pessoa{"Tiago", "Temporin", 31}
	joao := Pessoa{"João", "Santos", 24}
	dani := Pessoa{"Dani", "Almeida", 54}
	maria := Pessoa{"Maria", "Santos", 34}
	marcos := Pessoa{"Marcos", "Augusto", 44}

	list.Append(tiago)
	list.Append(dani)
	list.Append(joao)
	list.Append(maria)
	list.Append(marcos)

	list.Display()

	fmt.Println("--------------------")

	pessoa := list.Search("Maria")
	fmt.Println(pessoa)

	fmt.Println("--------------------")

	list.Delete("Dani")
	list.Display()
}
