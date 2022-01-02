package main //Mudar aqui!!

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
)

func task() {
	fmt.Println("I am running task.")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func main() {
	s1 := gocron.NewScheduler()
	s1.Every(2).Second().Do(taskWithParams, 1, "hello1")
	<-s1.Start()

	s2 := gocron.NewScheduler()
	s2.Every(2).Second().Do(taskWithParams, 2, "hello2")
	<-s2.Start()

	<-gocron.Start()

	// taskWithParams(1, "hello")
	// gocron.Every(1).Second().Do(taskWithParams, 1, "hello")
	// <-gocron.Start()
}
