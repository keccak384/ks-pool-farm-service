package repository

import "ks-data-prepare/internal/pkg/entity"

type IBundleElasticDatastoreRepository interface {
	Find() (entity.ElasticBundle, error)
	Save(bundle entity.ElasticBundle) (entity.ElasticBundle, error)
}
type IFactoryElasticDatastoreRepository interface {
	Find() (entity.ElasticFactory, error)
	Save(factory entity.ElasticFactory) (entity.ElasticFactory, error)
}

type IPoolElasticDatastoreRepository interface {
	FindAll() ([]entity.ElasticPool, error)
	SaveBatch([]entity.ElasticPool) ([]entity.ElasticPool, error)
}

type IPositionElasticDatastoreRepository interface {
	FindAll() ([]entity.ElasticPosition, error)
	SaveBatch([]entity.ElasticPosition) ([]entity.ElasticPosition, error)
}

type ITokenElasticDatastoreRepository interface {
	FindAll() ([]entity.ElasticToken, error)
	FindByAddresses([]string) ([]entity.ElasticToken, error)
	SaveBatch([]entity.ElasticToken) ([]entity.ElasticToken, error)
}

type ITickElasticDatastoreRepository interface {
	FindAll() ([]entity.ElasticTick, error)
	SaveBatch([]entity.ElasticTick) ([]entity.ElasticTick, error)
}

type IFarmElasticDatastoreRepository interface {
	FindAll() ([]entity.ElasticFarm, error)
	SaveBatch([]entity.ElasticFarm) ([]entity.ElasticFarm, error)
}

type IFarmPoolElasticDatastoreRepository interface {
	FindAll() ([]entity.ElasticFarmPool, error)
	SaveBatch([]entity.ElasticFarmPool) ([]entity.ElasticFarmPool, error)
}

type IJoinedPositionElasticDatastoreRepository interface {
	FindAll() ([]entity.ElasticJoinedPosition, error)
	SaveBatch([]entity.ElasticJoinedPosition) ([]entity.ElasticJoinedPosition, error)
}

type IGraphElasticExternalStoreRepository interface {
	FetchElasticBundleFromGraph() (entity.ElasticBundle, error)
	FetchElasticFactoryFromGraph() (entity.ElasticFactory, error)
	FetchElasticPoolsFromGraph() ([]entity.ElasticPool, error)
	FetchElasticPositionsFromGraph() ([]entity.ElasticPosition, error)
	FetchElasticTokensFromGraph() ([]entity.ElasticToken, error)
	FetchElasticTicksFromGraph() ([]entity.ElasticTick, error)
	FetchElasticFarmsFromGraph() ([]entity.ElasticFarm, error)
	FetchElasticFarmPoolsFromGraph() ([]entity.ElasticFarmPool, error)
	FetchElasticJoinedPositionsFromGraph() ([]entity.ElasticJoinedPosition, error)
}
