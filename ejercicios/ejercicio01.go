package ejercicios

import (
	"strconv"
)

func COnvNumerico(valor string) (int, string) {
	num, err := strconv.Atoi(valor)
	if err != nil {
		return 0, "Hubo un error " + err.Error()
	}
	if num > 100 {
		return num, "Es mayor a 100"
	} else {
		return num, "Es menor a 100"
	}

}
