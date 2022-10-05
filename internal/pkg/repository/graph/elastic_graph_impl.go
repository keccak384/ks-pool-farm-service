package graph

import (
	"fmt"
	"ks-data-prepare/internal/pkg/entity"

	"gorm.io/gorm"
)

type GraphElasticRepository struct {
	db *gorm.DB
}

func NewGraphElasticRepository(db *gorm.DB) *GraphElasticRepository {
	return &GraphElasticRepository{
		db,
	}
}

func (p *GraphElasticRepository) GroupedThenRaw(entity string) *gorm.DB {
	return p.db.Raw(fmt.Sprintf(`
		select grouped.vid, grouped.id, lower(t.block_range) as block, t.*
		from 
		(
			select id, max(vid) as vid from %s 
			group by id 
			order by vid
		) as grouped
		JOIN %s as t on grouped.vid=t.vid`, entity, entity))
}

func (p *GraphElasticRepository) FetchElasticBundleFromGraph() (entity.ElasticBundle, error) {
	var bundle entity.ElasticBundle
	res := p.GroupedThenRaw("sgd1.bundle").Scan(&bundle)
	return bundle, res.Error
}

func (p *GraphElasticRepository) FetchElasticFactoryFromGraph() (entity.ElasticFactory, error) {
	var factory entity.ElasticFactory
	res := p.GroupedThenRaw("sgd1.factory").Scan(&factory)
	return factory, res.Error
}

func (p *GraphElasticRepository) FetchElasticPoolsFromGraph() ([]entity.ElasticPool, error) {
	var pools []entity.ElasticPool
	res := p.GroupedThenRaw("sgd1.pool").Scan(&pools)
	return pools, res.Error
}

func (p *GraphElasticRepository) FetchElasticPositionsFromGraph() ([]entity.ElasticPosition, error) {
	var positions []entity.ElasticPosition
	res := p.GroupedThenRaw("sgd1.position").Scan(&positions)
	return positions, res.Error
}

func (p *GraphElasticRepository) FetchElasticTokensFromGraph() ([]entity.ElasticToken, error) {
	var tokens []entity.ElasticToken
	res := p.GroupedThenRaw("sgd1.token").Scan(&tokens)
	return tokens, res.Error
}

func (p *GraphElasticRepository) FetchElasticTicksFromGraph() ([]entity.ElasticTick, error) {
	var ticks []entity.ElasticTick
	res := p.GroupedThenRaw("sgd1.tick").Scan(&ticks)
	return ticks, res.Error
}

func (p *GraphElasticRepository) FetchElasticFarmsFromGraph() ([]entity.ElasticFarm, error) {
	var farms []entity.ElasticFarm
	res := p.GroupedThenRaw("sgd1.kyber_fair_launch").Scan(&farms)
	return farms, res.Error
}

func (p *GraphElasticRepository) FetchElasticFarmPoolsFromGraph() ([]entity.ElasticFarmPool, error) {
	var farmPools []entity.ElasticFarmPool
	res := p.GroupedThenRaw("sgd1.farming_pool").Scan(&farmPools)
	return farmPools, res.Error
}

func (p *GraphElasticRepository) FetchElasticJoinedPositionsFromGraph() ([]entity.ElasticJoinedPosition, error) {
	var joinedPositions []entity.ElasticJoinedPosition
	res := p.GroupedThenRaw("sgd1.joined_position").Scan(&joinedPositions)
	return joinedPositions, res.Error
}
