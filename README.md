# go_rpi
Reboot of a set experiments, libraries, and tools for working with a Raspberry Pi. 

## build command

For Raspberry Pi A, A+, B, B+, Zero
GOARM=6 GOARCH=arm GOOS=linux go build {filename.go}

Raspberry Pi 2, 3
GOARM=7 GOARCH=arm GOOS=linux go build {filename.go}

### Cross compile with a C library.

CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7 go build {filename.go}

## GO Packages for working with the Raspberry Pi GPIO

gobot - https://gobot.io/documentation/platforms/raspi/

go-rpio - https://github.com/stianeikeland/go-rpio

## Notes on pin

When using the gobot library refer to https://pinout.xyz

## General Reference

RPi Low-level peripherals - https://elinux.org/RPi_Low-level_peripherals

Raspberry Pi on Wikipedia - https://en.wikipedia.org/wiki/Raspberry_Pi
