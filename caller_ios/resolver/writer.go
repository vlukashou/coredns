package resolver

import (
	"net"

	"github.com/miekg/dns"
)

// ResponseWriter is a dns.ResponseWriter mock to retrieve the response for
// the DNS query.
type ResponseWriter struct {
	*dns.Msg
}

// NewResponseWriter function returns new ResponseWriter instance.
func NewResponseWriter() *ResponseWriter {
	return &ResponseWriter{&dns.Msg{}}
}

// LocalAddr ...
func (ResponseWriter) LocalAddr() net.Addr {
	return &net.IPAddr{net.IP{127, 0, 0, 1}, ""}
}

// RemoteAddr ...
func (ResponseWriter) RemoteAddr() net.Addr {
	return &net.IPAddr{net.IP{127, 0, 0, 1}, ""}
}

// WriteMsg ...
func (r *ResponseWriter) WriteMsg(m *dns.Msg) error {
	r.Msg = m
	return nil
}

// Write ...
func (r *ResponseWriter) Write(p []byte) (int, error) {
	return len(p), r.Msg.Unpack(p)
}

// Close ...
func (r *ResponseWriter) Close() error {
	return nil
}

// TsigStatus ...
func (r *ResponseWriter) TsigStatus() error {
	return nil
}

// TsigTimersOnly ...
func (r *ResponseWriter) TsigTimersOnly(_ bool) {}

// Hijack ...
func (r *ResponseWriter) Hijack() {}

var (
	_ dns.ResponseWriter = (*ResponseWriter)(nil)
)
