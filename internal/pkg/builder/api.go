package builder

import (
	"ks-data-prepare/internal/pkg/api"
	"ks-data-prepare/internal/pkg/config"
	"ks-data-prepare/internal/pkg/db"
	"ks-data-prepare/internal/pkg/repository/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
)

type apiServer struct {
	server *gin.Engine
}

func NewApi() IRunner {
	server := NewServer()
	v1 := server.Group("/api/v1")
	v1.GET("/live", func(c *gin.Context) { c.AbortWithStatus(http.StatusOK) })
	v1.GET("/ready", func(c *gin.Context) { c.AbortWithStatus(http.StatusOK) })

	elasticPoolDatastore := postgres.NewPoolElasticRepository(db.Instance())
	elasticFarmPoolDatastore := postgres.NewFarmPoolElasticRepository(db.Instance())
	elasticBundleDatastore := postgres.NewBundleElasticRepository(db.Instance())
	elasticFactoryDatastore := postgres.NewFactoryElasticRepository(db.Instance())
	elasticTokenDatastore := postgres.NewTokenElasticRepository(db.Instance())

	v1.GET("/elastic/bundle", api.NewElasticBundleApi(elasticBundleDatastore))
	v1.GET("/elastic/factory", api.NewElasticFactoryApi(elasticFactoryDatastore))
	v1.GET("/elastic/pools", api.NewElasticPoolApi(elasticPoolDatastore))
	v1.GET("/elastic/farmPools", api.NewElasticFarmPoolApi(elasticFarmPoolDatastore, elasticTokenDatastore))
	v1.GET("/elastic/tokens", api.NewElasticTokenApi(elasticTokenDatastore))
	return &apiServer{server: server}
}

func (f *apiServer) Run() error {
	return f.server.Run(config.Instance().Http.BindAddress)
}
