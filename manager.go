package main

import (
	"net/http"
	"encoding/json"
	"log"
	"strings"
	"strconv"
)

type Manager struct {
	leftMotor Motor
	rightMotor Motor
	frontSensor Sensor
	backSensor Sensor

	bindPort string
}

type SensorResponse struct {
	front float64
	back  float64
}

func (m *Manager) SensorHandler(w http.ResponseWriter, r *http.Request) {

	sensorData := SensorResponse{}
	sensorData.back = m.backSensor.GetValue()
	sensorData.front = m.frontSensor.GetValue()
	ret, err := json.Marshal(sensorData)
	if err != nil {
		w.WriteHeader(500)
		log.Println("Json parsing error")
		return
	}

	w.Write(ret)
	w.WriteHeader(200)
}

func (m *Manager) MotorHandler(w http.ResponseWriter, r *http.Request) {

	arr := strings.Split(r.URL.String(), "/")

	if len(arr) <  2 {
		log.Println("Incorrect request", r.URL.String())
		w.WriteHeader(500)
		return
	}

	left, err := strconv.Atoi(arr[1])
	if err != nil {
		log.Println("Incorrect request", r.URL.String())
		w.WriteHeader(500)
		return
	}

	right, err := strconv.Atoi(arr[2])
	if err != nil {
		log.Println("Incorrect request", r.URL.String())
		w.WriteHeader(500)
		return
	}

	m.leftMotor.SetDirection(left)
	m.rightMotor.SetDirection(right)

	w.Write([]byte("OK"))
	w.WriteHeader(200)
}

func (m *Manager) Run() {

	go m.rightMotor.Run()
	go m.leftMotor.Run()
	go m.backSensor.Run()
	go m.frontSensor.Run()

	http.HandleFunc("/sensors", m.SensorHandler)
	http.HandleFunc("/", m.MotorHandler)
	http.ListenAndServe(":"+ m.bindPort, nil)
}