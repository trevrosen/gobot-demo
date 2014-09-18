package main

import (
	"fmt"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
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

	// Create our collection of 'robots'
	gbot := gobot.NewGobot()

	// Start our web server for REST-based control/information
	server := api.NewAPI(gbot)
	server.Port = "7337"
	server.Start()

	blinkLedOverAndOver(gbot)
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
func blinkLedOverAndOver(gbot *gobot.Gobot){

	// Create an instance of our chosen adapter type
	// and pass it to the LED driver. The names given here
	// are used in the management functionality.
	beagleboneAdaptor := beaglebone.NewBeagleboneAdaptor("beaglebone")
	led								:= gpio.NewLedDriver(beagleboneAdaptor, "led", "P9_15")

	// Robots in the Gobot colletion run a "work" function
	// when they fire
	work := func() {
		gobot.Every(1 * time.Second, func() {
			led.Toggle()
		})
	}

	// A Robot is a board or device, and is one of the things managed by a Gobot.
	// Here we make one with the adaptor and led objects we made above.
	// The constructor creates a new named robot, provided a connection and a device
	// which will map to something like a GPIO pin.

	// A Robot can be composed of as many Connections and Devices as you like,
	// meaning that you can create something out of a group of supported hardware pieces
	// and treat it in code as a single logical unit.
	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{beagleboneAdaptor},
		[]gobot.Device{led},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
