package main

import (
	e "TrabalhoAv2Clp/projeto-Av2/estruturas"
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
			verEstatisticas(nil)
			break
		}

	}

	fmt.Println("Obrigado por utilizar o programa!")

}

func cadastrarTurma() {
	fmt.Println("Cadastro de turma......")
}

func verEstatisticas(turmas []e.Turma) {
	fmt.Println("Informe o número da turma")
	fmt.Println("Opções: ")
	for _, turma := range turmas {
		fmt.Printf("%d, ", turma.Numero)
	}
	fmt.Printf("\n")
	var opcao int
	fmt.Print("Turma: ")
	fmt.Scanln(&opcao)

	var t e.Turma
	t.Numero = 0

	for _, turma := range turmas {

		if turma.Numero == opcao {
			t = turma
		}

	}

	if t.Numero != 0 {
		t.ExibirRelatario()

	} else {
		fmt.Println("Turma não encontrada para o número informado!")

	}

}
