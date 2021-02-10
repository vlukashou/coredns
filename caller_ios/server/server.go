package server

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/coredns/caddy"

	"github.com/coredns/coredns/coremain"

	caller "github.com/coredns/coredns/caller_ios"
)

const (
	defaultCorefilePath = `coredns/Corefile`
)

type Server struct {
	*caller.Logger
	conf, p string
}

func (c *Server) Setup() {
	c.Logger = &caller.Logger{}
}

// Stop function stops the CoreDNS instance.
func (c *Server) Stop() error {
	return caddy.Stop()
}

// Run function runs the CoreDNS instance. It accepts both configuration string
// and path to the config file. When conf parameter is given, it overwrites contents
// of the Corefile p. Else, the default content of the p is used.
func (c *Server) Run(conf, p string) {

	var err error

	dir, _ := os.Getwd()
	log.Printf(`work_dir: %s`, dir)

	if p == "" {
		p = defaultCorefilePath
	}

	if conf != "" {
		if err = ioutil.WriteFile(p, []byte(conf), 0644); err != nil {
			log.Fatalf(`write_corefile: %v`, err)
		}
	}

	// set Corefile location.
	caddy.DefaultConfigFile = p

	log.Printf(`coremain: starting coredns instance: Corefile: %v`, p)

	coremain.Run()
}
