package servicenames

import (
	"errors"
	"testing"
)

func TestProtocolByNumber(t *testing.T) {
	s := New()

	protocol, err := s.ProtocolByNumber(6)
	if err != nil {
		t.Fatalf("ProtocolByNumber(6) unexpected error: %v", err)
	}
	if protocol.Keyword != "tcp" {
		t.Errorf("ProtocolByNumber(6) Keyword = %q, want %q", protocol.Keyword, "tcp")
	}

	_, err = s.ProtocolByNumber(200) // unassigned protocol number
	if !errors.Is(err, ErrNoProtocol) {
		t.Errorf("ProtocolByNumber(200) err = %v, want %v", err, ErrNoProtocol)
	}
}

func TestProtocolByKeyword(t *testing.T) {
	s := New()

	protocol, err := s.ProtocolByKeyword("tcp")
	if err != nil {
		t.Fatalf("ProtocolByKeyword(\"tcp\") unexpected error: %v", err)
	}
	if protocol.Proto != 6 {
		t.Errorf("ProtocolByKeyword(\"tcp\") Proto = %d, want 6", protocol.Proto)
	}

	_, err = s.ProtocolByKeyword("not-a-real-protocol")
	if !errors.Is(err, ErrNoProtocol) {
		t.Errorf("ProtocolByKeyword(unknown) err = %v, want %v", err, ErrNoProtocol)
	}

	// Regression test: several unassigned protocol numbers (61, 63, 68, 99,
	// 114, 253, 254) share an empty Keyword. The reverse lookup map must
	// skip them so an empty keyword lookup fails cleanly instead of
	// returning whichever one of them was inserted last.
	_, err = s.ProtocolByKeyword("")
	if !errors.Is(err, ErrNoProtocol) {
		t.Errorf("ProtocolByKeyword(\"\") err = %v, want %v", err, ErrNoProtocol)
	}
}

func TestServiceByProtoPort(t *testing.T) {
	s := New()

	service, err := s.ServiceByProtoPort(6, 22)
	if err != nil {
		t.Fatalf("ServiceByProtoPort(6, 22) unexpected error: %v", err)
	}
	if service.Name != "ssh" {
		t.Errorf("ServiceByProtoPort(6, 22) Name = %q, want %q", service.Name, "ssh")
	}
	if len(service.AltServices) != 0 {
		t.Errorf("ServiceByProtoPort(6, 22) AltServices = %+v, want none", service.AltServices)
	}

	_, err = s.ServiceByProtoPort(6, 65535) // unassigned TCP port
	if !errors.Is(err, ErrNoService) {
		t.Errorf("ServiceByProtoPort(6, 65535) err = %v, want %v", err, ErrNoService)
	}

	_, err = s.ServiceByProtoPort(200, 22) // unknown protocol entirely
	if !errors.Is(err, ErrNoService) {
		t.Errorf("ServiceByProtoPort(200, 22) err = %v, want %v", err, ErrNoService)
	}
}

func TestServiceByProtoPortAltServices(t *testing.T) {
	s := New()

	// UDP/2049 is registered under two names by IANA: "shilp" (first in
	// the source data, becomes primary) and "nfs" (becomes an alternate).
	service, err := s.ServiceByProtoPort(17, 2049)
	if err != nil {
		t.Fatalf("ServiceByProtoPort(17, 2049) unexpected error: %v", err)
	}
	if service.Name != "shilp" {
		t.Errorf("ServiceByProtoPort(17, 2049) Name = %q, want %q", service.Name, "shilp")
	}
	if len(service.AltServices) != 1 || service.AltServices[0].Name != "nfs" {
		t.Errorf("ServiceByProtoPort(17, 2049) AltServices = %+v, want one entry named nfs", service.AltServices)
	}
}
