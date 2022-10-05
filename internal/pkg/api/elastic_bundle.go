package api

import (
	"ks-data-prepare/internal/pkg/repository"

	"github.com/gin-gonic/gin"
)

func NewElasticBundleApi(bundleRepo repository.IBundleElasticDatastoreRepository) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		bundle, err := bundleRepo.Find()
		if err != nil {
			RespondFailure(ctx, err)
			return
		}

		RespondSuccess(ctx, bundle)
	}
}
