// Package hcsr04 : Library for working with the Ultrasonic Ranging Module HC - SR04
//
// From the datasheet:
// Ultrasonic ranging module HC - SR04 provides 2cm - 400cm non-contact
// measurement function, the ranging accuracy can reach to 3mm. The modules
// includes ultrasonic transmitters, receiver and control circuit. The basic principle
// of work:
//		(1) Using IO trigger for at least 10us high level signal,
// 		(2) The Module automatically sends eight 40 kHz and detect whether there is a
// pulse signal back.
// 		(3) IF the signal back, through high level, time of high output IO duration is
// the time from sending ultrasonic to returning.
// Test distance = (high level time×velocity of sound (340M/S) / 2,
//
// You can calculate the range through the time interval between sending trigger signal and
// receiving echo signal. Formula: uS / 58 = centimeters or uS / 148 =inch; or: the
// range = high level time * velocity (340M/S) / 2; we suggest to use over 60ms
// measurement cycle, in order to prevent trigger signal to the echo signal.
package hcsr04

import (
	"time"

	"github.com/stianeikeland/go-rpio"
)

const divisor = 58000.0
const hardStop = 1000000

// HCSR04 : Model to keep track the GPIO pins to use
type HCSR04 struct {
	EchoPin    rpio.Pin
	TriggerPin rpio.Pin
}

// NewHCSR04 : Sets up a new HC-SR04 sensor
//
// Uses the BCM number system
func NewHCSR04(echoPin int, triggerPin int) HCSR04 {
	err := rpio.Open()
	if err != nil {
		panic(err.Error)
	}

	var result HCSR04

	result.EchoPin = rpio.Pin(echoPin)
	result.TriggerPin = rpio.Pin(triggerPin)

	return result
}

// Measure : Takes a measurement then returns the distance in centimeter
//
// Sensor has a resolution of 0.3cm.
func (sensor *HCSR04) Measure() float64 {
	initalizeSensor(sensor)
	delay(6000) // minimum of 5000.
	trigger(sensor)
	dividend := timingCircuit(sensor)

	// Calculation
	// Golang doesn't offer a unix epoch time in microseconds. Need to shift by 1000 to get to microsecond
	// The datasheet says 'uS/58 = centimeters', if I multiple the 58 by 1000 this makes the Formula 'nS/58000 = cm'
	// I can save a little bit of time and resources by calculating the divisor and making it a constant
	dur := dividend / divisor

	return round(dur, 0.05)
}

func round(x, unit float64) float64 {
	// https://stackoverflow.com/questions/39544571/golang-round-to-nearest-0-05/39544897#39544897
	if x > 0 {
		return float64(int64(x/unit+0.5)) * unit
	}
	return float64(int64(x/unit-0.5)) * unit
}

func initalizeSensor(sensor *HCSR04) {
	sensor.TriggerPin.Output()
	sensor.EchoPin.Output()
	sensor.TriggerPin.Low()
	sensor.EchoPin.Low()

	sensor.EchoPin.Input()
}

func trigger(sensor *HCSR04) {
	sensor.TriggerPin.High()
	delay(15) // minimun value of 10.
	sensor.TriggerPin.Low()
}

func timingCircuit(sensor *HCSR04) float64 {
	for i := 0; i < hardStop && sensor.EchoPin.Read() != rpio.High; i++ {
	}

	start := time.Now()
	for i := 0; i < hardStop && sensor.EchoPin.Read() != rpio.Low; i++ {
		delay(1)
	}
	stop := time.Now()

	diff := float64(stop.UnixNano() - start.UnixNano())

	// log.Printf("timingCircuit - start: %v stop: %v diff: %v", start.UnixNano(), stop.UnixNano(), diff)
	return diff
}

func delay(mu int) {
	time.Sleep(time.Duration(mu) * time.Microsecond)
}
