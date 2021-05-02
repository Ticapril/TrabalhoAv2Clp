package main

import (
	e "estruturas"
	"fmt"
)

var turmas []e.Turma

func main() {

	fmt.Println("Calculador de Estatísticas de Notas Sobre uma Turma!")

	var opcao = 1

	for opcao != 0 {

		fmt.Printf("\n")
		fmt.Println("Digite uma opção:")
		fmt.Println("0 - Sair")
		fmt.Println("1 - Cadastrar Turma")
		fmt.Println("2 - Ver estatísticas")
		fmt.Print("Opção: ")
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			cadastrarTurma()
		case 2:
			verEstatisticas()
		}

	}
	fmt.Println("Obrigado por utilizar o programa!")
}

func cadastrarTurma() {
	//declarações de variavel
	var turma e.Turma
	var numeroTurma, quantidadeAlunos int
	var nome, matricula string
	var av1, av2, av3 float64

	fmt.Printf("\n")
	fmt.Printf("Entre com o número da turma: ")
	fmt.Scan(&numeroTurma)
	turma.DefinirNumero(numeroTurma) //atribuição ao numero da turma

	fmt.Printf("Entre com a quantidade de alunos a serem cadastrados: ")
	fmt.Scan(&quantidadeAlunos)

	for i := 0; i < quantidadeAlunos; i++ { //atribuindo alunos na turma
		var aluno e.Aluno
		fmt.Printf("Aluno %d \n", (i + 1))
		fmt.Print("Entre com a matricula do aluno: ")
		fmt.Scan(&matricula)
		fmt.Print("Entre com o nome do aluno: ")
		fmt.Scan(&nome)
		fmt.Print("Entre com a nota da av1: ")
		fmt.Scan(&av1)
		fmt.Print("Entre com a nota da av2: ")
		fmt.Scan(&av2)
		fmt.Print("Entre com a nota da av3: ")
		fmt.Scan(&av3)
		fmt.Printf("\n")
		aluno.AtribuirDados(nome, matricula, av1, av2, av3)
		turma.DefinirAlunos(append(turma.ObterAlunos(), aluno)) // eu acesso o slice na struct Turma e armazeo o aluno la
	}
	turmas = append(turmas, turma) // aqui eu só estou adicionando a nova turma ao slice de turmas global;
	fmt.Println("Cadastro realizado com sucesso!")
}

func verEstatisticas() {

	if len(turmas) > 0 {

		var opcao int
		fmt.Printf("\n")
		fmt.Println("Informe o Número da Turma")
		fmt.Println("Opções: ")
		for _, turma := range turmas {
			fmt.Printf("%d, ", turma.ObterNumero())
		}
		fmt.Printf("\n")
		fmt.Print("Turma: ")
		fmt.Scan(&opcao)

		var encontrado = false

		for _, turma := range turmas {
			if turma.ObterNumero() == opcao {
				turma.ExibirRelatario()
				encontrado = true
			}
		}

		if !encontrado {
			fmt.Println("Turma não encontrada para o número informado!")
		}

	} else {
		fmt.Println("Nenhuma Turma Cadastrada!")
	}

}
