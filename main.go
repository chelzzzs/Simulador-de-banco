package main

import (
	"banco/contas"
	"fmt"
)



func pagarBoleto (conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string 
}


func main() {

contaDoMichel := contas.ContaPoupanca{}
contaDoMichel.Depositar(100)
pagarBoleto(&contaDoMichel, 60)	
fmt.Println(contaDoMichel)


contaDaLuisa := contas.ContaCorrente{}
contaDaLuisa.Depositar(100)

pagarBoleto(&contaDaLuisa, 50)

fmt.Println(contaDaLuisa)
}