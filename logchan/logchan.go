package logchan

import (
	"fmt"
	"github.com/metakeule/channel"
)

type Logger interface {
	Print(vals ...interface{})
}

func New(l Logger, ch channel.Channel) channel.Channel {
	return &logged{l, ch}
}

func (l *logged) Subscribe(r channel.Receiver, msgs ...interface{}) {
	if len(msgs) == 0 {
		l.log.Print(fmt.Sprintf("subscribing receiver %T for all message types", r))
	} else {
		for _, msg := range msgs {
			l.log.Print(fmt.Sprintf("subscribing receiver %T for message type %T", r, msg))
		}
	}
	l.ch.Subscribe(r, msgs...)
}

func (l *logged) Unsubscribe(r channel.Receiver, msgs ...interface{}) {
	if len(msgs) == 0 {
		l.log.Print(fmt.Sprintf("unsubscribing receiver %T from receiving all message types", r))
	} else {
		for _, msg := range msgs {
			l.log.Print(fmt.Sprintf("unsubscribing receiver %T from message type %T", r, msg))
		}
	}

	l.ch.Unsubscribe(r, msgs...)
}

func (l *logged) Send(msg interface{}) {
	l.log.Print(fmt.Sprintf("triggered: %v (%T)", msg, msg))
	l.ch.Send(msg)
}

type logged struct {
	log Logger
	ch  channel.Channel
}
