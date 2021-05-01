package main
import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano())
	n := 0

	for {
		n++

		i := rand.Intn(4200)
		fmt.Println(i)

		if i%42 == 0 {
		break;

		}
	}

	fmt.Printl("Saída após %d iterações.\n",n)
//for aninhado

	var i int
	externo:
	for {
		for i = 0;  i < 10; i++ {
				if i == 5 {
					break externo
				}
		}
	}
}