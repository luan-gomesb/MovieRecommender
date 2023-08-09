package main

import (
	"fmt"
)

func main() {

	teste(worker)
}

func worker() {
	fmt.Println("worker working..")
}

func teste(a func()) {
	fmt.Println("function start")
	defer fmt.Println("function end")
	a()
}
