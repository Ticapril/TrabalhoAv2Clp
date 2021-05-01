package estruturas

type Turma struct {
	numero int
	alunos []Aluno
}

func (turma *Turma) adicionarAluno(aluno Aluno) {
	//atribuir aluno ao slide aluno da turma
}

func (turma *Turma) quantidadeDeAlunos() float64 {
	//calcular a quantidade de alunos na turma
	return 0
}

func (turma *Turma) calcularMediaGeral() float64 {
	return 0
	//calcular a média geral. calculo: soma de todas as médias dividido pela quantidade de alunos
}

func (turma *Turma) ordenarNotas() []Aluno {
	return nil // ordernar da maior pra a menor média as notas dos alunos
}

func (turma *Turma) exibirRelatario() {
	//Exibir um relatório contendo:
	//Número da Turma
	//Quantidade de alunos
	//Média da Turma
	//Dados dos Alunos:
	// matricula -> nome -> av1 -> av2 -> av3 -> media -> aprovado ou reprovado
}
