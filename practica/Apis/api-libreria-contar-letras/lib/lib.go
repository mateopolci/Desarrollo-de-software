package lib

func ContarLetras(palabra string) int {
	contador := 0

	for _, letra := range palabra {
		if (letra >= 'a' && letra <= 'z') || (letra >= 'A' && letra <= 'Z') {
			contador++
		}
	}

	return contador
}