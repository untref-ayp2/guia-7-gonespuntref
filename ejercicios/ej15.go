package ejercicios

// Se tiene un arreglo de n >= 3 elementos en forma de pico,
// esto es: estrictamente creciente hasta una determinada posición p,
// y estrictamente decreciente a partir de ella (con 0 < p < n-1).
// Por  ejemplo, en el arreglo [1,2,3,1,0,-2] la posición del pico es p=2.
// Se pide implementar un algoritmo de división y conquista de orden
// O(log n) que encuentre la posición p del pico.
func Pico(arreglo []int) int {
	return pico(arreglo, 0, len(arreglo)-1)
}

func pico(array []int, inicio int, fin int) int {
	if inicio == fin {
		return inicio
	}

	medio := (inicio + fin) / 2

	if array[medio] < array[medio+1] {
		return pico(array, medio+1, fin)
	} else {
		return pico(array, inicio, medio)
	}

}
