package protoframe

import (
    "net"
)

type Client struct {
    machine *StateMachine
}

func (c Client) Start(baseProto string, addr string) (err error) {
    conn, err := net.Dial(baseProto, addr)
    if err != nil {
        return
    }
}
