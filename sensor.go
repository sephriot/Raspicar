package main

import (
	"github.com/stianeikeland/go-rpio"
	"time"
)

type Sensor struct {
	trigger rpio.Pin
	echo rpio.Pin
	value float64
}

func (s *Sensor) Run() {

	s.trigger.Output()
	s.echo.Input()

	for {
		s.trigger.High()
		time.Sleep(time.Microsecond * 10)
		s.trigger.Low()
		timeoutIndicator := time.Now()
		startTime := time.Now()
		endTime := time.Now()

		for s.echo.Read() == rpio.Low && time.Since(timeoutIndicator) < time.Second {
			startTime = time.Now()
			time.Sleep(time.Microsecond*100)
		}

		for s.echo.Read() == rpio.High && time.Since(timeoutIndicator) < time.Second {
			endTime = time.Now()
			time.Sleep(time.Microsecond*100)
		}

		dist := endTime.Sub(startTime).Seconds() * 34300 / 2
		if dist > 0 && dist < 1000 {
			s.value = dist
		}
		time.Sleep(time.Second - time.Since(timeoutIndicator))
	}
}

func (s *Sensor) GetValue() float64 {
	return s.value
}