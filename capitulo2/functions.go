package main

import "fmt"

func showData(nome string, idade int){
	fmt.Printf("Usuário: %s\nIdade: %d", nome,idade)
}
func sumAge(idade1,idade2 int) int{
	var ageSum int = idade1 + idade2
	return ageSum
}

func main(){
	showData("Gabriel", 23)
	var sumAge = sumAge(23, 44)
	fmt.Printf("a soma das idades passadas foi de = %d", sumAge)
}