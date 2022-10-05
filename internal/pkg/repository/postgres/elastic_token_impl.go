package postgres

import (
	"ks-data-prepare/internal/pkg/entity"

	"gorm.io/gorm"
)

type TokenElasticRepository struct {
	db *gorm.DB
}

func NewTokenElasticRepository(db *gorm.DB) *TokenElasticRepository {
	return &TokenElasticRepository{
		db,
	}
}
func (p *TokenElasticRepository) SaveBatch(tokens []entity.ElasticToken) ([]entity.ElasticToken, error) {
	res := p.db.CreateInBatches(tokens, 100)
	if res.Error != nil {
		return nil, res.Error
	}
	return tokens, nil
}

func (p *TokenElasticRepository) FindAll() ([]entity.ElasticToken, error) {
	var tokens []entity.ElasticToken
	result := p.db.Find(&tokens)
	if result.Error != nil {
		return nil, result.Error
	}
	return tokens, nil
}

func (p *TokenElasticRepository) FindByAddresses(addresses []string) ([]entity.ElasticToken, error) {
	var tokens []entity.ElasticToken
	result := p.db.Where("id IN ?", addresses).Find(&tokens)
	if result.Error != nil {
		return nil, result.Error
	}
	return tokens, nil
}
