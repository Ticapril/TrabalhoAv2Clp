package estruturas

type Nota struct {
	av1   float64
	av2   float64
	av3   float64
	media float64
}

func (nota *Nota) atribuirNota(av1, av2, av3 float64) {
	//atribuir as notas do struct e calcula a media
	nota.av1 = av1
	nota.av2 = av2
	nota.av3 = av3
	nota.media = (av1 + av2 + av3)/3  
}

func (nota *Nota) calcularMedia() float64 {
	return 0
	//Calcular a média considerando as três notas
}
