package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/miekg/dns"

	"github.com/coredns/coredns/caller_ios/resolver"
	"github.com/coredns/coredns/caller_ios/server"
)

var (
	z      string
	t      string
	daemon bool
)

func init() {

	flag.StringVar(&z, "z", "", "the query domain name")
	flag.StringVar(&t, "t", "A", "the query type")
	flag.BoolVar(&daemon, "d", true, "run the server instead of performing query")

	flag.Parse()
}

var conf = `.:1253 {
	forward . 8.8.8.8
	log
	debug
}
`

func main() {

	if z != "" {
		r, err := resolver.New(conf, `/tmp/Corefile`)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer r.Shutdown()

		if t == "" {
			t = "A"
		}

		resp, err := r.Query(dns.Fqdn(z), dns.StringToType[strings.ToUpper(t)])
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("----------------")
		fmt.Println(resp)
		fmt.Println("----------------")
		return
	}

	srv := &server.Server{}
	srv.SetLogOutput(`/tmp/core.log`)
	srv.Run(conf, `/tmp/Corefile`)
}
