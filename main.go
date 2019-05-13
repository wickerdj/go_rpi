package main

import (
	"fmt"
	"time"

	hc "github.com/wickerdj/go_rpi/devices/hcsr04"
)

func exampleUsingHcsr() {
	// Uses the BCM number system
	h := hc.NewHCSR04(8, 24)

	for true {
		distance := h.Measure()
		fmt.Printf("object is %.2fcm away\n", distance)
		time.Sleep(time.Second)
	}
}

func main() {
	exampleUsingHcsr()
}
