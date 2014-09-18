package main

import (
	"fmt"
	"time"

	"github.com/hybridgroup/gobot/platforms/beaglebone"
	"github.com/hybridgroup/gobot/platforms/gpio"
)

const (
	OFF = iota
	ON
)

func main() {
	fmt.Println("[+] Relay going HIGH")
	setDirectPin(ON)
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("[+] Relay going LOW")
	setDirectPin(OFF)
}


// Run some simple GPIO interactions
// by taking control of the Direct Pin abstraction
func setDirectPin(state int){
	beagleboneAdaptor := beaglebone.NewBeagleboneAdaptor("beaglebone")
	gpioPin           := gpio.NewDirectPinDriver(beagleboneAdaptor, "myDevice", "P9_12")

	// Initialize the internal representation of the pinout
	beagleboneAdaptor.Connect()

	// Cast to byte because we are returning an int from a function
	// and not passing in an int literal.
	gpioPin.DigitalWrite(byte(state))
}


// Demos the Gobot Every function, which provides a way
// to trigger recurring functionality.
func blinkLEDmanually(){

}


func robotEverydemo(){

}



