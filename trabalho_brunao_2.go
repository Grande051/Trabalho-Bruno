package main

import (
	"fmt"
	"os"
)

const (
	idiomaPortugues = "pt"
	idiomaIngles    = "en"
	idiomaEspanhol  = "es"
	saudacaoPadrao  = "Idioma não suportado."
)

func saudacao(nome, idioma string) {
	if nome == "" {
		nome = "visitante"
	}

	prefixo := saudacaoPadrao
	switch idioma {
	case idiomaPortugues:
		prefixo = "Olá, "
	case idiomaIngles:
		prefixo = "Hello, "
	case idiomaEspanhol:
		prefixo = "Hola, "
	}

	fmt.Println(prefixo + nome + "!")
}

func main() {
	fmt.Println("### Saudações Multi-idioma ###")
	fmt.Println("Digite 'sair' para fechar o programa a qualquer momento.")

	for {
		var nome, idioma, entrada string

		fmt.Print("\nDigite seu nome (ou 'sair' para fechar): ")
		fmt.Scanln(&entrada)
		if entrada == "sair" {
			break
		}
		nome = entrada

		fmt.Print("Digite o código do idioma ('pt', 'en', 'es'): ")
		fmt.Scanln(&entrada)
		if entrada == "sair" {
			break
		}
		idioma = entrada
		saudacao(nome, idioma)
	}

	fmt.Println("\nPrograma encerrado. Até mais!")
	os.Exit(0)
}
