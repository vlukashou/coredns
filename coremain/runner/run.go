package runner

import (
	"github.com/coredns/coredns/coremain"

	_ "github.com/coredns/coredns/plugin/forward"
)

// NewCoreDns makes a new CoreDns
func NewCoreDns() *CoreDns {
	return &CoreDns{c: &coremain.CoreDns{}}
}

type CoreDns struct {
	c *coremain.CoreDns
}

func (c *CoreDns) Run() {
	c.c.Run()
}

func (c *CoreDns) Init() {
	c.c = &coremain.CoreDns{}
	c.c.Init()
}

func (c *CoreDns) GetLog() string {
	return c.c.GetLog()
}
