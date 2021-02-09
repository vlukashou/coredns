package getcontext

import (
	"sync"

	"github.com/coredns/caddy"
)

func init() {

	caddy.RegisterPlugin("__getcontext", caddy.Plugin{
		ServerType: "dns",
		Action:     func(_ *caddy.Controller) error { return nil },
	})

	caddy.RegisterParsingCallback("dns", "__getcontext", caddy.ParsingCallback(ParsingCallback))
}

var (
	getContext = new(GetContext)
)

type GetContext struct {
	sync.Mutex
	caddy.Context
	slist []caddy.Server
}

func ParsingCallback(ctx caddy.Context) error {
	getContext.Lock()
	defer getContext.Unlock()
	getContext.Context = ctx
	return nil
}

func GetServers() []caddy.Server {

	getContext.Lock()
	defer getContext.Unlock()

	if getContext.slist == nil {
		getContext.slist, _ = getContext.MakeServers()
	}

	return getContext.slist
}
