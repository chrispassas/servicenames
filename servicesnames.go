package servicenames

import "fmt"

var (
	ErrNoService  = fmt.Errorf("unknown service")
	ErrNoProtocol = fmt.Errorf("unknown protocol")
)

type Services struct {
	lookupServices         map[uint8]map[uint16]Service
	lookupProtocols        map[uint8]Protocol
	reverseLookupProtocols map[string]Protocol
}

type Service struct {
	Protocol    uint8  `json:"protocol"`
	Port        uint16 `json:"port"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// AltServices holds other names/descriptions IANA registers for this
	// same Protocol+Port pair (e.g. UDP/2049 is registered as both "nfs"
	// and "shilp"). It is only populated on the primary Service returned
	// by ServiceByProtoPort and is empty for the vast majority of ports
	// that have a single registered name.
	AltServices []Service `json:"alt_services,omitempty"`
}

type Protocol struct {
	Proto   uint8  `json:"proto"`
	Keyword string `json:"keyword"`
}

func New() *Services {
	s := &Services{
		lookupServices:         make(map[uint8]map[uint16]Service),
		lookupProtocols:        make(map[uint8]Protocol),
		reverseLookupProtocols: make(map[string]Protocol),
	}

	for _, p := range protocolData {
		s.lookupProtocols[p.Proto] = p

		// Several unassigned protocol numbers share an empty keyword.
		// Skip them so they don't collide in the reverse lookup map.
		if p.Keyword == "" {
			continue
		}
		s.reverseLookupProtocols[p.Keyword] = p
	}

	for _, sn := range serviceData {
		if _, ok := s.lookupServices[sn.Protocol]; !ok {
			s.lookupServices[sn.Protocol] = make(map[uint16]Service)
		}

		existing, ok := s.lookupServices[sn.Protocol][sn.Port]
		if !ok {
			s.lookupServices[sn.Protocol][sn.Port] = sn
			continue
		}

		// IANA registers more than one name/description for this
		// Protocol+Port pair. Keep the first entry as primary and stash
		// the rest as alternates instead of silently discarding them.
		existing.AltServices = append(existing.AltServices, sn)
		s.lookupServices[sn.Protocol][sn.Port] = existing
	}

	return s
}

func (s *Services) ProtocolByKeyword(keyword string) (protocol Protocol, err error) {
	var ok bool

	if protocol, ok = s.reverseLookupProtocols[keyword]; !ok {
		err = ErrNoProtocol
		return protocol, err
	}

	return protocol, err
}

func (s *Services) ProtocolByNumber(proto uint8) (protocol Protocol, err error) {
	var ok bool

	if protocol, ok = s.lookupProtocols[proto]; !ok {
		err = ErrNoProtocol
		return protocol, err
	}

	return protocol, err
}

func (s *Services) ServiceByProtoPort(proto uint8, port uint16) (service Service, err error) {
	var ok bool

	if _, ok = s.lookupServices[proto]; !ok {
		err = ErrNoService
		return service, err
	}

	if service, ok = s.lookupServices[proto][port]; !ok {
		err = ErrNoService
		return service, err
	}

	return service, err
}
