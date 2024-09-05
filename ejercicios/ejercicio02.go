package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var texto string

func TablaMultiplicar() string { //Modificamos la funcion para indicar que va a devolver un tipo string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese un número")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}
	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, i*numero) // A la variable texto se le asigna al final el valor del Sprintf
	}

	return texto
}
