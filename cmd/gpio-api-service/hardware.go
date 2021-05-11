package main

import (
	"fmt"
	"log"

	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

type GPIOInterface interface {
	Setup()
	PinOn(pin int) (err error)
	PinOff(pin int) (err error)
	PinValue(pin int, value float32) (err error)
}

type FakeGPIO struct {
}

func (g *FakeGPIO) Setup() {

}

func (g *FakeGPIO) PinOn(_ int) (err error) {
	return nil
}

func (g *FakeGPIO) PinOff(_ int) (err error) {
	return nil
}

func (g *FakeGPIO) PinValue(_ int, _ float32) (err error) {
	return nil
}

type RealGPIO struct {
}

func (g *RealGPIO) Setup() {
	host.Init()
	if _, err := driverreg.Init(); err != nil {
		log.Fatal(err)
	}
	fmt.Print("GPIO pins available:\n")
	for _, p := range gpioreg.All() {
		fmt.Printf("- %s: %s\n", p, p.Function())
	}
}

func (g *RealGPIO) PinOn(pin int) (err error) {

	p := g.getPin(pin)

	if err := p.Out(gpio.High); err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func (g *RealGPIO) PinOff(pin int) (err error) {

	p := g.getPin(pin)

	if err := p.Out(gpio.Low); err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func (g *RealGPIO) PinValue(pin int, value float32) (err error) {
	return
}

func (g *RealGPIO) getPin(pin int) (p gpio.PinIO) {

	p = gpioreg.ByName(fmt.Sprintf("GPIO%d", pin))

	if p == nil {
		log.Printf("Failed to find GPIO%d", pin)
	}
	return
}
