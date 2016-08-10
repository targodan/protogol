package example

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/targodan/protogol"

	"github.com/Workiva/go-datastructures/queue"
)

type packageReader struct{}

func (p packageReader) Handle(data interface{}) (interface{}, error) {
	reader := data.(*bufio.Reader)
	var len uint32
	err := binary.Read(reader, binary.LittleEndian, &len)
	if err != nil {
		return nil, err
	}

	stringbuf := new(bytes.Buffer)
	var read uint32
	for read = 0; read < len; read++ {
		r, _, err := reader.ReadRune()
		if err != nil {
			return nil, err
		}
		stringbuf.WriteRune(r)
	}
	return stringbuf.String(), nil
}

type packageWriter struct{}

func (p packageWriter) Handle(data interface{}) (interface{}, error) {
	msg := data.(string)

	buf := new(bytes.Buffer)

	len := uint32(utf8.RuneCountInString(msg))
	err := binary.Write(buf, binary.LittleEndian, len)
	if err != nil {
		return nil, err
	}

	for i := uint32(0); i < len; i++ {
		r, size := utf8.DecodeRuneInString(msg)
		msg = msg[size:]
		buf.WriteRune(r)
	}

	return buf.Bytes(), nil
}

func NewExReaderChain() *protogol.ReaderChain {
	c := protogol.NewReaderChain()
	c.AddHandler(packageReader{})
	return c
}

func NewExWriterChain() *protogol.WriterChain {
	c := protogol.NewWriterChain()
	c.AddHandler(packageWriter{})
	return c
}

type serverState struct{}

func (s serverState) NextState(reader protogol.ReaderChain, writer protogol.WriterChain, stack *queue.Queue) (protogol.State, error) {
	tmp, err := reader.RecvPackage()
	if err != nil {
		return nil, err
	}
	msg := tmp.(string)
	_, err = writer.SendPackage(msg)
	if err != nil {
		return nil, err
	}
	return s, nil
}

type clientState struct {
	reader *bufio.Reader
}

func (s clientState) NextState(reader protogol.ReaderChain, writer protogol.WriterChain, stack *queue.Queue) (protogol.State, error) {
	msg, err := s.reader.ReadString('\n')

	_, err = writer.SendPackage(msg)
	if err != nil {
		return nil, err
	}
	tmp, err := reader.RecvPackage()
	msg = tmp.(string)
	if err != nil {
		return nil, err
	}
	fmt.Println(msg)
	return s, nil
}

func NewExServerMachine() *protogol.StateMachine {
	return protogol.NewStateMachine(serverState{}, nil)
}

func NewExClientMachine() *protogol.StateMachine {
	return protogol.NewStateMachine(clientState{bufio.NewReader(os.Stdin)}, nil)
}
