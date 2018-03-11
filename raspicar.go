package main

import (
	"fmt"
	"flag"
	"github.com/stianeikeland/go-rpio"
	"os"
)

func ReadParams() Manager {

	manager := Manager{}

	manager.leftMotor.gpio1 = rpio.Pin(*flag.Int("LMG1", 16, "Left motor GPIO input pin number 1"))
	manager.leftMotor.gpio2 = rpio.Pin(*flag.Int("LMG2", 20, "Left motor GPIO input pin number 2"))
	manager.leftMotor.enable = rpio.Pin(*flag.Int("LME", 21, "Left motor GPIO enable pin number"))

	manager.rightMotor.gpio1 = rpio.Pin(*flag.Int("RMG1", 19, "Right motor GPIO input pin number 1"))
	manager.rightMotor.gpio2 = rpio.Pin(*flag.Int("RMG2", 13, "Right motor GPIO input pin number 2"))
	manager.rightMotor.enable = rpio.Pin(*flag.Int("RME", 26, "Right motor GPIO enable pin number"))

	manager.frontSensor.trigger = rpio.Pin(*flag.Int("FST", 8, "Front sensor trigger pin number"))
	manager.frontSensor.echo = rpio.Pin(*flag.Int("FSE", 7, "Front sensor echo pin number"))

	manager.backSensor.trigger = rpio.Pin(*flag.Int("BST", 23, "Back sensor trigger pin number"))
	manager.backSensor.echo = rpio.Pin(*flag.Int("BSE", 24, "Back sensor echo pin number"))

	manager.bindPort = *flag.String("Port", "8000", "Http server port")

	flag.Parse()

	return manager
}

func main() {

	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()

	manager := ReadParams()
	manager.Run()
}
