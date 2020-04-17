package iteration

func Repeat(character string) string {
	var repeated string
	const repeatcount = 5
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
