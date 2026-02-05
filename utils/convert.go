package utils

func Converter(source string, target string, quantity float64) float64 {
	var result float64
	if source == "USD" {
		result = Data.Rates[target] * quantity
		return result
	}
	if target == "USD" {
		result = quantity / Data.Rates[source]
		return result
	}
	// The rates we are given based on USD.
	// What if I want to learn 50 eur to gbp
	// Here is the logic
	result = (1.0 / Data.Rates[source]) * Data.Rates[target] * quantity
	return result

}
