package main

import (
	"github.com/stianeikeland/go-rpio"
	"time"
)

const (
	Forward = -1
	Stop = 0
	Backward = 1
)

type Motor struct {

	gpio1 rpio.Pin
	gpio2 rpio.Pin
	enable rpio.Pin
	direction int
	lastRequest time.Time
}

func (m *Motor) Run() {

	m.direction = 0
	m.gpio1.Output()
	m.gpio2.Output()
	m.enable.Output()
	sleepTime := time.Second / 10

	for {
		if m.direction == Forward {
			m.gpio1.Low()
			m.gpio2.High()
			m.enable.High()
		} else if m.direction == Backward {
			m.gpio1.High()
			m.gpio2.Low()
			m.enable.High()
		} else {
			m.enable.Low()
		}

		if time.Since(m.lastRequest) > sleepTime {
			m.direction = Stop
			time.Sleep(sleepTime)
		}
	}
}

func (m *Motor) SetDirection(direction int) {
	if direction < 0 {
		m.direction = Forward
	} else if direction > 0 {
		m.direction = Backward
	}

	m.lastRequest = time.Now()
}