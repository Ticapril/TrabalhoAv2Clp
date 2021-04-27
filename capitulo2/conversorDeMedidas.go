package main // criação do pacote main

import  (
		"fmt" //pacote de tratamento de string
		"os" // 
		"strconv" // pacote de conversão de string
) 

func main(){
	
	//len(os.Args) é o tamanho do array de argumentos tipo o usuário digita 3 valores 
	if len(os.Args) < 3 {
		fmt.Printf("Uso: conversor <valores> <unidade>")
		os.Exit(1)
	}

	unidadeOrigem := os.Args[len(os.Args)-1] // eu pego o ultimo valor do array  ele aqui pega celsius e armazena na variavel 
	valoresOrigem := os.Args[1 : len(os.Args) - 1] // pega o valor origem que é 32

	var unidadeDestino string

	if unidadeOrigem == "celsius" {
		unidadeDestino = "fahrenheit"
	} else if unidadeOrigem == "quilometros"{
		unidadeDestino = "milhas"
	} else {
		fmt.Printf("%s não é uma unidade conhecida!",unidadeDestino)
		os.Exit(1)
	}

	//olha que lindo isso aqui

	for i, v := range valoresOrigem { //percorre cada valor do array
		valoresOrigem, err := strconv.ParseFloat(v,64)
		if err!= nil {
			fmt.Printf("O valor 5s na posição %d não é um número válido!\n",v, i)
			os.Exit(1)
		}

		var valorDestino float64

		if unidadeOrigem == "celsius" {
			valorDestino = valoresOrigem * 1.8 + 32
		} else {
			valorDestino = valoresOrigem / 1.60934
		}
		fmt.Printf("%.2f %s = %.2f %s\n", valoresOrigem, unidadeOrigem,valorDestino,unidadeDestino)
	}
}