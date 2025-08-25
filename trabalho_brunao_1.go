package main

import (
	"fmt"
	"os"
)

func calculadora(a, b float64, operador string) (float64, bool) {

	if operador != "+" && operador != "-" && operador != "*" && operador != "/" {
		fmt.Println("Erro: Operador inválido. Use +, -, *, ou /.")
		return 0, false
	}
	if operador == "/" && b == 0 {
		fmt.Println("Erro: Divisão por zero não é permitida.")
		return 0, false
	}
	var resultado float64
	switch operador {
	case "+":
		resultado = a + b
	case "-":
		resultado = a - b
	case "*":
		resultado = a * b
	case "/":
		resultado = a / b
	}
	return resultado, true
}
func main() {
	fmt.Println("### Calculadora ###")
	fmt.Println("Digite 'sair' para fechar o programa a qualquer momento.")
	for {
		var num1, num2 float64
		var operador string
		var entrada string

		fmt.Print("\nDigite o primeiro número (ou 'sair'): ")
		fmt.Scanln(&entrada)
		if entrada == "sair" {
			break
		}
		_, err := fmt.Sscanf(entrada, "%f", &num1)
		if err != nil {
			fmt.Println("Entrada inválida. Por favor, digite um número.")
			continue
		}
		fmt.Print("Digite o operador (+, -, *, /): ")
		fmt.Scanln(&operador)
		if operador == "sair" {
			break
		}
		fmt.Print("Digite o segundo número (ou 'sair'): ")
		fmt.Scanln(&entrada)
		if entrada == "sair" {
			break
		}
		_, err = fmt.Sscanf(entrada, "%f", &num2)
		if err != nil {
			fmt.Println("Entrada inválida. Por favor, digite um número.")
			continue
		}
		resultado, valido := calculadora(num1, num2, operador)
		if valido {
			fmt.Printf("Resultado: %.2f\n", resultado)
		}
	}
	fmt.Println("\nPrograma encerrado. Até mais!")
	os.Exit(0)
}
