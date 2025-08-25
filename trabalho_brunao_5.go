package main

import (
	"fmt"
	"os"
)

func calcularSituacao(nota1, nota2, nota3 float64) string {
	media := (nota1 + nota2 + nota3) / 3

	if media >= 7.0 {
		return "Aprovado"
	} else if media >= 5.0 && media < 7.0 {
		return "Recuperação"
	} else {
		return "Reprovado"
	}
}

func main() {
	fmt.Println("### Calculadora de Média e Situação do Aluno ###")
	fmt.Println("Digite 'sair' a qualquer momento para fechar o programa.")

	for {
		var nota1, nota2, nota3 float64
		var entrada string

		fmt.Print("\nDigite a primeira nota (ou 'sair'): ")
		fmt.Scanln(&entrada)
		if entrada == "sair" {
			break
		}
		_, err := fmt.Sscanf(entrada, "%f", &nota1)
		if err != nil {
			fmt.Println("Entrada inválida. Por favor, digite um número.")
			continue
		}

		fmt.Print("Digite a segunda nota: ")
		fmt.Scanln(&entrada)
		if entrada == "sair" {
			break
		}
		_, err = fmt.Sscanf(entrada, "%f", &nota2)
		if err != nil {
			fmt.Println("Entrada inválida. Por favor, digite um número.")
			continue
		}

		fmt.Print("Digite a terceira nota: ")
		fmt.Scanln(&entrada)
		if entrada == "sair" {
			break
		}
		_, err = fmt.Sscanf(entrada, "%f", &nota3)
		if err != nil {
			fmt.Println("Entrada inválida. Por favor, digite um número.")
			continue
		}

		situacao := calcularSituacao(nota1, nota2, nota3)
		fmt.Printf("Média das notas: %.2f\n", (nota1+nota2+nota3)/3)
		fmt.Printf("Situação do aluno: %s\n", situacao)
	}

	fmt.Println("\nPrograma encerrado. Até mais!")
	os.Exit(0)
}
