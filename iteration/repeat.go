package iteration

import "strings"

func Repeat(character string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated = repeated + character
	}
	return repeated
}
func LastName(firstName, lastName string) bool {
	return strings.Contains(firstName, lastName)
}
