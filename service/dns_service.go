package service

import (
	"demo/model/dns"
)

// TODO: Extract into env var injected at build time? Depends on usecase
const SectorID = 1

var (
	DNSService IDNSService = &dnsService{}
)

type dnsService struct{}

type IDNSService interface {
	GetDNS(*dns.FindDNS) (float64, error)
}

func (s *dnsService) GetDNS(dnsRequest *dns.FindDNS) (float64, error) {
	loc := dnsRequest.X*SectorID + dnsRequest.Y*SectorID + dnsRequest.Z*SectorID + dnsRequest.Vel
	return loc, nil
}
