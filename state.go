package protoframe

import (
    "net"
    "github.com/Workiva/go-datastructures/queue"
)

type State interface {
    NextState(reader *ReaderChain, writer *WriterChain, stack *queue.Queue) (State, error)
}

func NewStateMachine(initial State, final State) *StateMachine {
    return StateMachine{
        initial: initial,
        current: initial,
        final: final,
        stack: queue.New(0)
    }
}

func CloneStateMachine(m StateMachine) *StateMachine {
    return StateMachine{
        initial: m.initial,
        current: m.current,
        final: m.final,
        stack: queue.New(0)
    }
}

type StateMachine {
    initial State
    currnet State
    final State
    stack *queue.Queue
}

func (sm *StateMachine) Handle(reader *ReaderChain, writer *WriterChain) (done bool, err error) {
    sm.current, err = this.initial.NextState(reader, writer, sm.stack)
    done := sm.current == sm.final
    return
}
