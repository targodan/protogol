package protoframe

import (
    "errors"
)

type Handler interface {
    func Handle(data interface{}) (interface{}, error)
}

func NewChain() *Chain {
    return Chain {
        first: nil,
        last: nil
    }
}

type Chain struct {
    first *link
    last *link
}

func (this *Chain) AddHandler(h Handler) {
    l := newLink(h);
    if this.first == nil {
        this.first = l
    } else if this.last == nil {
        this.first.next = l
        this.last = l
    } else {
        this.last.next = l
        this.last = l
    }
}

func (this Chain) Handle(interface{} data) (data interface{}, err error) {
    if this.first == nil {
        err := errors.New("Chain is empty. Please add Hanlders first.")
        return
    }
    link := this.first
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

type link struct {
    handler Handler
    next *link
}

func newLink(h Handler) *link {
    return link {
        handler: h,
        next: nil
    }
}
