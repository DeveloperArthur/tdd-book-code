package stocks

import (
	"strconv"
)

func ExecuteCompressionLossLess(input string) string {
	var result string = ""

	/*Variável que conta quantos caracteres iguais existem
	o primeiro caracter já deve ser contato
	por isso count deve ser inicializado como 1
	se não, caso a entrada fosse: AAA, o resultado seria A2*/
	var count int = 1

	for i := 0; i < len(input); i++ {
		char := string(input[i])

		//para evitar erro de index out
		var ultimaIteracao bool = (i == (len(input) - 1))
		if ultimaIteracao {
			result = result + concat(char, count)
			break
		}

		nextChar := string(input[i+1])

		if char == nextChar {
			count = count + 1
		} else {
			result = result + concat(char, count)
			count = 1
		}
	}

	return result
}

func concat(char string, count int) string {
	return (char + strconv.Itoa(count))
}
