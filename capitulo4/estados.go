package main
import "fmt"

func main(){
	type Aluno struct {
		nome string
		matricula int
		data_nascimento string		
		media_escolar float64
	}
	type Disciplina struct {
		nome string
		codigo_disciplina int		
		carga_horaria float64
		media_da_disciplina float64 // soma de todas as notas dos alunos na disciplina  e uma média
	}
	alunos := make(map[int]Aluno,5)
	alunos[0] = Aluno{"Gabriel",1723334045,"14-10-1997",5.90}
	alunos[1] = Aluno{"Marcos",1813334045,"10-12-1996",7.90}
	alunos[2] = Aluno{"Fernando",1913334045,"09-03-1990",8.90}

	alunos[3] = Aluno {}
	// var i = 0

	// for i <= 3 {
	// 	fmt.Println(alunos[i])
	// 	i++
	// }
	fmt.Println(alunos)
}