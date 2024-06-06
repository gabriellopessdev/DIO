package main
import "fmt"

// func torrar(pao string, nome ... string){
//     cliente := "Não teve cliente indentificado"//variavel para trtar paramentro opcinal representado por "nome ..."
//     switch {
// 	case len(nome) > 0: // se nome tiver um parametro indentificado "cliente = nome[0]"
//         	cliente = nome[0]
//     		fmt.Println("Torrada feita com", pao)
//     		fmt.Println("é um pedido de", cliente)
// 	default:
// 		fmt.Println("Torrada feita com", pao)
// 		fmt.Println(cliente)
// 	}
// }

func torrar(pao string, nome ...string) {
    cliente := "Não teve cliente indentificado"

    if len(nome) > 0 {
        cliente = nome[0]
        fmt.Println("Torrada feita com", pao)
        fmt.Println("é um pedido de", cliente)
    } else {
        fmt.Println("Torrada feita com", pao)
        fmt.Println(cliente)
    }
}

func main() {
    torrar("pão de forma", "Gabriel")
    torrar("pão de sal")

}
