package estruturas

import "fmt"

type Turma struct {
	Numero int
	Alunos []Aluno
}

func (turma *Turma) AdicionarAluno(aluno Aluno) {
	//atribuir aluno ao slide aluno da turma
	turma.alunos = append(alunos, aluno)
}

func (turma *Turma) quantidadeDeAlunos() int {
	//calcular a quantidade de alunos na turma
	return int(len(turma.alunos))
}

func (turma *Turma) CalcularMediaGeral() float64 {
	//calcular a média geral. calculo: soma de todas as médias dividido pela quantidade de alunos
	var somatorio = 0.00

	for _, aluno := range turma.Alunos {
		somatorio += aluno.CalcularMedia()
	}

	var media = somatorio / float64(turma.QuantidadeDeAlunos())

	return media
}

func ordenarNotas(alunos []Aluno) []Aluno {
	// ordernar da maior pra a menor média as notas dos alunos
	if len(alunos) <= 1 {
		return alunos
	}

	n := make([]Aluno, len(alunos))
	copy(n, alunos)

	indicePivo := len(n) / 2
	pivo := n[indicePivo]

	n = append(n[:indicePivo], n[indicePivo+1:]...)

	var menores []Aluno
	var maiores []Aluno

	for _, n := range alunos {

		if n.CalcularMedia() <= pivo.CalcularMedia() {
			menores = append(menores, n)
		} else {
			maiores = append(maiores, n)
		}

	}

	return append(append(ordenarNotas(menores), pivo), ordenarNotas(maiores)...)
}

func (turma *Turma) ExibirRelatario() {
	//Exibir um relatório contendo:
	//Número da Turma
	//Quantidade de alunos
	//Média da Turma
	//Dados dos Alunos:
	// matricula -> nome -> av1 -> av2 -> av3 -> media -> aprovado ou reprovado
	fmt.Printf("Relatório da Turma: %d\n", turma.Numero)
	fmt.Printf("Quantidade de Aluno: %d\n", turma.QuantidadeDeAlunos())
	fmt.Printf("Média Geral da Turma: %.2f\n", turma.CalcularMediaGeral())
	fmt.Println("Dados dos Alunos:")
	fmt.Println("------------------------------------------------------------------------")
	fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\n", "Matricula", "Nome", "AV1", "AV2", "AV3", "Média", "Situação")

	var alunosOrdenados = ordenarNotas(turma.Alunos)

	for _, aluno := range alunosOrdenados {

		fmt.Printf("%s\t%s\t%.2f\t%2.f\t%.2f\t%.2f\t%s\n", aluno.Matricula, aluno.Nome,
			aluno.Nota.Av1, aluno.Nota.Av2, aluno.Nota.Av3, aluno.CalcularMedia(), aluno.VerificarSituacao())

		println("----------------------------------------------------------------------------")

	}

	println("----------------------------------------------------------------------------")

}
