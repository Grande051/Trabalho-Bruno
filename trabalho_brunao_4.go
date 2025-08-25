package main

import (
	"fmt"
	"os"
)

func exibirMenu() {
	fmt.Println("\n============================")
	fmt.Println("       MENU PRINCIPAL       ")
	fmt.Println("============================")
	fmt.Println("1. Opção A")
	fmt.Println("2. Opção B")
	fmt.Println("3. Opção C")
	fmt.Println("0. Sair")
	fmt.Println("============================")
}

func main() {
	fmt.Println("### Exibindo Menu ###")

	for {
		exibirMenu()

		var escolha int
		fmt.Print("Escolha uma opção: ")
		_, err := fmt.Scanln(&escolha)

		if err != nil {
			fmt.Println("Entrada inválida. Por favor, digite um número.")
			var temp string
			fmt.Scanln(&temp)
			continue
		}

		if escolha == 0 {
			break
		}

		switch escolha {
		case 1:
			fmt.Println("Você escolheu a Opção A.")
		case 2:
			fmt.Println("Você escolheu a Opção B.")
		case 3:
			fmt.Println("Você escolheu a Opção C.")
		default:
			fmt.Println("Opção inválida. Por favor, tente novamente.")
		}
	}

	fmt.Println("\nPrograma encerrado. Até mais!")
	os.Exit(0)
}
