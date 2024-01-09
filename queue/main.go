package main

import "fmt"

func main() {
	queue := Queue{}
	queue.Enqueue("Tiago")
	queue.Enqueue("Dani")
	queue.Enqueue("Maria")
	queue.Enqueue("Pedro")
	queue.Enqueue("José")

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}
