package main

import  (
	"fmt" //pacote de tratamento de string
	"os" 
	"strconv" // pacote de conversões de tipo 
) 

func main(){
	entrada := os.Args[1:] //faz ele pegar o range de valores passados 
	numeros := make([]int, len(entrada)) // faça um array de inteiros com os valores da entrada declarei ta vazio

	for i, v := range entrada { // range me devolve um indice e o valor desse indice
		numero, err := strconv.Atoi(v) // strconv me retorna dois valores o valor em si e o err se o err for igual a nil deu certo a conversão
		// so executa o if se der algum erro, ou seja, ele passar algum valor diferente de um inteiro
		if err != nil {
			fmt.Printf("%s não é um número válido!\n", v)
			os.Exit(1)
		}
		numeros[i] = numero // atribui o valor na posição 
	}
	
	fmt.Println(quicksort(numeros))
}
	//verifica se esta vazia ou se possui apenas um numero
	func quicksort(numeros [] int) [] int {
		if len(numeros) <= 1 {
			return numeros
		}
		//faço uma copia de numeros
	
		n := make([]int, len(numeros))
		copy(n,numeros)
	
		//pivo definido foi o do meio se for par 
		indicePivo := len(n)/2
		pivo := n[indicePivo]
	
		//forma geral append novoSlice := append(slice, elemento)
		n = append(n[:indicePivo], n[indicePivo+1:]...) 
		//eu crio duas listas listaAtePivo e listaDepoisPivo e junto as duas listas que seria uma lista sem o pivo
		menores,maiores := particionar(n, pivo)

		return append(append(quicksort(menores),pivo),quicksort(maiores)...)
	}
	func particionar(numeros [] int, pivo int) (menores [] int, maiores[] int){
		for _,n := range numeros{
			if n <= pivo {
				menores = append(menores, n)
			} else {
				maiores = append(maiores, n)
			}
		}
		return menores, maiores
	}
