package protoframe

import (
    "net"
    "bufio"
)

func NewServer(chain Chain, machine *StateMachine) Server {
    return Server{
        chain: chain,
        machine: machine
    }
}

type Server struct {
    chain Chain
    machine *StateMachine
}

func (s Server) Start(baseProto string, addr string) (err error) {
    ln, err := net.Listen(baseProto, addr)
    if err != nil {
        return
    }

    for {
        conn, err := ln.Accept()
        if err != nil {
            return
        }
    	go s.handleConnection(conn)
    }
}

func (s Server) handleConnection(conn net.Conn) {
    m := *s.machine
    machine := &m

    buffer := bufio.NewReader(conn)
    io := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
    data := make(chan interface{})
    go s.chain.Handle(io, data)

    for {
        machine.Handle(conn, <-data)
    }
}
