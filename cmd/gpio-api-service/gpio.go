package main

type GPIOPin interface {
	On(value float32) (success bool, err error)
	Off() (success bool, err error)
	Name() (name string)
	Pin() (pin int)
}

type Relay struct {
	pin  int
	name string
}

func (r *Relay) Name() string                           { return r.name }
func (r *Relay) Pin() int                               { return r.pin }
func (r *Relay) On(_ float32) (success bool, err error) { return }
func (r *Relay) Off() (success bool, err error)         { return }

type PWM struct {
	pin  int
	name string
}

func (p *PWM) Name() string                               { return p.name }
func (p *PWM) Pin() int                                   { return p.pin }
func (p *PWM) On(value float32) (success bool, err error) { return }
func (p *PWM) Off() (success bool, err error)             { return }

func PinFactory(pinType string, gpio int, name string) (p GPIOPin) {
	switch pinType {
	case "relay":
		p = &Relay{gpio, name}
	case "pwm":
		p = &PWM{gpio, name}
	}
	return
}
