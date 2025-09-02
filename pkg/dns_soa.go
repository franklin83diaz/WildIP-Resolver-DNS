package pkg

import (
	"time"

	"github.com/miekg/dns"
)

func apexSOA(zone string, ns string) dns.RR {
	return &dns.SOA{
		Hdr:     dns.RR_Header{Name: dns.Fqdn(zone), Rrtype: dns.TypeSOA, Class: dns.ClassINET, Ttl: 3600},
		Ns:      ns + ".",
		Mbox:    "hostmaster." + dns.Fqdn(zone),
		Serial:  uint32(time.Now().Unix()), // o lleva tu propio contador
		Refresh: 3600, Retry: 600, Expire: 604800, Minttl: 300,
	}
}
