package logchan_test

import (
	"fmt"
	"github.com/metakeule/channel"
	"github.com/metakeule/channel/logchan"
	"log"
	"os"
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
	l := log.New(os.Stdout, "INFO: ", log.Lshortfile)
	ch := logchan.New(l, channel.New())
	ch.Subscribe(printer{}, message(""), eMailAddress(""), eMail{})
	ch.Send(message("Hello World!"))
	ch.Unsubscribe(printer{}, eMailAddress(""))
	ch.Send(eMailAddress("test@example.com"))
	ch.Send(eMail{"Hello", "World", "sender@example.com", "receiver@example.com"})

	// Output:
	//
	// INFO: logchan.go:21: subscribing receiver logchan_test.printer for message type logchan_test.message
	// INFO: logchan.go:21: subscribing receiver logchan_test.printer for message type logchan_test.eMailAddress
	// INFO: logchan.go:21: subscribing receiver logchan_test.printer for message type logchan_test.eMail
	// INFO: logchan.go:40: triggered: Hello World! (logchan_test.message)
	// got message: "Hello World!"
	// INFO: logchan.go:32: unsubscribing receiver logchan_test.printer from message type logchan_test.eMailAddress
	// INFO: logchan.go:40: triggered: test@example.com (logchan_test.eMailAddress)
	// INFO: logchan.go:40: triggered: {Hello World sender@example.com receiver@example.com} (logchan_test.eMail)
	// got eMail from sender@example.com: Hello World
}
