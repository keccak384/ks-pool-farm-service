package api

import (
	"ks-data-prepare/internal/pkg/entity"
	"ks-data-prepare/internal/pkg/repository"

	"github.com/gin-gonic/gin"
)

func NewElasticPoolApi(poolRepo repository.IPoolElasticDatastoreRepository) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		pools, err := poolRepo.FindAll()
		if err != nil {
			RespondFailure(ctx, err)
			return
		}

		RespondSuccess(ctx, toJSONPools(pools))
	}
}

func toJSONPools(pools []entity.ElasticPool) []entity.ElasticPoolJSON {
	res := make([]entity.ElasticPoolJSON, 0, len(pools))
	for _, item := range pools {
		res = append(res, item.ToJSON())
	}
	return res
}
