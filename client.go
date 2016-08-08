package protogol

import (
	"net"
)

type ErrorHandler func(err error) (abort bool)

type Client struct {
	readerChain ReaderChain
	writerChain WriterChain
	machine     *StateMachine
}

func NewClient(readerChain ReaderChain, writerChain WriterChain, machine *StateMachine) Client {
	return Client{
		readerChain: readerChain,
		writerChain: writerChain,
		machine:     machine,
	}
}

func (c Client) Start(baseProto string, addr string, errHandler ErrorHandler) {
	conn, err := net.Dial(baseProto, addr)
	if err != nil {
		if errHandler(err) {
			return
		}
	}

	machine := CloneStateMachine(*c.machine)

	reader := c.readerChain.Bind(conn)
	writer := c.writerChain.Bind(conn)

	done := false
	for !done {
		done, err = machine.Handle(reader, writer)
		if err != nil {
			done = errHandler(err)
		}
	}
	conn.Close()
}
