package channel_test

import (
	"fmt"
	"github.com/metakeule/channel"
	"github.com/metakeule/channel/logchan"
	"os"
	// for concurrency
	// "github.com/metakeule/channel/ccchan"
	// for broadcasting
	// "github.com/metakeule/channel/bcchan"
	"log"
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
	ch := logchan.New(log.New(os.Stdout, "INFO: ", log.Lshortfile), channel.New())
	ch.Subscribe(printer{}, message(""), eMailAddress(""), eMail{})
	ch.Send(message("Hello World!"))
	ch.Send(eMailAddress("test@example.com"))
	ch.Send(eMail{"Hello", "World", "sender@example.com", "receiver@example.com"})

	// Output:
	//
	// INFO: logchan.go:17: subscribing receiver channel_test.printer for message type []interface {}
	// INFO: logchan.go:27: triggered: Hello World! (channel_test.message)
	// got message: "Hello World!"
	// INFO: logchan.go:27: triggered: test@example.com (channel_test.eMailAddress)
	// got eMailAddress: "test@example.com"
	// INFO: logchan.go:27: triggered: {Hello World sender@example.com receiver@example.com} (channel_test.eMail)
	// got eMail from sender@example.com: Hello World
}
