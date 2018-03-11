# Raspicar project

### Description

This project creates an application that allows to control motors
and read data from sensors attached to RaspberryPi 3.
I've used it to build my own RC car.

### How to build
Linux build command (static linkage)
        
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o raspicar .
    
Raspberry Pi build command (static linkage)

    GOARCH=arm GOARM=5 CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o raspicar .
    
### Installation

Simply copy executable file to your RaspberryPi 3 and run.
After you make sure that Raspicar application is running send
one of following http requests to either steer your device
or collect data from sensors.

Get data from sensors
    
    curl raspihost:8000/
    
Make motors move forward
    
    curl raspihost:8000/motors/-1/-1
    
### Contact

If you have found any issues with this project contact me at: wojtjk@gmail.com
