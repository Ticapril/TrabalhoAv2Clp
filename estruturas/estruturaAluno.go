package estruturas

type Aluno struct {
	matricula string
	nome      string
	nota      Nota
}

func (aluno *Aluno) AtribuirDados(nome, matricula string, av1, av2, av3 float64) {
	aluno.nome = nome
	aluno.matricula = matricula
	aluno.nota.AtribuirNota(av1, av2, av3)
}

func (aluno *Aluno) verificarSituacao() string {
	if aluno.nota.media >= 6.0 {
		return "Aprovado"
	} else {
		return "Reprovado"
	}
}
