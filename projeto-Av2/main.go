package main

import (
	"fmt"
)

func main() {

	fmt.Println("Calculador de Estatísticas de Notas Sobre uma Turma!")

	var opcao = 1
	var turmas []Turma // lista de turmas

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

func (turmas []Turma) cadastrarTurma() {
	//declarações de variavel
	var turma Turma
	var numeroTurma, quantidadeAlunos int
	var nome, matricula string
	var av1, av2, av3 float64

	fmt.printf("Entre com o número da turma -->")
	fmt.scanf("%d", &numeroTurma)
	turma.numero = numeroTurma //atribuição ao numero da turma

	fmt.printf("Entre com a quantidade de alunos a serem cadastrados -->")
	fmt.scanf("%d", &quantidadeAlunos)

	for i := 0; i < quantidadeAlunos; i++ { //atribuindo alunos na turma
		var aluno Aluno
		fmt.println("Entre com o nome do aluno")
		fmt.scanf("%s", &nome)
		fmt.println("Entre com a matricula do aluno")
		fmt.scanf("%s", &matricula)
		fmt.println("Entre com a nota da av1")
		fmt.scanf("%s", &av1)
		fmt.println("Entre com a nota da av2")
		fmt.scanf("%s", &av2)
		fmt.println("Entre com a nota da av3")
		fmt.scanf("%s", &av3)
		turma.alunos[i].atribuirDados(nome, matricula, av1, av2, av3)
		aluno.AtribuirDados(nome, matricula, av1, av2, av3)
		turma.alunos = append(turma.alunos, aluno) // eu acesso o slice na struct Turma e armazeo o aluno la
	}
	turmas = append(turmas, turma) // aqui eu só estou adicionando direto no slice de turmas e nova turma criada
	fmt.Println("Cadastro de turma realizado com sucesso!")
}

func verEstatisticas() {
	fmt.Println("Ver estatísticas.....")
}
