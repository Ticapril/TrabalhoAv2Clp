package estruturas

type Aluno struct {
	Matricula string
	Nome      string
	Nota      Nota
}

func (aluno *Aluno) atribuirDados(nome, matricula string, av1, av2, av3 float64) {
	//atribuir os valores a aluno e as notas dele utilizando o método atribuirNota
	aluno.nome = nome
	aluno.matricula = matricula
	aluno.nota = nota.atribuirNota(av1, av2, av3)
}

func (aluno *Aluno) CalcularMedia() float64 {
	return 0
	//obter média a partir do método calcularMedia da estrutura de nota
}

func (aluno *Aluno) verificarSituacao() string {
	if aluno.nota.media >= 6.0 {
		return "Aprovado"
	} else {
		return "Reprovado"
	}
	//verifica se o aluno está aprovado ou reprovado
}
