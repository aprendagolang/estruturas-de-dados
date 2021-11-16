package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
	Sexo      string
}

func main() {
	pessoas := []Pessoa{
		{"Maria", "da Graça", 50, "F"},
		{"Maria", "da Silva", 45, "F"},
		{"Tiago", "Temporin", 31, "M"},
		{"João", "Santos", 24, "M"},
	}

	table := HashTable{}

	keys := make([]int, len(pessoas))
	for i, pessoa := range pessoas {
		keys[i] = table.Put(pessoa)
	}

	for _, key := range keys {
		ps := table.Get(key)
		for _, p := range ps {
			fmt.Println(p.Nome, p.Sobrenome)
		}
	}

	table.Remove("Maria")

	tiago := table.Search("Tiago")
	fmt.Println(tiago)
}
