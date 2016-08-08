package protoframe

import (
    "errors"
)

type Handler interface {
    func Handle(data interface{}) (interface{}, error)
}

func NewChain() *Chain {
    return &Chain {
        first: nil,
        last: nil
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
    // wrap it around
    c.last.next = c.first
}

func (c Chain) Handle(interface{} data, chan<- out) (err error) {
    if c.first == nil {
        err := errors.New("Chain is empty. Please add Hanlders first.")
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
            out <- data
        }
    }
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
