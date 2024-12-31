package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const sleep = "sleep"
const write = "write"

type SpySleeper struct {
	Calls int
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to go", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleepePrinter := &SpyCountdownOperations{}

		Countdown(buffer, spySleepePrinter)
		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}

	})
	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}

		Countdown(spySleepPrinter, spySleepPrinter)
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %q got %q", want, spySleepPrinter.Calls)
		}
	})
}
func TestConfigurableSleeper(t *testing.T) {
	spyTime := &SpyTime{}
	sleepTime := 5 * time.Second
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("Should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}

}
