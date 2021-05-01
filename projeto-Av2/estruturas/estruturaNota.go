package estruturas

type Nota struct {
	Av1   float64
	Av2   float64
	Av3   float64
	Media float64
}

func (nota *Nota) AtribuirNota(av1, av2, av3 float64) {
	//atribuir as notas do struct e calcula a media
}

func (nota *Nota) CalcularMedia() float64 {
	return 0
	//Calcular a média considerando as três notas
}
