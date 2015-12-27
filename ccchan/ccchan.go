package ccchan

import (
	"github.com/metakeule/channel"
	"sync"
)

// New returns a new channel that can be used concurrently
func New() channel.Channel {
	return &ccchan{ch: channel.New()}
}

type ccchan struct {
	sync.RWMutex
	ch channel.Channel
}

func (c *ccchan) Subscribe(r channel.Receiver, msg ...interface{}) {
	c.Lock()
	c.ch.Subscribe(r, msg...)
	c.Unlock()
}

func (c *ccchan) Send(msg interface{}) {
	c.RLock()
	c.ch.Send(msg)
	c.RUnlock()
}

func (c *ccchan) Unsubscribe(r channel.Receiver, msg ...interface{}) {
	c.Lock()
	c.ch.Unsubscribe(r, msg...)
	c.Unlock()
}
