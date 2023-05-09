package ejercicios

// Escriba un método recursivo que calcule el máximo
// común divisor entre dos números enteros.
// Nota: Se puede usar el algoritmo de Euclides para
// resolver este problema.
func MCD(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	return MCD(b, a%b)

}
