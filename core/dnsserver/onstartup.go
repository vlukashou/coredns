package dnsserver

import (
	"fmt"
	"strings"
	"net"
)


// startUpZones creates the text that we show when starting up:
// grpc://example.com.:1055
// example.com.:1053 on 127.0.0.1
func startUpZones(protocol, addr string, zones map[string]*Config) string {
	s := ""
	var ip, port string

	var err error

	for zone := range zones {
		// split addr into protocol, IP and Port
		parts := strings.Split(addr, "://")

		switch len(parts) {
			case 1:
				ip, port, err = net.SplitHostPort(parts[0])
			
			case 2:
				ip, port, err = net.SplitHostPort(parts[1])
		
			default:
				ip = ""
				port = ""
		        err = fmt.Errorf("provided value is not in an address format : %s", addr)
		}


		if err != nil {
			// this should not happen, but we need to take care of it anyway
			s += fmt.Sprintln(protocol + zone + ":" + addr)
			continue
		}
		if ip == "" {
			s += fmt.Sprintln(protocol + zone + ":" + port)
			continue
		}
		// if the server is listening on a specific address let's make it visible in the log,
		// so one can differentiate between all active listeners
		s += fmt.Sprintln(protocol + zone + ":" + port + " on " + ip)
	}
	return s
}
