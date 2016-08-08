package protoframe

import (
	"bufio"
	"errors"
	"net"
)

type Handler interface {
	Handle(data interface{}) (interface{}, error)
}

type chain struct {
	first *link
	last  *link
}

func (c *chain) AddHandler(h Handler) {
	l := newLink(h)
	if c.first == nil {
		c.first = l
		c.last = l
	} else {
		c.last.next = l
		c.last = l
	}
}

func NewReaderChain() *ReaderChain {
	return new(ReaderChain)
}

func NewWriterChain() *WriterChain {
	return new(WriterChain)
}

type ReaderChain struct {
	chain
	reader *bufio.Reader
	conn   net.Conn
}

type WriterChain struct {
	chain
	writer *bufio.Writer
	conn   net.Conn
}

func (wc WriterChain) Bind(conn net.Conn) WriterChain {
	ret := wc
	ret.writer = bufio.NewWriter(conn)
	ret.conn = conn
	return ret
}

func (rc ReaderChain) Bind(conn net.Conn) ReaderChain {
	ret := rc
	ret.reader = bufio.NewReader(conn)
	ret.conn = conn
	return ret
}

func (c ReaderChain) RecvPackage() (data Package, err error) {
	if c.first == nil {
		err = errors.New("Chain is empty. Please add Hanlders first.")
		return
	}
	link := c.first
	data = c.reader
	for {
		data, err = link.handler.Handle(data)
		if err != nil {
			return
		}
		link = link.next
		if link == nil {
			break
		}
	}
	return
}

func (c WriterChain) SendPackage(data Package) (nn int, err error) {
	nn, err = 0, nil
	if c.first == nil {
		err = errors.New("Chain is empty. Please add Hanlders first.")
		return
	}
	link := c.first
	for {
		data, err = link.handler.Handle(data)
		if err != nil {
			return
		}
		link = link.next
		if link == nil {
			break
		}
	}
	nn, err = c.writer.Write(data.([]byte))
	return
}

type link struct {
	handler Handler
	next    *link
}

func newLink(h Handler) *link {
	return &link{
		handler: h,
		next:    nil,
	}
}
