package main

import (
	"encoding/json"
	"github.com/stianeikeland/go-rpio"
)

type Pin struct {
	json.Marshaler
	gpio rpio.Pin
}

var gpios []int = []int{2, 3, 4, 17, 27, 22, 10, 9, 11, 14, 15, 18, 23, 24, 25, 8, 7}
var pins []*Pin

func NewPin(id int64) *Pin {
	p := &Pin{}

	p.gpio = rpio.Pin(id)
	p.gpio.Output()
	p.Close()

	return p
}

func GetPin(id int) *Pin {
	for _, pin := range pins {
		if int(pin.gpio) == id {
			return pin
		}
	}
	return nil
}

func InitializePins(gpios []int) {
	for _, id := range gpios {
		pins = append(pins, NewPin(int64(id)))
	}
}

func (p *Pin) Id() int8 {
	return int8(p.gpio)
}

func (p *Pin) Open() {
	p.gpio.Low()
}

func (p *Pin) Close() {
	p.gpio.High()
}

func (p *Pin) State() int {
	return int(p.gpio.Read())
}

func (p *Pin) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id    int `json:"id"`
		State int `json:"state"`
	}{
		State: p.State(),
		Id:    int(p.gpio),
	})
}
