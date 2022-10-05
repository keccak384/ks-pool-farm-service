package postgres

import (
	"ks-data-prepare/internal/pkg/entity"

	"gorm.io/gorm"
)

type PoolElasticRepository struct {
	db *gorm.DB
}

func NewPoolElasticRepository(db *gorm.DB) *PoolElasticRepository {
	return &PoolElasticRepository{
		db,
	}
}
func (p *PoolElasticRepository) SaveBatch(pools []entity.ElasticPool) ([]entity.ElasticPool, error) {
	res := p.db.CreateInBatches(pools, 100)
	if res.Error != nil {
		return nil, res.Error
	}
	return pools, nil
}

func (p *PoolElasticRepository) FindAll() ([]entity.ElasticPool, error) {
	var pools []entity.ElasticPool
	result := p.db.Model(&pools).Preload("Token0").Preload("Token1").Find(&pools)
	if result.Error != nil {
		return nil, result.Error
	}
	return pools, nil
}
