package resolver

import (
	"net"

	"github.com/miekg/dns"
)

type ResponseWriter struct {
	*dns.Msg
}

func NewResponseWriter() *ResponseWriter {
	return &ResponseWriter{&dns.Msg{}}
}

func (ResponseWriter) LocalAddr() net.Addr {
	return &net.IPAddr{net.IP{127, 0, 0, 1}, ""}
}

func (ResponseWriter) RemoteAddr() net.Addr {
	return &net.IPAddr{net.IP{127, 0, 0, 1}, ""}
}

func (r *ResponseWriter) WriteMsg(m *dns.Msg) error {
	r.Msg = m
	return nil
}

func (r *ResponseWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

func (r *ResponseWriter) Close() error {
	return nil
}

func (r *ResponseWriter) TsigStatus() error {
	return nil
}

func (r *ResponseWriter) TsigTimersOnly(_ bool) {}

func (r *ResponseWriter) Hijack() {}

var (
	_ dns.ResponseWriter = (*ResponseWriter)(nil)
)
