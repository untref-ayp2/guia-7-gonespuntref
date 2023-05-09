package ejercicios

// Escribir un método recursivo que devuelva el
// cociente y el resto de la división entera mediante
// restas sucesivas
func DivisionEntera(dividendo, divisor int) (cociente, resto int) {

	if dividendo == 0 {
		return 0, 0
	}
	if divisor == 0 {
		return -1, 0
	}
	if dividendo < divisor {
		return 1, dividendo - divisor
	}
	return DivisionEntera(dividendo-divisor, divisor)
}

/*
if dividendo == 0 {
	return 0, 0
}
if divisor == 0 {
	return -1, 0
}
return 1 + DivisionEntera(dividendo-divisor, divisor)

func Division(a, b int) int {

	if a == 0 || a < b {
		return 0
	}
	if b == 0 {
		return -1
	}
	return 1 + Division(a-b, b)
}
*/
