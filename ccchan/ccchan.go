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

func (c *ccchan) Subscribe(r channel.Receiver, msgs ...interface{}) {
	c.Lock()
	c.ch.Subscribe(r, msgs...)
	c.Unlock()
}

func (c *ccchan) Unsubscribe(r channel.Receiver, msgs ...interface{}) {
	c.Lock()
	c.ch.Unsubscribe(r, msgs...)
	c.Unlock()
}

func (c *ccchan) Send(msg interface{}) {
	c.RLock()
	c.ch.Send(msg)
	c.RUnlock()
}
