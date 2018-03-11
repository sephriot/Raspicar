package main

import (
	"net/http"
	"log"
	"strings"
	"strconv"
	"fmt"
)

type Manager struct {
	leftMotor Motor
	rightMotor Motor
	frontSensor Sensor
	backSensor Sensor

	bindPort string
}

var (
	requestNumber = 0
)

func (m *Manager) SensorHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling request", requestNumber, r.URL.String())
	requestNumber++
	fmt.Fprintf(w, "{\"front\":\"%f\",\"back\":\"%f\"}", m.frontSensor.GetValue(), m.backSensor.GetValue())
}

func (m *Manager) MotorHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Handling request", requestNumber, r.URL.String())
	requestNumber++
	arr := strings.Split(r.URL.String(), "/")

	if len(arr) <  4 {
		log.Println("Incorrect request", r.URL.String())
		w.WriteHeader(500)
		return
	}

	left, err := strconv.Atoi(arr[2])
	if err != nil {
		log.Println("Incorrect request", r.URL.String())
		w.WriteHeader(500)
		return
	}

	right, err := strconv.Atoi(arr[3])
	if err != nil {
		log.Println("Incorrect request", r.URL.String())
		w.WriteHeader(500)
		return
	}

	if requestNumber % 2 == 0 {
		m.leftMotor.SetDirection(left)
		m.rightMotor.SetDirection(right)
	} else {
		m.rightMotor.SetDirection(right)
		m.leftMotor.SetDirection(left)
	}

	w.Write([]byte("OK"))
}

func (m *Manager) Run() {

	go m.rightMotor.Run()
	go m.leftMotor.Run()
	go m.backSensor.Run()
	go m.frontSensor.Run()

	http.HandleFunc("/", m.SensorHandler)
	http.HandleFunc("/motors/", m.MotorHandler)
	http.ListenAndServe(":"+ m.bindPort, nil)
}