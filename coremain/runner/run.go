package runner

import (
	"github.com/coredns/coredns/coremain"

	_ "github.com/coredns/coredns/core/plugin"
)

// CoremainStarter runs coremain
type CoremainStarter interface {
	Run()
	Init()
	GetLog() string
}

// NewCoreDns makes a new CoreDns
func NewCoreDns() *coremain.CoreDns {
	return &coremain.CoreDns{}
}
