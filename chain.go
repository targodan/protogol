package protogol

import (
	"bufio"
	"errors"
	"net"
)

type Package struct {
	Parent *Package
	Data   interface{}
}

func (p Package) Pack(data interface{}) Package {
	return Package{Parent: &p, Data: data}
}

func (p Package) Unpack() Package {
	return *p.Parent
}

type Handler func(pkg Package) (Package, error)

type chain struct {
	handlers []Handler
}

func (c *chain) AddHandler(h Handler) {
	c.handlers = append(c.handlers, h)
}

func NewReaderChain() *ReaderChain {
	return &ReaderChain{chain: chain{make([]Handler, 0, 1)}}
}

func NewWriterChain() *WriterChain {
	return &WriterChain{chain: chain{make([]Handler, 0, 1)}}
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

func (c ReaderChain) RecvPackage() (pkg Package, err error) {
	if len(c.handlers) == 0 {
		err = errors.New("Chain is empty. Please add Hanlders first.")
		return
	}
	d := c.reader
	for _, handler := range c.handlers {
		pkg, err = handler(Package{nil, d})
		if err != nil {
			return
		}
	}
	return
}

func (c WriterChain) SendPackage(pkg Package) (nn int, err error) {
	nn, err = 0, nil
	if len(c.handlers) == 0 {
		err = errors.New("Chain is empty. Please add Hanlders first.")
		return
	}
	for _, handler := range c.handlers {
		pkg, err = handler(pkg)
		if err != nil {
			return
		}
	}
	nn, err = c.writer.Write(pkg.Data.([]byte))
	c.writer.Flush()
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
