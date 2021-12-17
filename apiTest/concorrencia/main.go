package main

import (
	"fmt"
	"time"
)

func main() {

	//Criando canal de comunicação com o worker
	channel := make(chan int)

	//Criando uma função assíncrona que vai executar o worker
	go func() {
		//Criando 5 workers(processos)
		for i := 1; i <= 50; i++ {
			go worker(channel)
		}
	}()

	for i := 0; i < 100; i++ {
		//Mandando informação para o canal
		channel <- i
	}

}

/**
 * Função de worker que recebe a mensageem do canal
 */
func worker(channel chan int) {

	for i := range channel {
		fmt.Println(i)
		time.Sleep(time.Second * 5)
	}
}
