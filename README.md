[![GoDoc](https://pkg.go.dev/github.com/chrispassas/servicenames?status.svg)](https://pkg.go.dev/github.com/chrispassas/servicenames)


# servicenames
This module uses www.iana.org data to lookup protocol or service names using protocol and port numbers.

Data Sources:

https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml
https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.xhtml



## Example
Look up service or protocol

```go
package main

import (
	"log"

	"github.com/chrispassas/servicenames"
)

func main() {

	s := servicenames.New()

	service, _ := s.ServiceByProtoPort(6, 22)
	log.Printf("proto:%d port:%d name:%s description:%s", service.Protocol, service.Port, service.Name, service.Description)
	// proto:6 port:22 name:ssh description:the secure shell (ssh) protocol

	service, _ = s.ServiceByProtoPort(17, 53)
	log.Printf("proto:%d port:%d name:%s description:%s", service.Protocol, service.Port, service.Name, service.Description)
	// proto:17 port:53 name:domain description:domain name server

	p, _ := s.ProtocolByNumber(uint8(6))
	log.Printf("proto:%d keyword:%s", p.Proto, p.Keyword)
	//  proto:6 keyword:tcp

	p, _ = s.ProtocolByNumber(uint8(17))
	log.Printf("proto:%d keyword:%s", p.Proto, p.Keyword)
	// proto:17 keyword:udp

	p, _ = s.ProtocolByKeyword("icmp")
	log.Printf("proto:%d keyword:%s", p.Proto, p.Keyword)
	// proto:1 keyword:icmp

}
```
