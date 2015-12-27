package bcchan

import (
	"github.com/metakeule/channel"
)

func New() channel.Channel {
	b := bcchan([]channel.Receiver{})
	return &b
}

type bcchan []channel.Receiver

func (b *bcchan) Subscribe(r channel.Receiver, msg ...interface{}) {
	*b = append(*b, r)
}

func (b *bcchan) Unsubscribe(r channel.Receiver, msg ...interface{}) {
	var j int = -1

	for i, rc := range *b {
		if rc == r {
			j = i
		}
	}

	if j == -1 {
		return
	}

	if j == 0 {
		*b = (*b)[1:]
		return
	}

	if j == len(*b)-1 {
		*b = (*b)[:len(*b)-2]
		return
	}

	*b = append((*b)[0:j], (*b)[j+1:]...)
}

func (b *bcchan) Send(msg interface{}) {
	for _, r := range *b {
		r.Receive(false, msg)
	}

}
