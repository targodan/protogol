package protoframe

import (
    "net"
)

type State interface {
    NextState(data interface{}, conn net.Conn) (State, error)
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

func (this *StateMachine) Handle(data interface{}, conn net.Conn) (err error) {
    this.current, err = this.initial.NextState(data, conn)
    return
}
