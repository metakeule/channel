package ccchan_test

import (
	"fmt"
	"github.com/metakeule/channel/ccchan"
)

type message string
type eMailAddress string
type eMail struct {
	subject, text, from, to string
}

type printer struct{}

func (printer) Receive(check bool, msg interface{}) {
	switch m := msg.(type) {
	case message:
		if !check {
			fmt.Printf("got message: %#v\n", m)
		}
	case eMailAddress:
		if !check {
			fmt.Printf("got eMailAddress: %#v\n", m)
		}
	case eMail:
		if !check {
			fmt.Printf("got eMail from %s: %s %s\n", m.from, m.subject, m.text)
		}
	default:
		panic(fmt.Sprintf("unsupported: %T", msg))
	}
}

func ExampleChannel() {
	ch := ccchan.New()
	ch.Subscribe(printer{}, message(""), eMailAddress(""), eMail{})
	stop := make(chan bool)
	go func() {
		ch.Send(message("Hello World!"))
		ch.Unsubscribe(printer{}, eMail{})
		stop <- true
	}()
	<-stop
	ch.Send(eMailAddress("test@example.com"))
	ch.Send(eMail{"Hello", "World", "sender@example.com", "receiver@example.com"})

	// Output:
	//
	// got message: "Hello World!"
	// got eMailAddress: "test@example.com"
}
