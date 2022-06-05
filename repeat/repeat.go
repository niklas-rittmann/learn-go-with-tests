package iteration

// Repeat given string n times
func Repeat(text string, times int) (repeated string) {
	for i := 0; i < times; i++ {
		repeated += text
	}
	return
}
