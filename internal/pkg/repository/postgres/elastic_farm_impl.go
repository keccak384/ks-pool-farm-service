package postgres

import (
	"ks-data-prepare/internal/pkg/entity"

	"gorm.io/gorm"
)

type FarmElasticRepository struct {
	db *gorm.DB
}

func NewFarmElasticRepository(db *gorm.DB) *FarmElasticRepository {
	return &FarmElasticRepository{
		db,
	}
}
func (p *FarmElasticRepository) SaveBatch(farms []entity.ElasticFarm) ([]entity.ElasticFarm, error) {
	res := p.db.CreateInBatches(farms, 100)
	if res.Error != nil {
		return nil, res.Error
	}
	return farms, nil
}

func (p *FarmElasticRepository) FindAll() ([]entity.ElasticFarm, error) {
	var farms []entity.ElasticFarm
	result := p.db.Find(&farms)
	if result.Error != nil {
		return nil, result.Error
	}
	return farms, nil
}
