package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pabalberione/godesde0/ejercicios"
)

var filename string = "./files/txt/tabla.txt"

func GrabaTabla() { //Esta funcion va a pisar el archivo si ya existe
	var texto string = ejercicios.TablaMultiplicar()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return //Esto no muestra nada, solo sale de la funcion
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close() //Siempre tiene que cerrarse el buffer
}

func SumaTabla() { //Esta funcion va a concatenar lo que haya creado en el archivo
	var texto string = ejercicios.TablaMultiplicar()
	if !Append(filename, texto) { //Si NO ocurrió el Append, se está negando al principio de la funcion
		fmt.Println("Error al concatenar contenido")
	}

}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644) //Permite abrir el archivo para concatenarle contenido
	if err != nil {
		fmt.Println("Error durante el Append" + err.Error())
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString" + err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error al leer el archivo" + err.Error())
	}
	fmt.Println(string(archivo)) //Hay que convertir al archivo en string porque es un flujo de bytes
}

func LeoArchivoDos() {
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error al leer el archivo" + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo) //En vez de ingresar desde teclado ingresar al buffer desde el archivo
	for scanner.Scan() {                 //Devuelve un booleano si anduvo bien o mal
		registro := scanner.Text() //Devuelve cada uno de las lineas del archivo
		fmt.Println("> " + registro)
	}
	archivo.Close() //Siempre cerrar la apertura del archivo
}
