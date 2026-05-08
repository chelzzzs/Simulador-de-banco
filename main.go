package main

import (
	"fmt"
	c "projects/banco/contas"
)

func main() {
    contaDaSilvia := c.ContaCorrente{Titular:"Silvia", Saldo: 300}
    contaDoGustavo := c.ContaCorrente{Titular:"Gustavo", Saldo: 100}
    
    fmt.Println(contaDaSilvia)
    fmt.Println(contaDoGustavo)

	
		status := contaDoGustavo.Transferir(-200,  &contaDaSilvia)

		fmt.Println(status)
		fmt.Println(contaDaSilvia)
		fmt.Println(contaDoGustavo)
}