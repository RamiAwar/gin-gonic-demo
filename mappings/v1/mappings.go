package v1

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(g *gin.RouterGroup) {
	SetupDNSRoutes(g)
}
