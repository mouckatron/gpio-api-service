package main

import "testing"

func setup() (*[]GPIOPin, GPIOInterface) {
	pins := make([]GPIOPin, 0)
	hw := FakeGPIO{}
	return &pins, &hw
}

func TestAddRelay(t *testing.T) {
	pins, _ := setup()
	//addRelay(pins *[int]Pin, name string, gpio int)

	pinId := 1
	name := "MyRelay"

	err := addRelay(pins, pinId, "MyRelay")

	if err != nil {
		t.Errorf("Wanted pin added, got error %s", err)
	}

	if l := len(*pins); l != 1 {
		t.Errorf("want len(pins) == 1, got %d", l)
	}

	/* if ok := (*pins)[0]; !ok {
		t.Error("want pins[0] is our pin, it wasn't there, got %s", pins)
	} */

	if p := (*pins)[0]; p.Name() != name {
		t.Errorf("Pin name %s, not %s", p.Name, name)
	}
}

func TestAddRelayIncorrectID(t *testing.T) {}

func TestDeleteRelay(t *testing.T) {

}

func TestAddPWM(t *testing.T) {

}

func TestDeletePWM(t *testing.T) {

}

func TestSetName(t *testing.T) {

}

func TestSetStateOn(t *testing.T) {

}

func TestSetStateOff(t *testing.T) {

}

func TestSetPWMValue(t *testing.T) {

}

func TestSetLockClosed(t *testing.T) {

}

func TestSetLockOpen(t *testing.T) {

}

func TestChangeStateWhithLockOpen(t *testing.T) {

}

func TestChangeStateWithLockClosed(t *testing.T) {

}
