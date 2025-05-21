package main

import "fmt"

type Generator interface {
	Next() int
}

func doWork(g Generator) {
	n := g.Next()
	fmt.Println(n)
	//do something with n
}
