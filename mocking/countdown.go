package mocking

import (
	"fmt"
	"io"
	"time"
)

const countdownStart = 3
const lastWord = "Go!"

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}
type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, lastWord)

}
