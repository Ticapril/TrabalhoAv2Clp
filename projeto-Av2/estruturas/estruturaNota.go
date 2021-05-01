package estruturas

type Nota struct {
	av1   float64
	av2   float64
	av3   float64
	media float64
}

func (nota *Nota) atribuirNota(av1, av2, av3 float64) {
	//atribuir as notas do struct e calcula a media
}

func (nota *Nota) calcularMedia() float64 {
	return 0
	//Calcular a média considerando as três notas
}
