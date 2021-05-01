package main

import (
	"fmt"
)

func main(){
	sliceAlunos := make([]int, 10)
	fmt.Println(sliceAlunos, len(sliceAlunos), cap(sliceAlunos))

	for i := range sliceAlunos {
		sliceAlunos[i] += 25;
	}
	fmt.Println(sliceAlunos)
	
}