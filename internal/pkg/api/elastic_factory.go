package api

import (
	"ks-data-prepare/internal/pkg/entity"
	"ks-data-prepare/internal/pkg/repository"

	"github.com/gin-gonic/gin"
)

func NewElasticFactoryApi(factoryRepo repository.IFactoryElasticDatastoreRepository) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		factory, err := factoryRepo.Find()
		if err != nil {
			RespondFailure(ctx, err)
			return
		}

		RespondSuccess(ctx, getResp(factory))
	}
}

type ElasticFactoryResp struct {
	Id                           string
	Vid                          int
	Block                        int
	PoolCount                    string
	TxCount                      string
	TotalVolumeUsd               float64
	TotalVolumeEth               float64
	TotalFeesUsd                 float64
	TotalFeesEth                 float64
	UntrackedVolumeUsd           float64
	TotalValueLockedUsd          float64
	TotalValueLockedEth          float64
	TotalValueLockedUsdUntracked float64
	TotalValueLockedEthUntracked float64
	Owner                        string
}

func getResp(f entity.ElasticFactory) ElasticFactoryResp {
	return ElasticFactoryResp{
		Id:                           f.Id,
		Vid:                          f.Vid,
		Block:                        f.Block,
		PoolCount:                    f.PoolCount.Int.String(),
		TotalVolumeUsd:               f.TotalVolumeUsd,
		TotalVolumeEth:               f.TotalVolumeEth,
		TotalFeesUsd:                 f.TotalFeesUsd,
		TotalFeesEth:                 f.TotalFeesEth,
		UntrackedVolumeUsd:           f.UntrackedVolumeUsd,
		TotalValueLockedUsd:          f.TotalValueLockedUsd,
		TotalValueLockedEth:          f.TotalValueLockedEth,
		TotalValueLockedUsdUntracked: f.TotalValueLockedUsdUntracked,
		Owner:                        f.Owner,
	}
}
