package main

import (
	"fmt"
	"time"

	hc "github.com/wickerdj/go_rpi/devices/hcsr04"
)

func main() {
	// Uses the BCM number system
	h := hc.NewHCSR04(8, 24)

	for true {
		distance := h.Measure()
		fmt.Printf("object is %fcm away\n", distance)
		time.Sleep(time.Second)
	}
}
