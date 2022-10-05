package builder

import (
	cfg "ks-data-prepare/internal/pkg/config"
	"ks-data-prepare/internal/pkg/constant"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	gin.SetMode(cfg.Instance().Http.Mode)
	server := gin.Default()
	// server.Use(sentrygin.New(sentrygin.Options{
	// 	Repanic: true,
	// }))
	setCORS(server)
	return server
}

func setCORS(engine *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AddAllowMethods(http.MethodOptions)
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders(constant.HeaderAuthorization)
	corsConfig.AddAllowHeaders(constant.HeaderRequestIdLower)
	corsConfig.AddAllowHeaders(constant.HeaderRequestId)
	corsConfig.AddAllowHeaders(constant.HeaderAPIKey)
	engine.Use(cors.New(corsConfig))
}
