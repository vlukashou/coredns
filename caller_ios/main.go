package caller

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/coredns/caddy"

	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/coremain"

	// _ "github.com/coredns/caddy/onevent"
	// _ "github.com/coredns/coredns/plugin/acl"
	// _ "github.com/coredns/coredns/plugin/any"
	// _ "github.com/coredns/coredns/plugin/auto"
	// _ "github.com/coredns/coredns/plugin/autopath"
	// _ "github.com/coredns/coredns/plugin/azure"
	// _ "github.com/coredns/coredns/plugin/bind"
	// _ "github.com/coredns/coredns/plugin/bufsize"
	// _ "github.com/coredns/coredns/plugin/cache"
	// _ "github.com/coredns/coredns/plugin/cancel"
	// _ "github.com/coredns/coredns/plugin/chaos"
	// _ "github.com/coredns/coredns/plugin/clouddns"
	_ "github.com/coredns/coredns/plugin/debug"
	// _ "github.com/coredns/coredns/plugin/dns64"
	// _ "github.com/coredns/coredns/plugin/dnssec"
	// _ "github.com/coredns/coredns/plugin/dnstap"
	// _ "github.com/coredns/coredns/plugin/erratic"
	_ "github.com/coredns/coredns/plugin/errors"
	// _ "github.com/coredns/coredns/plugin/etcd"
	// _ "github.com/coredns/coredns/plugin/file"
	_ "github.com/coredns/coredns/plugin/forward"
	// _ "github.com/coredns/coredns/plugin/grpc"
	// _ "github.com/coredns/coredns/plugin/health"
	// _ "github.com/coredns/coredns/plugin/hosts"
	// _ "github.com/coredns/coredns/plugin/k8s_external"
	// _ "github.com/coredns/coredns/plugin/kubernetes"
	// _ "github.com/coredns/coredns/plugin/loadbalance"
	// _ "github.com/coredns/coredns/plugin/local"
	_ "github.com/coredns/coredns/plugin/log"
	// _ "github.com/coredns/coredns/plugin/loop"
	// _ "github.com/coredns/coredns/plugin/metadata"
	// _ "github.com/coredns/coredns/plugin/metrics"
	// _ "github.com/coredns/coredns/plugin/nsid"
	// _ "github.com/coredns/coredns/plugin/pprof"
	// _ "github.com/coredns/coredns/plugin/ready"
	// _ "github.com/coredns/coredns/plugin/reload"
	// _ "github.com/coredns/coredns/plugin/rewrite"
	// _ "github.com/coredns/coredns/plugin/root"
	// _ "github.com/coredns/coredns/plugin/route53"
	// _ "github.com/coredns/coredns/plugin/secondary"
	// _ "github.com/coredns/coredns/plugin/sign"
	// _ "github.com/coredns/coredns/plugin/template"
	// _ "github.com/coredns/coredns/plugin/tls"
	// _ "github.com/coredns/coredns/plugin/trace"
	// _ "github.com/coredns/coredns/plugin/transfer"
	// _ "github.com/coredns/coredns/plugin/whoami"
)

var directives = []string{
	// "metadata",
	// "cancel",
	// "tls",
	// "reload",
	// "nsid",
	// "bufsize",
	// "root",
	// "bind",
	"debug",
	// "trace",
	// "ready",
	// "health",
	// "pprof",
	// "prometheus",
	"errors",
	"log",
	// "dnstap",
	// "local",
	// "dns64",
	// "acl",
	// "any",
	// "chaos",
	// "loadbalance",
	// "cache",
	// "rewrite",
	// "dnssec",
	// "autopath",
	// "template",
	// "transfer",
	// "hosts",
	// "route53",
	// "azure",
	// "clouddns",
	// "k8s_external",
	// "kubernetes",
	// "file",
	// "auto",
	// "secondary",
	// "etcd",
	// "loop",
	"forward",
	// "grpc",
	// "erratic",
	// "whoami",
	// "on",
	// "sign",
}

const (
	DefaultLogOutputPath = `coredns/core.log`
	DefaultCorefilePath  = `coredns/Corefile`
)

type CoreDns struct {
	logFile, corefile string
	config            string
}

func (c *CoreDns) SetLogOutput(p string) {
	c.logFile = p
}

func (c *CoreDns) SetCorefilePath(p string) {
	c.corefile = p
}

func (c *CoreDns) Run(corednsConfig string) {

	var err error

	if c.logFile != "" {
		// set log output.
		if os.Stdout, err = os.OpenFile(c.logFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644); err != nil {
			panic(err)
		}

		os.Stderr = os.Stdout

		defer os.Stdout.Close()

		log.SetOutput(os.Stdout)
	}

	dir, _ := os.Getwd()
	log.Printf(`getwd: %s`, dir)

	if c.corefile == "" {
		c.corefile = DefaultCorefilePath
	}

	if err = ioutil.WriteFile(c.corefile, []byte(corednsConfig), 0644); err != nil {
		log.Fatalf(`write_corefile: %v`, err)
	}

	// set directives.
	dnsserver.Directives = directives

	// set Corefile location.
	caddy.DefaultConfigFile = c.corefile

	log.Printf(`coremain: run`)

	coremain.Run()
}
