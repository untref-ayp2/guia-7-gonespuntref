package ejercicios

import "fmt"

// Escribir un método recursivo que dado un número
// entero decimal devuelva su equivalente en
// binario en forma de string
func DecimalABinario(n int) string {
	if n < 2 {
		return fmt.Sprint(n % 2)
	}
	return DecimalABinario(n/2) + fmt.Sprint(n%2)
}
