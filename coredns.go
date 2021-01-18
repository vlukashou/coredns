package main

//go:generate go run directives_generate.go
//go:generate go run owners_generate.go

import (
	_ "github.com/coredns/coredns/core/plugin" // Plug in CoreDNS.
	"github.com/coredns/coredns/coremain"
)

func main() {
	// starterCoreDNS := coremain.NewStarter(&coremain.CoreDNS{}) 
	// starterCoreDNS.Init() 
	// starterCoreDNS.Start() 
	
	coreDns := coremain.NewCoreDns("corefile")
	coreDns.Init()
	coreDns.Run()
}
