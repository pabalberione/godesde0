package iteraciones

import "fmt"

func Iterar() {

	for i := 0; i < 100; i += 5 {
		fmt.Println(i)

	}
}

func IterarHasta8() {

	for i := 0; i < 10; i++ {
		if i == 8 {
			break
		}
		fmt.Println(i)

	}
}

func IterarSin8() {
	for i := 0; i < 10; i++ {
		if i == 8 {
			continue
		}
		fmt.Println(i)

	}
}
