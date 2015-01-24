// Common package converts common units
// into units of SI (International System).
package common

func Fahrenheit(x float64) float64 {
	return (x + 459.67) * 5 / 9
}

func Celsius(x float64) float64 {
	return x + 273.15
}

func Horsepower(x float64) float64 {
	return x * 746
}

func Watt(x float64) float64 {
	return x
}

func Kmh(x float64) float64 {
	return x * 0.27777777777778
}

func Mph(x float64) float64 {
	return x * 0.44704
}

func Kn(x float64) float64 {
	return x * 0.514444
}

func Mps(x float64) float64 {
	return x
}

/*
* carat
* gram
* ounce
* pound
* stone
 */

func Carat(x float64) float64 {
	return x * 0.2
}
func Gram(x float64) float64 {
	return x
}
func Ounce(x float64) float64 {
	return x * 28
}
func Pound(x float64) float64 {
	return x * 453.59237
}
func Stone(x float64) float64 {
	return x * 6350.29318
}
