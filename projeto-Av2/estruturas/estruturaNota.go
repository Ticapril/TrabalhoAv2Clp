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
	if nota.av3 == 0 {
		return (nota.av1 + nota.av2) / 2
	} else {
		if nota.av3 > nota.av1 {
			return (nota.av3 + nota.av2) / 2
		} else if nota.av3 > nota.av2 {
			return (nota.av3 + nota.av1) / 2
		} else {
			return (nota.av1 + nota.av2) / 2
		}
	}
	//Calcular a média considerando as três notas
}
