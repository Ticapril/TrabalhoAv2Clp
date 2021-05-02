package estruturas

import "fmt"

type Turma struct {
	numero int
	alunos []Aluno
}

func (turma *Turma) DefinirNumero(numero int) {
	turma.numero = numero
}

func (turma *Turma) ObterNumero() int {
	return turma.numero
}

func (turma *Turma) ObterAlunos() []Aluno {
	return turma.alunos
}

func (turma *Turma) DefinirAlunos(alunos []Aluno) {
	turma.alunos = alunos
}

func (turma *Turma) AdicionarAluno(aluno Aluno) {
	turma.alunos = append(turma.alunos, aluno)
}

func (turma *Turma) QuantidadeDeAlunos() int {
	return len(turma.alunos)
}

func (turma *Turma) CalcularMediaDaTurma() float64 {
	var somatorio = 0.00

	for _, aluno := range turma.alunos {
		somatorio += aluno.nota.media
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

		if n.nota.media <= pivo.nota.media {
			menores = append(menores, n)
		} else {
			maiores = append(maiores, n)
		}

	}

	return append(append(ordenarNotas(menores), pivo), ordenarNotas(maiores)...)
}

func (turma *Turma) ExibirRelatario() {
	fmt.Printf("Relatório da Turma: %d\n", turma.numero)
	fmt.Printf("Quantidade de Aluno: %d\n", turma.QuantidadeDeAlunos())
	fmt.Printf("Média Geral da Turma: %.2f\n", turma.CalcularMediaDaTurma())

	fmt.Println("Dados dos Alunos:")
	fmt.Println("------------------------------------------------------------------------")
	fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\n", "Matricula", "Nome", "AV1", "AV2", "AV3", "Média", "Situação")

	var alunosOrdenados = ordenarNotas(turma.alunos)

	for _, aluno := range alunosOrdenados {

		fmt.Printf("%s\t%s\t%.2f\t%2.f\t%.2f\t%.2f\t%s\n", aluno.matricula, aluno.nome,
			aluno.nota.av1, aluno.nota.av2, aluno.nota.av3, aluno.nota.media, aluno.verificarSituacao())

	}

}
