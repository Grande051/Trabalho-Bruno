package main

import (
	"fmt"
	"os"
)

func ehPrimo(numero int) bool {
	if numero <= 1 {
		return false
	}

	if numero == 2 {
		return true
	}

	if numero%2 == 0 {
		return false
	}

	for i := 3; i*i <= numero; i += 2 {
		if numero%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("### Verificador de Números Primos ###")
	fmt.Println("Digite 'sair' para fechar o programa a qualquer momento.")

	for {
		var numero int
		var entrada string

		fmt.Print("\nDigite um número inteiro (ou 'sair'): ")
		fmt.Scanln(&entrada)

		if entrada == "sair" {
			break
		}

		_, err := fmt.Sscanf(entrada, "%d", &numero)
		if err != nil {
			fmt.Println("Entrada inválida. Por favor, digite um número inteiro.")
			continue
		}

		if ehPrimo(numero) {
			fmt.Printf("%d é primo.\n", numero)
		} else {
			fmt.Printf("%d não é primo.\n", numero)
		}
	}

	fmt.Println("\nPrograma encerrado. Até mais!")
	os.Exit(0)
}
