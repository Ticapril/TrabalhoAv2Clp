package estruturas

type Nota struct {
	Av1   float64
	Av2   float64
	Av3   float64
	Media float64
}

func (nota *Nota) AtribuirNota(av1, av2, av3 float64) {
	//atribuir as notas do struct e calcula a media
	nota.av1 = av1
	nota.av2 = av2
	nota.av3 = av3
	nota.media = (av1 + av2 + av3)/3  
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
