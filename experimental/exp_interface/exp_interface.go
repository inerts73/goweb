package main

import (
	"fmt"
	"github.com/inerts73/myGo/experimental/exp_interface/mymath"
)

type echoEr interface {
	echoName(name string) 
}

type person string

func (p person) echoName(name string){
	fmt.Println(name)
}

func add() int {
	return 123
}

func main(){
	p := person("ADAM")
	var e echoEr = p
	e.echoName("ADAM HSU")

	type (
		s1 string
		s2 string
		f1 func() int
	)


	name1 := s1("S1")
	name2 := s2("S2")
	fmt.Println(name1, name2)
	f := f1(add)
	fmt.Println(f())

	fmt.Println(mymath.AddInt(3,9))
}