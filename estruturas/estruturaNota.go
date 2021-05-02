package estruturas

type Nota struct {
	av1   float64
	av2   float64
	av3   float64
	media float64
}

func (nota *Nota) AtribuirNota(av1, av2, av3 float64) {
	nota.av1 = av1
	nota.av2 = av2
	nota.av3 = av3
	nota.media = nota.calcularMedia()
}

func (nota *Nota) calcularMedia() float64 {
	//Mudar a Lógica Aqui Rafael
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
}
