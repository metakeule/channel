package bcchan_test

import (
	"fmt"
	"github.com/metakeule/channel/bcchan"
)

type message string
type eMailAddress string
type eMail struct {
	subject, text, from, to string
}

type printer string

func (p printer) Receive(check bool, msg interface{}) {
	switch m := msg.(type) {
	case message:
		if !check {
			fmt.Printf("%s: message: %#v\n", p, m)
		}
	case eMailAddress:
		if !check {
			fmt.Printf("%s: eMailAddress: %#v\n", p, m)
		}
	case eMail:
		if !check {
			fmt.Printf("%s: eMail from %s: %s %s\n", p, m.from, m.subject, m.text)
		}
	default:
		panic(fmt.Sprintf("%s: unsupported: %T", p, msg))
	}
}

func ExampleBroadcastChannel() {
	ch := bcchan.New()
	ch.Subscribe(printer("A"))
	ch.Subscribe(printer("B"))
	ch.Send(message("Hello World!"))
	ch.Send(eMailAddress("test@example.com"))
	ch.Unsubscribe(printer("A"))
	ch.Send(eMail{"Hello", "World", "sender@example.com", "receiver@example.com"})

	// Output:
	//
	// A: message: "Hello World!"
	// B: message: "Hello World!"
	// A: eMailAddress: "test@example.com"
	// B: eMailAddress: "test@example.com"
	// B: eMail from sender@example.com: Hello World
}
