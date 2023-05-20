package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/chokey2nv/ens-resolve/config"
	"github.com/chokey2nv/ens-resolve/resolver"
	"github.com/gin-gonic/gin"
)

func Server(config *config.Config) {
	// Set up the Gin router
	app := gin.Default()
	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app.GET("resolve/:domain", func(ctx *gin.Context) {
		// example domain = adebimpe.xyz
		domain := ctx.Param("domain")
		resolved, err := resolver.Resolve(config, domain)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": resolved,
		})
	})

	app.Use(CORSMiddleware())
	addr := fmt.Sprintf(":%s", port)
	app.Run(addr)
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
