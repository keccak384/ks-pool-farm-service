package api

import (
	"ks-data-prepare/internal/pkg/entity"
	"ks-data-prepare/internal/pkg/repository"

	"github.com/gin-gonic/gin"
)

func NewElasticTokenApi(tokenRepo repository.ITokenElasticDatastoreRepository) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		tokens, err := tokenRepo.FindAll()
		if err != nil {
			RespondFailure(ctx, err)
			return
		}

		RespondSuccess(ctx, ToJSONTokens(tokens))
	}
}

func ToJSONTokens(tokens []entity.ElasticToken) []entity.ElasticTokenJSON {
	res := make([]entity.ElasticTokenJSON, 0, len(tokens))
	for _, item := range tokens {
		res = append(res, item.ToJSON())
	}
	return res
}
