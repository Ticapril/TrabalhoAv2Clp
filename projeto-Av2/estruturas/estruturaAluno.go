package estruturas

type Aluno struct {
	nome      string
	matricula string
	nota      Nota
}

func (aluno *Aluno) AtribuirDados(nome, matricula string, av1, av2, av3 float64) {
	//atribuir os valores a aluno e as notas dele utilizando o método atribuirNota
	aluno.nome = nome
	aluno.matricula = matricula
	aluno.nota.atribuirNota(av1, av2, av3)
}

func (aluno *Aluno) calcularMedia() float64 {
	return 0
	//obter média a partir do método calcularMedia da estrutura de nota
}

func (aluno *Aluno) verificarSituacao() string {
	return ""
	//verifica se o aluno está aprovado ou reprovado
}
