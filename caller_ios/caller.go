package caller

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/coredns/caddy"

	"github.com/coredns/coredns/coremain"
)

const (
	defaultLogOutputPath = `coredns/core.log`
	defaultCorefilePath  = `coredns/Corefile`
)

type Server struct {
	conf, p string
	log     *os.File
}

var (
	oldStdout = os.Stdout
	oldStderr = os.Stderr
)

// SetLogOutput function sets the ouput file for the logging instead of defined
// stdout/stderr.
func (c *Server) SetLogOutput(p string) {

	var err error

	if p == "" {
		p = defaultLogOutputPath
	}

	if c.log, err = os.OpenFile(p, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644); err != nil {
		panic(err)
	}

	os.Stdout, os.Stderr = c.log, c.log
	log.SetOutput(c.log)
}

// ResetLogOutput resets the output file for the log by reverting back stdout/stderr.
func (c *Server) ResetLogOutput() {
	if c.log != nil {
		c.log.Close()
	}
	os.Stdout, os.Stderr = oldStdout, oldStderr
	log.SetOutput(os.Stdout)
}

// Stop function stops the CoreDNS instance.
func (c *Server) Stop() error {
	c.ResetLogOutput()
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
