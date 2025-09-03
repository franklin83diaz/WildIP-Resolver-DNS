package pkg

import (
	"github.com/miekg/dns"
)

func apexNS(zone string, ns string) dns.RR {
	print(zone + " " + ns)
	return &dns.NS{
		Hdr: dns.RR_Header{
			Name:   zone + ".",
			Rrtype: dns.TypeNS,
			Class:  dns.ClassINET,
			Ttl:    3600,
		},
		Ns: ns + ".",
	}
}
