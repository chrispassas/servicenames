package main

import (
	"log"

	"github.com/chrispassas/servicenames"
)

func main() {

	s := servicenames.New()

	service, _ := s.ServiceByProtoPort(6, 22)
	log.Printf("proto:%d port:%d name:%s description:%s", service.Protocol, service.Port, service.Name, service.Description)

	service, _ = s.ServiceByProtoPort(17, 53)
	log.Printf("proto:%d port:%d name:%s description:%s", service.Protocol, service.Port, service.Name, service.Description)

	p, _ := s.ProtocolByNumber(uint8(6))
	log.Printf("proto:%d keyword:%s", p.Proto, p.Keyword)

	p, _ = s.ProtocolByNumber(uint8(17))
	log.Printf("proto:%d keyword:%s", p.Proto, p.Keyword)

	p, _ = s.ProtocolByKeyword("icmp")
	log.Printf("proto:%d keyword:%s", p.Proto, p.Keyword)

}
