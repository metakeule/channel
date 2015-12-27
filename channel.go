package channel

import (
	"reflect"
)

// Receiver is receiving messages.
type Receiver interface {

	// Receive is called when the Receiver is registered and when the corresponding
	// message is sent.
	// when the receiver is registered, it is called with check=true to
	// check, if it can handle the registered message type.
	// When a real message is sent, the Receiver is called with check=false
	Receive(check bool, msg interface{})
}

type ReceiverFunc func(check bool, msg interface{})

func (r ReceiverFunc) Receive(check bool, msg interface{}) {
	r(check, msg)
}

// Channel allows subscription to message types and sending of events
type Channel interface {
	Subscribe(r Receiver, msg ...interface{})
	Unsubscribe(r Receiver, msg ...interface{})
	Send(msg interface{})
}

// New returns a new channel that cannot be used concurrently
func New() Channel {
	return channel(map[string][]Receiver{})
}

/*
this is also a event library. then you can think of Channel as an EventBus, Receiver as
and EventHandler, Sending as Triggering and Messages as Events
*/

//type channel map[reflect.Type][]Receiver
type channel map[string][]Receiver

// Subscribe subscribes a receiver to types of messages
// if msg is not passed, it will receive all messages from the channel
func (c channel) Subscribe(r Receiver, msg ...interface{}) {
	var ty string
	if len(msg) == 0 {
		c[ty] = append(c[ty], r)
		return
	}

	for _, m := range msg {
		r.Receive(true, m) // checks if the notification works
		ty = reflect.TypeOf(m).Name()
		c[ty] = append(c[ty], r)
	}
}

func (c channel) unsubscribe(r Receiver, ty string) {
	rcs, has := c[ty]
	if !has {
		return
	}

	var j int = -1

	for i, rc := range rcs {
		if rc == r {
			j = i
		}
	}

	if j == -1 {
		return
	}

	if j == 0 {
		c[ty] = rcs[1:]
		return
	}

	if j == len(rcs)-1 {
		c[ty] = rcs[:len(rcs)-2]
		return
	}

	c[ty] = append(rcs[0:j], rcs[j+1:]...)
}

func (c channel) Unsubscribe(r Receiver, msg ...interface{}) {

	if len(msg) == 0 {
		c.unsubscribe(r, "")
	}

	for _, m := range msg {
		c.unsubscribe(r, reflect.TypeOf(m).Name())
	}

}

// Send sends the message to the receiver of the message type
// If msg is nil the behaviour is undefined
func (c channel) Send(msg interface{}) {
	for _, r := range c[""] {
		r.Receive(false, msg)
	}

	if msg == nil {
		return
	}

	receivers := c[reflect.TypeOf(msg).Name()]

	for _, r := range receivers {
		r.Receive(false, msg)
	}
}
