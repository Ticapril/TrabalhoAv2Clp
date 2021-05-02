package main

import (
	e "TrabalhoAv2Clp/projeto-Av2/estruturas"
	"fmt"
)

var turmas []e.Turma

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
	var numeroTurma uint8
	var nome, matricula string
	var av1, av2, av3 float64
	fmt.printf("Entre com o número da turma -->")
	scanf("%d", &numeroTurma)
	turma.numero = numeroTurma //atribuição ao numero da turma

	for i, elemento := range turma.alunos { //atribuindo alunos na turma
		fmt.Println("Entre com o nome do aluno")
		fmt.Scanf("%s", nome)
		fmt.Println("Entre com a matricula do aluno")
		scanf("%s", matricula)
		println("Entre com a nota da av1")
		scanf("%s", &av1)
		println("Entre com a nota da av2")
		scanf("%s", &av2)
		println("Entre com a nota da av3")
		scanf("%s", &av3)
		turma.alunos[i].atribuirDados(nome, matricula, av1, av2, av3)
	}
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
