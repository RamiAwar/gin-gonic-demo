package service

import (
	"demo/model/dns"
	"testing"
)

// Test business logic here with edge cases
// We don't have any constraints on DNS values or location values for now, so no edge cases to test
// Coords and velocity can be negative in this case

func TestGetDNS(t *testing.T) {
	findDNS := dns.FindDNS{X: 1, Y: 2, Z: 3, Vel: 4}
	got, err := DNSService.GetDNS(&findDNS)
	if err != nil {
		t.Fatal(err)
	}

	want := float64(1 + 2 + 3 + 4)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
