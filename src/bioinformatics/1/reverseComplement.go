package main

func ReverseComplement(Pattern string) string {
	return Reverse(Complement(Pattern))
} // top down programming

func Complement(Pattern string) string {

	sliceOfComplement := make([]byte, len(Pattern))

	for i, symbol := range Pattern {
		char := byte(symbol)

		switch char {
		case 'T':
			sliceOfComplement[i] = 'A'
		case 'A':
			sliceOfComplement[i] = 'T'
		case 'C':
			sliceOfComplement[i] = 'G'
		case 'G':
			sliceOfComplement[i] = 'C'
		}
	}

	return string(sliceOfComplement)
}

func Reverse(Pattern string) string {
	sliceOfReverse := make([]byte, len(Pattern))

	for i := len(Pattern); i > 0; i-- {
		sliceOfReverse[len(Pattern)-i] = Pattern[i-1]
	}

	return string(sliceOfReverse)
}
