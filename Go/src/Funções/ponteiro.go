// ponterio (Ptr) armazena um valor na memoria
package main

import "fmt"

func inicial(xPtr *int){
	*xPtr= 0
}

func main() {
	x := 5
	inicial(&x)
	fmt.Println(x)
	
}