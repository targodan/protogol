package protoframe

import (
	"net"
)

type ErrorHandler func(err error) (abort bool)

type Client struct {
	chain   Chain
	machine *StateMachine
}

func NewClient(chain Chain, machine *StateMachine) Client {
	return Server{
		chain:   chain,
		machine: machine,
	}
}

func (c Client) Start(baseProto string, addr string, errHandler ErrorHandler) (err error) {
	conn, err := net.Dial(baseProto, addr)
	if err != nil {
		return
	}

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
