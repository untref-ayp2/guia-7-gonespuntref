package ejercicios

// Escribir un método recursivo que tome un
// array de números enteros y devuelva la suma
// de todos sus elementos
/*
func SumaArray(v []int) int {
	if len(v) < 1 {
		return 0
	}
	return v[0] + SumaArray(v[1:])
}
*/

func SumaArray(v []int) int {
	if len(v) < 1 {
		return 0
	}
	if len(v) == 1 {
		return v[0]
	}
	medioArray := len(v) / 2
	a := v[:medioArray]
	b := v[medioArray:]
	return SumaArray(a) + SumaArray(b)
}
