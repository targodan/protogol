package protoframe

import (
    "fmt"
)

type State interface {
    NextState(data interface{}) (State, error)
}

func NewStateMachine(initial State) *StateMachine {
    return StateMachine{
        initial: initial,
        current: initial
    }
}

type StateMachine {
    initial State
    currnet State
}

func (this *StateMachine) Handle(data interface{}) (err error) {
    this.current, err = this.initial.NextState(data)
    return
}
