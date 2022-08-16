package dns

import (
	"demo/core/errors"
	"demo/model/dns"
	"demo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDNS(c *gin.Context) {
	// Initialize default values if any
	var dnsRequest dns.GetDNSRequest
	dnsRequest.Default()

	// NOTE: I would've designed this to be a GET request, but the docs specify JSON input
	// and typically GET requests shouldn't have body data (only path and query params)
	// Bind body model
	if err := c.ShouldBind(&dnsRequest); err != nil {
		restError := errors.BadRequestError(err.Error())
		c.JSON(restError.Status, restError)
		return
	}

	// Validate query model
	if err := dnsRequest.Validate(); err != nil {
		c.JSON(err.Status, err)
		return
	}

	// Parse model
	// NOTE: the frontend sending strings instead of floats is pretty annoying. Makes for some ugly code in this demo
	parsedDNS, err := dnsRequest.Parse()
	if err != nil {
		// Why not return RestError from inside service?
		// Service functions are meant to be reused in the code. They shouldn't know about REST at all.
		// REST related functionality is handled at the router level.
		restError := errors.BadRequestError(err.Error())
		c.JSON(restError.Status, restError)
		return
	}

	result, err := service.DNSService.GetDNS(parsedDNS)
	if err != nil {
		// Throw internal error, will be captured by middleware and turned into InternalServerError automatically
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dns.GetDNSResponse{Loc: result})
}
