package api

import (
	"net/http"

	ens "github.com/chokey2nv/ens-resolve/services/ENSResolvers"
	"github.com/chokey2nv/ens-resolve/config"
	"github.com/gin-gonic/gin"
)

func ENSRoutes(router *gin.RouterGroup, config *config.Config) {

	router.GET("/:domain", func(ctx *gin.Context) {
		domain := ctx.Param("domain")
		record, err := ens.Resolve(config, domain)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, record)
	})
}
