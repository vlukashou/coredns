package resolver

import (
	"net"

	"github.com/miekg/dns"
)

// responseWriter is a dns.ResponseWriter mock to retrieve the response for
// the DNS query.
type responseWriter struct {
	*dns.Msg
}

// newResponseWriter function returns new ResponseWriter instance.
func newResponseWriter() *responseWriter {
	return &responseWriter{&dns.Msg{}}
}

// LocalAddr ...
func (responseWriter) LocalAddr() net.Addr {
	return &net.IPAddr{net.IP{127, 0, 0, 1}, ""}
}

// RemoteAddr ...
func (responseWriter) RemoteAddr() net.Addr {
	return &net.IPAddr{net.IP{127, 0, 0, 1}, ""}
}

// WriteMsg ...
func (r *responseWriter) WriteMsg(m *dns.Msg) error {
	r.Msg = m
	return nil
}

// Write ...
func (r *responseWriter) Write(p []byte) (int, error) {
	return len(p), r.Msg.Unpack(p)
}

// Close ...
func (r *responseWriter) Close() error {
	return nil
}

// TsigStatus ...
func (r *responseWriter) TsigStatus() error {
	return nil
}

// TsigTimersOnly ...
func (r *responseWriter) TsigTimersOnly(_ bool) {}

// Hijack ...
func (r *responseWriter) Hijack() {}

var (
	_ dns.ResponseWriter = (*responseWriter)(nil)
)
