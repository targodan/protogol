package main

import (
	"fmt"
	"protoframe"
	"protoframe/example"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	s := protoframe.NewServer(*example.NewExReaderChain(), *example.NewExWriterChain(), example.NewExServerMachine())
	go s.Start("tcp", "127.0.0.1:7535", func(err error) bool {
		fmt.Printf("Error: %s\n", err.Error())
		return false
	})
}
