package postgres

import (
	"ks-data-prepare/internal/pkg/entity"

	"gorm.io/gorm"
)

type PositionElasticRepository struct {
	db *gorm.DB
}

func NewPositionElasticRepository(db *gorm.DB) *PositionElasticRepository {
	return &PositionElasticRepository{
		db,
	}
}
func (p *PositionElasticRepository) SaveBatch(positions []entity.ElasticPosition) ([]entity.ElasticPosition, error) {
	res := p.db.CreateInBatches(positions, 100)
	if res.Error != nil {
		return nil, res.Error
	}
	return positions, nil
}

func (p *PositionElasticRepository) FindAll() ([]entity.ElasticPosition, error) {
	var positions []entity.ElasticPosition
	result := p.db.Find(&positions)
	if result.Error != nil {
		return nil, result.Error
	}
	return positions, nil
}
