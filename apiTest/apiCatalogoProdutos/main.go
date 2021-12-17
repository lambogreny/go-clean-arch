package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	numberOfWorkers int = 8
)

func main() {

	//Criando canal de comunicação com o worker
	channel := make(chan int)

	//Criando uma função assíncrona que vai executar o worker
	go func() {
		//Criando 5 workers(processos)
		for i := 1; i <= numberOfWorkers; i++ {
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
		//Aqui dentro que faz o processo
		url := "http://172.30.0.77:8883/api/v1/catalog/categories?teste=asdasd"

		payload := strings.NewReader("{\n\t\"Teste\":\"asdasdasd\"\n}")

		req, _ := http.NewRequest("GET", url, payload)

		req.Header.Add("Content-Type", "application/json")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		fmt.Println(res)
		fmt.Println(string(body))

		time.Sleep(time.Second * 5)
	}
}
