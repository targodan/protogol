package protoframe

import (
    "errors"
    "bufio"
    "net"
)

type Handler interface {
    func Handle(data interface{}) (interface{}, error)
}

func NewChain() *Chain {
    return &Chain {
        first: nil,
        last: nil,
        reader: nil,
        writer: nil,
        conn: nil
    }
}

type Chain struct {
    first *link
    last *link
}

func (c *Chain) AddHandler(h Handler) {
    l := newLink(h);
    if c.first == nil {
        c.first = l
        c.last = l
    } else {
        c.last.next = l
        c.last = l
    }
}

func (c *Chain) GetReaderChain(conn net.Conn) *ReaderChain {
    return &ReaderChain{
        chain: c,
        reader: bufio.NewReader(conn)
    }
}

func (c *Chain) GetWriterChain(onn net.Conn) *ReaderChain {
    return &WriterChain{
        chain: c,
        writer: bufio.NewWriter(conn)
    }
}

type ReaderChain struct {
    chain *Chain
    reader *bufio.Reader
}

type WriterChain struct {
    chain *Chain
    reader *bufio.Writer
}

func (c ReaderChain) RecvPackage() (data Package, err error) {
    if c.chain.first == nil {
        err := errors.New("Chain is empty. Please add Hanlders first.")
        return
    }
    link := c.chain.first
    data := c.reader
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
    nn, err := 0, nil
    if c.chain.first == nil {
        err := errors.New("Chain is empty. Please add Hanlders first.")
        return
    }
    link := c.chain.first
    data := data
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
    next *link
}

func newLink(h Handler) *link {
    return &link {
        handler: h,
        next: nil
    }
}
