package estruturas

type Aluno struct {
	matricula string
	nome      string
	nota      Nota
}

func (aluno *Aluno) atribuirDados(nome, matricula string, av1, av2, av3 float64) {
	//atribuir os valores a aluno e as notas dele utilizando o método atribuirNota
	aluno.nome = nome
	aluno.matricula = matricula
	aluno.nota.AtribuirNota(av1, av2, av3)
}

func (aluno *Aluno) CalcularMedia() float64 {
	return aluno.nota.calcularMedia()
}

func (aluno *Aluno) verificarSituacao() string {
	if aluno.nota.media >= 6.0 {
		return "Aprovado"
	} else {
		return "Reprovado"
	}
	//verifica se o aluno está aprovado ou reprovado
}
