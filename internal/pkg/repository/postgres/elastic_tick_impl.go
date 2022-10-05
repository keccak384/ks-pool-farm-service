package postgres

import (
	"ks-data-prepare/internal/pkg/entity"

	"gorm.io/gorm"
)

type TickElasticRepository struct {
	db *gorm.DB
}

func NewTickElasticRepository(db *gorm.DB) *TickElasticRepository {
	return &TickElasticRepository{
		db,
	}
}
func (p *TickElasticRepository) SaveBatch(ticks []entity.ElasticTick) ([]entity.ElasticTick, error) {
	res := p.db.CreateInBatches(ticks, 100)
	if res.Error != nil {
		return nil, res.Error
	}
	return ticks, nil
}

func (p *TickElasticRepository) FindAll() ([]entity.ElasticTick, error) {
	var ticks []entity.ElasticTick
	result := p.db.Find(&ticks)
	if result.Error != nil {
		return nil, result.Error
	}
	return ticks, nil
}
