package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("expected %q but repeated %q", expected, repeated)
	}
}

func TestLastName(t *testing.T) {
	fullName := "Mohammed Sadique"
	lastName := "Sadique"
	if LastName(fullName, lastName) != true {
		t.Errorf("expected %q is a substring of %q", lastName, fullName)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", i)
	}
}
