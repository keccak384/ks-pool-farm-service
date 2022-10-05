package postgres

import (
	"fmt"
	"ks-data-prepare/internal/pkg/entity"

	"gorm.io/gorm"
)

type FarmPoolElasticRepository struct {
	db *gorm.DB
}

func NewFarmPoolElasticRepository(db *gorm.DB) *FarmPoolElasticRepository {
	return &FarmPoolElasticRepository{
		db,
	}
}
func (p *FarmPoolElasticRepository) SaveBatch(farmPools []entity.ElasticFarmPool) ([]entity.ElasticFarmPool, error) {

	res := p.db.CreateInBatches(farmPools, 100)
	if res.Error != nil {
		fmt.Println("[err] save farm pools", res.Error)
		return nil, res.Error
	}
	return farmPools, nil
}

func (p *FarmPoolElasticRepository) FindAll() ([]entity.ElasticFarmPool, error) {
	var farmPools []entity.ElasticFarmPool
	result := p.db.Model(&farmPools).Preload("Pool").Preload("Pool.Token0").Preload("Pool.Token1").Find(&farmPools)
	if result.Error != nil {
		return nil, result.Error
	}
	return farmPools, nil
}
