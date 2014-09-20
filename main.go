package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/miekg/dns"
)

func handleRequest(w dns.ResponseWriter, r *dns.Msg) {
	q := r.Question[0]

	info := fmt.Sprintf("Question: Type=%s Class=%s Name=%s", dns.TypeToString[q.Qtype], dns.ClassToString[q.Qclass], q.Name)

	if q.Qtype == dns.TypeA && q.Qclass == dns.ClassINET {
		m := new(dns.Msg)
		m.SetReply(r)
		a := new(dns.A)
		a.Hdr = dns.RR_Header{Name: q.Name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 600}
		a.A = net.IPv4(127, 0, 0, 1)
		m.Answer = []dns.RR{a}
		w.WriteMsg(m)
		log.Printf("%s (RESOLVED)\n", info)
	} else {
		m := new(dns.Msg)
		m.SetReply(r)
		m.Rcode = dns.RcodeNameError // NXDOMAIN
		w.WriteMsg(m)
		log.Printf("%s (NXDOMAIN)\n", info)
	}
}

func main() {
	var addr = flag.String("addr", "127.0.0.1:5300", "listen address")
	flag.Parse()

	server := &dns.Server{Addr: *addr, Net: "udp"}
	server.Handler = dns.HandlerFunc(handleRequest)
	log.Printf("Listening on %s\n", *addr)
	log.Fatal(server.ListenAndServe())
}
