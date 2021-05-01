package estruturas

type Aluno struct {
	Matricula string
	Nome      string
	Nota      Nota
}

func (aluno *Aluno) AtribuirDados(nome, matricula string, av1, av2, av3 float64) {
	//atribuir os valores a aluno e as notos dele utilizando o método atribuirNota
}

func (aluno *Aluno) CalcularMedia() float64 {
	return 0
	//obter média a partir do método calcularMedia da estrutura de nota
}

func (aluno *Aluno) VerificarSituacao() string {
	return ""
	//verifica se o aluno está aprovado ou reprovado
}
