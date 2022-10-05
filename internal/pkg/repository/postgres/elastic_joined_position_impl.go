package postgres

import (
	"ks-data-prepare/internal/pkg/entity"

	"gorm.io/gorm"
)

type JoinedPositionElasticRepository struct {
	db *gorm.DB
}

func NewJoinedPositionElasticRepository(db *gorm.DB) *JoinedPositionElasticRepository {
	return &JoinedPositionElasticRepository{
		db,
	}
}
func (p *JoinedPositionElasticRepository) SaveBatch(joinedPositions []entity.ElasticJoinedPosition) ([]entity.ElasticJoinedPosition, error) {
	res := p.db.CreateInBatches(joinedPositions, 100)
	if res.Error != nil {
		return nil, res.Error
	}
	return joinedPositions, nil
}

func (p *JoinedPositionElasticRepository) FindAll() ([]entity.ElasticJoinedPosition, error) {
	var joinedPositions []entity.ElasticJoinedPosition
	result := p.db.Find(&joinedPositions)
	if result.Error != nil {
		return nil, result.Error
	}
	return joinedPositions, nil
}
