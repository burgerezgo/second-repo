package main

import (
	"fmt"
)

func main() {
	fmt.Println("first line here")

	beyondhello()

}

func beyondhello() {
    var uid int
    uid=100
	sum, prod := third()
    total := fourth(uid)
	fmt.Println(sum, prod)
    fmt.Println(total)
    fmt.Println(fourth(uid))
}

func third() (summ, prodd int) {
	x := 3
	y := 4

	z := x + y
	a := x * y
	return z, a

}

func fourth(id int) (bsdfdfdfdfdfdfdftotal int){
    var hh, tt int
    tt=100
    hh=id+tt
    return hh
}
