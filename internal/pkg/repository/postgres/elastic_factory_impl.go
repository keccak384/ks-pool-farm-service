package postgres

import (
	"ks-data-prepare/internal/pkg/entity"

	"gorm.io/gorm"
)

type FactoryElasticRepository struct {
	db *gorm.DB
}

func NewFactoryElasticRepository(db *gorm.DB) *FactoryElasticRepository {
	return &FactoryElasticRepository{
		db,
	}
}
func (p *FactoryElasticRepository) Save(factory entity.ElasticFactory) (entity.ElasticFactory, error) {
	res := p.db.Create(&factory)
	if res.Error != nil {
		return entity.ElasticFactory{}, res.Error
	}
	return factory, nil
}

func (p *FactoryElasticRepository) Find() (entity.ElasticFactory, error) {
	var factory entity.ElasticFactory
	result := p.db.Find(&factory)
	if result.Error != nil {
		return entity.ElasticFactory{}, result.Error
	}
	return factory, nil
}
