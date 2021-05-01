package main

import (
	"fmt"
)

func main() {

	fmt.Println("Calculador de Estatísticas de Notas Sobre uma Turma!")

	var opcao = 1

	for opcao != 0 {
		fmt.Println("Digite uma opção:")
		fmt.Println("0 - Sair")
		fmt.Println("1 - Cadastrar Turma")
		fmt.Println("2 - Ver estatísticas")
		fmt.Print("Opção: ")
		fmt.Scanln(&opcao)
		switch opcao {
		case 1:
			cadastrarTurma()
			break
		case 2:
			verEstatisticas()
			break
		}

	}

	fmt.Println("Obrigado por utilizar o programa!")

}

func cadastrarTurma() {
	fmt.Println("Cadastro de turma......")
}

func verEstatisticas() {
	fmt.Println("Ver estatísticas.....")
}
