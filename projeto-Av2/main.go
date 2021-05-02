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

func (turma *Turma) cadastrarTurma() {
	var numeroTurma uint8
	var nome,matricula string
	var av1,av2,av3 float64
	fmt.printf("Entre com o número da turma -->")
	scanf("%d",&numeroTurma)
	turma.numero = numeroTurma //atribuição ao numero da turma 
	
	for i, elemento := range turma.alunos { //atribuindo alunos na turma
		println("Entre com o nome do aluno")
		scanf("%s",nome)
		println("Entre com a matricula do aluno")
		scanf("%s",matricula)
		println("Entre com a nota da av1")
		scanf("%s",&av1)
		println("Entre com a nota da av2")
		scanf("%s",&av2)
		println("Entre com a nota da av3")
		scanf("%s",&av3)
		turma.alunos[i].atribuirDados(nome,matricula,av1,av2,av3) 
	}
	fmt.Println("Cadastro de turma......")
}

func verEstatisticas() {
	fmt.Println("Ver estatísticas.....")
}
