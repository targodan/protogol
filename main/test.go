package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/targodan/protogol"
	"github.com/targodan/protogol/example"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var mode string
	if len(os.Args) < 2 {
		mode = "server"
	} else {
		mode = os.Args[1]
	}
	if mode == "server" {
		s := protogol.NewServer(*example.NewExReaderChain(), *example.NewExWriterChain(), example.NewExServerMachine())
		s.Start("tcp", "127.0.0.1:7535", func(err error) bool {
			fmt.Printf("Error: %s\n", err.Error())
			return false
		})
	} else if mode == "client" {
		c := protogol.NewClient(*example.NewExReaderChain(), *example.NewExWriterChain(), example.NewExClientMachine())
		c.Start("tcp", "127.0.0.1:7535", func(err error) bool {
			fmt.Printf("Error: %s\n", err.Error())
			return false
		})
	} else {
		fmt.Println("Please tell me if I should be a server or a client.")
	}
}
