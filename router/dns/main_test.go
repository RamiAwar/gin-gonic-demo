package dns

import (
	"demo/app"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func TestMain(m *testing.M) {
	r = app.SetupApp()

	routes_v1 := r.Group("/v1")
	routes_v1.POST("/dns/find", GetDNS)

	exitVal := m.Run()
	os.Exit(exitVal)
}
