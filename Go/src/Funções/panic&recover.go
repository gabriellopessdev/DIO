// panic: erro so programador
// recover: ela interrompe o panic
package main

import "fmt"

func main() {
	defer func ()  {
	
	x := recover()
	fmt.Println(x)
	}()
	panic("Pânico")
	
}