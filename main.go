package main

import (
	"context"
	"demo/app"
	"demo/log"
	v1 "demo/mappings/v1"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleErrors(c *gin.Context) {
	c.Next() // execute all the handlers

	// TODO: Send errors to some error collector, ex. Sentry

	// Log error internally and hide from user
	err := c.Errors.ByType(gin.ErrorTypeAny).Last()
	if err != nil {
		log.Error(err.Error(), err)
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal Server Error",
		})
	}
}

func main() {
	r := app.SetupApp()

	// Only need below code in production, not included in test setup
	// Logging and recovery
	r.Use(log.GinZapLogger(log.GetLogger()), log.GinZapRecovery(log.GetLogger(), true))

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(HandleErrors)

	// Setup routes
	routes_v1 := r.Group("/v1")
	v1.InitializeRoutes(routes_v1)

	// Start server
	srv := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	// Graceful server shutdown
	// https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to initialize server: %v\n", err)
		}
	}()

	log.Info(fmt.Sprintf("Listening on port %v\n", srv.Addr))

	// Wait for kill signal of channel
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// This blocks until a signal is passed into the quit channel
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Info("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: %v\n", err)
	}

}
