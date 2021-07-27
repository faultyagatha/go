package main

import "fmt"

// ---------------
// Switch stays the same
// we use only implementors
// of Switch interface
//  ---------------
type Switch struct {
	State State
}

func NewSwitch() *Switch {
	return &Switch{NewOffState()}
}

func (s *Switch) On() {
	s.State.On(s)
}

func (s *Switch) Off() {
	s.State.Off(s)
}

// ---------------
// State manages state
// of Switch
//  ---------------
type State interface {
	On(sw *Switch)
	Off(sw *Switch)
}

// ---------------
// State provides default behaviours
// for the State interface
//  ---------------
type BaseState struct{}

func (s *BaseState) On(sw *Switch) {
	fmt.Println("Light is already on")
}

func (s *BaseState) Off(sw *Switch) {
	fmt.Println("Light is already off")
}

type OnState struct {
	BaseState
}

func NewOnState() *OnState {
	fmt.Println("Light turned on")
	return &OnState{BaseState{}}
}

func (o *OnState) Off(sw *Switch) {
	fmt.Println("Turning light off...")
	sw.State = NewOffState()
}

type OffState struct {
	BaseState
}

func NewOffState() *OffState {
	fmt.Println("Light turned off")
	return &OffState{BaseState{}}
}

func (o *OffState) On(sw *Switch) {
	fmt.Println("Turning light on...")
	sw.State = NewOnState()
}

func main() {
	sw := NewSwitch()
	sw.On()
	sw.Off()
	sw.Off()
}
