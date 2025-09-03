package pkg

import (
	"WildIP-Resolver-DNS/pkg/config"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/miekg/dns"
)

func DNS() {
	zone := config.Fqdn
	addr := config.Address + ":" + fmt.Sprint(config.Port)
	ttl := config.TTL
	ns := config.NS

	mux := dns.NewServeMux()
	// Manager Zone
	mux.HandleFunc(zone, func(w dns.ResponseWriter, r *dns.Msg) {
		msg := new(dns.Msg)
		msg.SetReply(r)
		msg.Authoritative = true

		if len(r.Question) == 0 {
			_ = w.WriteMsg(msg)
			return
		}

		//for each question
		for _, q := range r.Question {

			// Handle SOA requests
			if q.Qtype == dns.TypeSOA {
				soa := apexSOA(zone, ns)
				msg.Answer = append(msg.Answer, soa)
				_ = w.WriteMsg(msg)
				return
			}
			// Handle NS requests
			if q.Qtype == dns.TypeNS {
				ns := apexNS(zone, ns)
				msg.Answer = append(msg.Answer, ns)
				_ = w.WriteMsg(msg)
				return
			}

			// Just respond to A (and ANY returning A)
			if q.Qtype != dns.TypeA && q.Qtype != dns.TypeANY {
				_ = w.WriteMsg(msg)
				return
			}

			name := dns.Fqdn(q.Name)
			zoneFQDN := dns.Fqdn(zone)

			if !dns.IsSubDomain(zoneFQDN, name) {
				msg.Rcode = dns.RcodeNameError // NXDOMAIN
				_ = w.WriteMsg(msg)
				return
			}

			// extract local part
			local := strings.TrimSuffix(name, zoneFQDN)
			local = strings.TrimSuffix(local, ".")

			// Support dashes and underscores as separators.
			ipText := strings.NewReplacer("-", ".", "_", ".").Replace(local)

			ip := net.ParseIP(ipText)
			if ip == nil || ip.To4() == nil {
				msg.Rcode = dns.RcodeNameError // NXDOMAIN if it doesn't match the pattern
				_ = w.WriteMsg(msg)
				return
			}

			rr := &dns.A{
				Hdr: dns.RR_Header{
					Name:   name,
					Rrtype: dns.TypeA,
					Class:  dns.ClassINET,
					Ttl:    uint32(ttl),
				},
				A: ip.To4(),
			}
			msg.Answer = append(msg.Answer, rr)
			_ = w.WriteMsg(msg)
		}
	})

	// Run the DNS server
	udpSrv := &dns.Server{Addr: addr, Net: "udp", Handler: mux}
	tcpSrv := &dns.Server{Addr: addr, Net: "tcp", Handler: mux}

	go func() {
		if err := udpSrv.ListenAndServe(); err != nil {
			log.Fatalf("UDP listen failed: %v", err)
		}
	}()
	if err := tcpSrv.ListenAndServe(); err != nil {
		log.Fatalf("TCP listen failed: %v", err)
	}
}
