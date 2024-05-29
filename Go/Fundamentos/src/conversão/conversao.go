package main

import (
	"fmt"
	"strconv"
)

var a int = 20
var b string = "name"

func main() {
	var c float64 = float64(a)
	var d string = strconv.Itoa(a)
	fmt.Println(c, d)
}