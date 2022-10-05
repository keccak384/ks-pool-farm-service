package api

import (
	"ks-data-prepare/internal/pkg/entity"
	"ks-data-prepare/internal/pkg/repository"

	"github.com/gin-gonic/gin"
)

func NewElasticFarmPoolApi(farmPoolRepo repository.IFarmPoolElasticDatastoreRepository, tokenRepo repository.ITokenElasticDatastoreRepository) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		farmPools, err := farmPoolRepo.FindAll()
		rewardTokensList := make([][]entity.ElasticToken, 0, len(farmPools))
		for _, f := range farmPools {
			rewardTokens, err := tokenRepo.FindByAddresses(f.RewardTokenIDs)
			if err != nil {
				RespondFailure(ctx, err)
				return
			}
			rewardTokensList = append(rewardTokensList, rewardTokens)
		}
		if err != nil {
			RespondFailure(ctx, err)
			return
		}
		RespondSuccess(ctx, toJSONFarmPools(farmPools, rewardTokensList))
	}
}

func toJSONFarmPools(farmPools []entity.ElasticFarmPool, rewardTokens [][]entity.ElasticToken) []entity.ElasticFarmPoolJSON {
	res := make([]entity.ElasticFarmPoolJSON, 0, len(farmPools))
	for index, item := range farmPools {
		res = append(res, item.ToJSON(rewardTokens[index]))
	}
	return res
}
