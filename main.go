package main

import (
	"log"

	toggle "github.com/SBhanushali/oapi-gin-api-toggle"
	"github.com/SBhanushali/oapi-gin-api-toggle-example/config"
	"github.com/SBhanushali/oapi-gin-api-toggle-example/server"
	"github.com/SBhanushali/oapi-gin-api-toggle-example/server/api"
	"github.com/SBhanushali/oapi-gin-api-toggle-example/spec"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %s", err)
	}

	swagger, err := spec.GetSwagger()
	if err != nil {
		log.Fatalf("error loading swagger spec: %s", err)
	}

	router := gin.Default()

	toggleConfig := toggle.Config{
		ExtensionName: "x-feature-flag",
		FeatureFlags:  config.FeatureFlags,
	}

	router.Use(toggle.New(swagger, toggleConfig))

	// To use directly with gin
	// router.GET("v1/hello", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "Hello, World!"})
	// })

	// To use with generated server
	svc := api.New()
	handler := server.NewStrictHandler(svc, nil)
	server.RegisterHandlersWithOptions(router, handler, server.GinServerOptions{
		BaseURL: "/v1",
	})

	router.Run(":8080")
}
