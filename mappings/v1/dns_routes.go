package v1

import (
	dns_router "demo/router/dns"

	"github.com/gin-gonic/gin"
)

func SetupDNSRoutes(router *gin.RouterGroup) {
	router.POST("/dns/find", dns_router.GetDNS)
}
