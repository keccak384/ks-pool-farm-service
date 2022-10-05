package postgres

import (
	"ks-data-prepare/internal/pkg/entity"

	"gorm.io/gorm"
)

type BundleElasticRepository struct {
	db *gorm.DB
}

func NewBundleElasticRepository(db *gorm.DB) *BundleElasticRepository {
	return &BundleElasticRepository{
		db,
	}
}
func (p *BundleElasticRepository) Save(bundle entity.ElasticBundle) (entity.ElasticBundle, error) {
	res := p.db.Create(&bundle)
	if res.Error != nil {
		return entity.ElasticBundle{}, res.Error
	}
	return bundle, nil
}

func (p *BundleElasticRepository) Find() (entity.ElasticBundle, error) {
	var bundle entity.ElasticBundle
	result := p.db.Find(&bundle)
	if result.Error != nil {
		return entity.ElasticBundle{}, result.Error
	}
	return bundle, nil
}
