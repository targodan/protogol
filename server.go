package protoframe

import (
	"bufio"
	"net"
)

func NewServer(chain Chain, machine *StateMachine) Server {
	return Server{
		chain:   chain,
		machine: machine,
	}
}

type Server struct {
	chain   Chain
	machine *StateMachine
}

func (s Server) Start(baseProto string, addr string, errHandler ErrorHandler) (err error) {
	ln, err := net.Listen(baseProto, addr)
	if err != nil {
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		go s.handleConnection(conn, errHandler)
	}
}

func (s Server) handleConnection(conn net.Conn, errHandler ErrorHandler) {
	machine := CloneStateMachine(*s.machine)

	reader := s.chain.GetReaderChain(conn)
	writer := s.chain.GetWriterChain(conn)

	done := false
	for !done {
		done, err = machine.Handle(reader, writer)
		if err != nil {
			done = errHandler(err)
		}
	}
	conn.Close()
}
