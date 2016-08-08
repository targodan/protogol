package protogol

import (
	"fmt"
	"net"
)

func NewServer(readerChain ReaderChain, writerChain WriterChain, machine *StateMachine) Server {
	return Server{
		readerChain: readerChain,
		writerChain: writerChain,
		machine:     machine,
	}
}

type Server struct {
	readerChain ReaderChain
	writerChain WriterChain
	machine     *StateMachine
}

func (s Server) Start(baseProto string, addr string, errHandler ErrorHandler) {
	ln, err := net.Listen(baseProto, addr)
	if err != nil {
		if errHandler(err) {
			return
		}
	}

	for {
		fmt.Println("pls accept")
		conn, err := ln.Accept()
		fmt.Println("accepted")
		if err != nil {
			if errHandler(err) {
				return
			}
		}
		go s.handleConnection(conn, errHandler)
	}
}

func (s Server) handleConnection(conn net.Conn, errHandler ErrorHandler) {
	machine := CloneStateMachine(*s.machine)

	reader := s.readerChain.Bind(conn)
	writer := s.writerChain.Bind(conn)

	done := false
	var err error
	for !done {
		done, err = machine.Handle(reader, writer)
		if err != nil {
			done = errHandler(err)
		}
	}
	conn.Close()
}
