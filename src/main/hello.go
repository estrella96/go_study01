package main
import (
	"fmt"
)

const(
	PI float64=3.14
	NUM int64=2
	Name string="joe"
)

var(
	a float64=3.14
	num int64=5
	sex string="female"
)

type Gopher struct {

}
func(gopher Gopher) Go(){

}

type Writer interface {

}

type Int int32
func main()  {
	//go run hello.go
	var abc Int=1
	var a float64=1.1
	var b int
	b=int(a)
	c:=int(a)
	fmt.Println("hello,world")
	fmt.Println(abc,b,c)

}