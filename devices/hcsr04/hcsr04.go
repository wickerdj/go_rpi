package hcsr04

import (
	"log"
	"time"

	"github.com/stianeikeland/go-rpio"
)

// HCSR04 : Model to keep track the GPIO pins to use
type HCSR04 struct {
	EchoPin    rpio.Pin
	TriggerPin rpio.Pin
}

// NewHCSR04 : Sets up a new HC-SR04 sensor
//
// Uses the BCM number system
func NewHCSR04(echoPin int, triggerPin int) (result HCSR04) {
	err := rpio.Open()
	if err != nil {
		panic(err.Error)
	}

	result.EchoPin = rpio.Pin(echoPin)
	result.TriggerPin = rpio.Pin(triggerPin)

	return
}

// Measure : Takes a measurement then returns the distance in centimetre (centimeter or cm)
func (sensor *HCSR04) Measure() float32 {
	sensor.TriggerPin.Output()
	sensor.EchoPin.Output()
	sensor.TriggerPin.Low()
	sensor.EchoPin.Low()

	sensor.EchoPin.Input()

	strobeZero := 0
	strobeOne := 0

	delay(200)
	sensor.TriggerPin.High()
	delay(15)
	sensor.TriggerPin.Low()

	hardStop := 1000000

	log.Printf(" sensor.EchoPin: %v", sensor.EchoPin.Read())

	for strobeZero = 0; strobeZero < hardStop && sensor.EchoPin.Read() != rpio.High; strobeZero++ {
	}
	start := time.Now()
	for strobeOne = 0; strobeOne < hardStop && sensor.EchoPin.Read() != rpio.Low; strobeOne++ {
		delay(1)
	}
	stop := time.Now()

	log.Printf(" sensor.EchoPin: %v", sensor.EchoPin.Read())

	// speed of sound 343 meters per second
	// 343 m/s = 34300cm/s
	// 34300 cm/s = 0.0343 cm/mu (mu=microsecond)
	// 34300 cm/s = 34.3 cm/ms (ms=millisecond)
	// 34300 cm/s = 3.43e-5 cm/ns (ns=nanosecond)
	// Need to account for amount of time to go there and back
	// 17150 is half of 34300

	// dur := stop.Sub(start)
	dur := float32(stop.UnixNano()-start.UnixNano()) / (58.0 * 1000)
	log.Printf("start: %v stop: %v dur: %v", start.UnixNano(), stop.UnixNano(), dur)

	return dur

}

func delay(mu int) {
	time.Sleep(time.Duration(mu) * time.Microsecond)
}
