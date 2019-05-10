# go_rpi
Rebooting of experiments, libraries, and tools for working with a Raspberry Pi. 

## Installation
`go get -d -u github.com/wickerdj/go_rpi`

If the above throws an error, remove the switches -d and -u and try again. The new command would be `go get github.com/wickerdj/go_rpi`. Double check that files were downloaded.

## Documentation
Not yet

## Build Command
### For Raspberry Pi A, A+, B, B+, Zero
`GOARM=6 GOARCH=arm GOOS=linux go build {filename.go}`

### For Raspberry Pi 2, 3
`GOARM=7 GOARCH=arm GOOS=linux go build {filename.go}`

### Cross compile with a C library.
`CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7 go build {filename.go}`

# Devices
## HC-SR04 Sensor
### Datasheet 
* [Ultrasonic Ranging Module HC-SR04](https://cdn.sparkfun.com/datasheets/Sensors/Proximity/HCSR04.pdf)
* [Ultrasonic ranging module : HC-SR04 version 2](https://www.electroschematics.com/wp-content/uploads/2013/07/HC-SR04-datasheet-version-2.pdf)
* [HC-SR04 User Guide](https://www.mpja.com/download/hc-sr04_ultrasonic_module_user_guidejohn.pdf)

### Helpful Blog posts
* [HC-SR04 Ultrasonic Range Sensor on the Raspberry Pi](https://www.modmypi.com/blog/hc-sr04-ultrasonic-range-sensor-on-the-raspberry-pi)
* [Using a Raspberry Pi distance sensor (ultrasonic sensor HC-SR04)](https://tutorials-raspberrypi.com/raspberry-pi-ultrasonic-sensor-hc-sr04/)
* [Raspberry Pi Distance Sensor: How to setup the HC-SR04](https://pimylifeup.com/raspberry-pi-distance-sensor/)
* [Ultrasonic Sensor (HC-SR04) + Raspberry Pi](https://classes.engineering.wustl.edu/ese205/core/index.php?title=Ultrasonic_Sensor_(HC-SR04)_%2B_Raspberry_Pi)

# Misc
## Raspberry Pi
* [Raspberry Pi hardware](https://www.raspberrypi.org/documentation/hardware/raspberrypi/)

## GO Packages for working with the Raspberry Pi 
* [gobot](https://gobot.io/documentation/platforms/raspi/)
* [go-rpio](https://github.com/stianeikeland/go-rpio)
* [go-wiringpi](https://github.com/eternal-flame-AD/go-wiringpi)
* [periph - Peripherals I/O in go](https://periph.io) - [github](https://github.com/google/periph)

## Notes on pin
* [Raspberry Pi Pinout](https://pinout.xyz) - [github](https://github.com/Gadgetoid/Pinout.xyz)

## General Reference
* [RPi Low-level peripherals](https://elinux.org/RPi_Low-level_peripherals)
* [Raspberry Pi on Wikipedia](https://en.wikipedia.org/wiki/Raspberry_Pi)
