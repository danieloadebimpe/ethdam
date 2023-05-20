package api

import (
	"net/http"

	"github.com/chokey2nv/ens-resolve/config"
	nftResolvers "github.com/chokey2nv/ens-resolve/services/NFTResolvers"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func NFTRoutes(router *gin.RouterGroup, config *config.Config) {

	router.GET("/:address", func(ctx *gin.Context) {
		address := ctx.Param("address")
		var nfts []interface{}
		n, err := nftResolvers.Resolve1155(config, common.HexToAddress(address))
		if err != nil {
			panic(err)
		}
		nfts = append(nfts, n...)
		m, err := nftResolvers.Resolve721(config, common.HexToAddress(address))
		if err != nil {
			panic(err)
		}
		nfts = append(nfts, m...)
		ctx.JSON(http.StatusOK, nfts)
	})
}
