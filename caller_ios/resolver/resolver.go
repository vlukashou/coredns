package resolver

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"
	"unsafe"

	"github.com/coredns/caddy"
	"github.com/miekg/dns"

	caller "github.com/coredns/coredns/caller_ios"
	"github.com/coredns/coredns/caller_ios/resolver/getcontext"
)

const (
	serverType = "dns"
)

// Server ...
type server interface {
	// ServeDNS ...
	ServeDNS(context.Context, dns.ResponseWriter, *dns.Msg)
}

// Resolver structure is a DNS resolver based on the CoreDNS/Caddy plugin system.
// You can configure the resolve to be a blocking forwarding-proxy.
type Resolver struct {
	*caller.Logger

	inst *caddy.Instance
	s    server
}

// Resolve function resolves the raw DNS datagram and returns the bytearray
// of the response DNS datagram.
func (r *Resolver) Resolve(p []byte) ([]byte, error) {

	var (
		w = newResponseWriter()
		m = new(dns.Msg)

		err error
	)

	if err = m.Unpack(p); err != nil {
		return nil, err
	}

	r.s.ServeDNS(context.Background(), w, m)
	return w.Pack()
}

// Query function resolves the DNS query of type t for the domainname z.
func (r *Resolver) Query(z string, t int) (*dns.Msg, error) {

	var (
		w = newResponseWriter()
		m = new(dns.Msg)
	)

	m.SetQuestion(z, uint16(t))

	r.s.ServeDNS(context.Background(), w, m)
	return w.Msg, nil
}

// Shutdown function stops running goroutines initiated by plugins (like forward).
func (r *Resolver) Shutdown() {
	for _, onShutdown := range r.inst.OnShutdown {
		onShutdown()
	}
}

// NewForConfigFile function returns the new resolver for the config file path p.
func NewForConfigFile(p string) (*Resolver, error) {

	var (
		b   []byte
		err error
	)

	if b, err = ioutil.ReadFile(p); err != nil {
		return nil, err
	}

	return New(string(b), p)

}

func (r *Resolver) Setup(c, p string) error {

	var (
		ret *Resolver
		err error
	)

	if ret, err = New(c, p); err != nil {
		return err
	}

	if r.inst != nil {
		r.Shutdown()
	}

	*r = *ret

	return nil
}

// New function returns the new resolver for the config string c and config file path p
// (which is the path where Corefile will be written).
func New(c, p string) (*Resolver, error) {

	var (
		r = new(Resolver)

		err error
	)

	r.Logger = &caller.Logger{}

	// replace ending with getcontext stanza.
	c = strings.Replace(c, "\n}", "\n    __getcontext\n}", -1)

	if err = ioutil.WriteFile(p, []byte(c), 0644); err != nil {
		return nil, err
	}

	input := caddy.CaddyfileInput{
		Contents:       []byte(c),
		Filepath:       p,
		ServerTypeName: serverType,
	}

	r.inst = new(caddy.Instance)

	// NOTE: setting this directive allows us to instantinate plugins.
	*(*string)(unsafe.Pointer(uintptr(unsafe.Pointer(r.inst)))) = serverType

	if err = caddy.ValidateAndExecuteDirectives(input, r.inst, false); err != nil {
		return nil, err
	}

	servers := getcontext.GetServers()

	if len(servers) == 0 {
		return nil, fmt.Errorf(`no servers found for the resolver. conf: %v`, c)
	}

	for _, onStartup := range r.inst.OnStartup {
		if err = onStartup(); err != nil {
			return nil, err
		}
	}

	r.s = servers[0].(server)

	return r, nil
}
