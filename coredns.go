package main

import (
	"github.com/coredns/coredns/caller_ios"
)

var conf = `.:1253 {
	forward . 8.8.8.8
	log
	debug
}
`

func main() {

	coredns := &caller.CoreDns{}
	coredns.SetLogOutput(`/tmp/core.log`)
	coredns.SetCorefilePath(`/tmp/Corefile`)

	coredns.Run(conf)
}
