package builder

import (
	"ks-data-prepare/internal/pkg/db"
	"ks-data-prepare/internal/pkg/repository"
	"ks-data-prepare/internal/pkg/repository/graph"
	"ks-data-prepare/internal/pkg/repository/postgres"
)

type migrationServer struct {
	graphRepo          repository.IGraphElasticExternalStoreRepository
	bundleRepo         repository.IBundleElasticDatastoreRepository
	factoryRepo        repository.IFactoryElasticDatastoreRepository
	poolRepo           repository.IPoolElasticDatastoreRepository
	positionRepo       repository.IPositionElasticDatastoreRepository
	tokenRepo          repository.ITokenElasticDatastoreRepository
	tickRepo           repository.ITickElasticDatastoreRepository
	farmRepo           repository.IFarmElasticDatastoreRepository
	farmPoolRepo       repository.IFarmPoolElasticDatastoreRepository
	joinedPositionRepo repository.IJoinedPositionElasticDatastoreRepository
}

func NewElasticMigration() IRunner {
	graphElastic := graph.NewGraphElasticRepository(db.InstanceGraphElastic())
	bundleElasticDataStore := postgres.NewBundleElasticRepository(db.Instance())
	factoryElasticDataStore := postgres.NewFactoryElasticRepository(db.Instance())
	poolElasticDataStore := postgres.NewPoolElasticRepository(db.Instance())
	positionElasticDataStore := postgres.NewPositionElasticRepository(db.Instance())

	tokenElasticDataStore := postgres.NewTokenElasticRepository(db.Instance())
	tickElasticDataStore := postgres.NewTickElasticRepository(db.Instance())
	farmElasticDataStore := postgres.NewFarmElasticRepository(db.Instance())
	farmPoolElasticDataStore := postgres.NewFarmPoolElasticRepository(db.Instance())
	joinedPositionElasticDataStore := postgres.NewJoinedPositionElasticRepository(db.Instance())

	return &migrationServer{
		graphElastic,
		bundleElasticDataStore,
		factoryElasticDataStore,
		poolElasticDataStore,
		positionElasticDataStore,
		tokenElasticDataStore,
		tickElasticDataStore,
		farmElasticDataStore,
		farmPoolElasticDataStore,
		joinedPositionElasticDataStore,
	}
}

func (f *migrationServer) Run() error {
	return f.bootstrapElastic()
}

func (f *migrationServer) bootstrapElastic() error {
	// f.bootstrapPools()
	bootstraps := []func() error{
		f.bootstrapBundle,
		f.bootstrapFactory,
		f.bootstrapTokens,
		f.bootstrapTicks,
		f.bootstrapPools,
		f.bootstrapPositions,
		f.bootstrapFarms,
		f.bootstrapFarmPools,
		f.bootstrapJoinedPositions,
	}

	for _, f := range bootstraps {
		// go func(f func() error) {
		f()
		// }(f)
	}

	for {

	}
}

func (f *migrationServer) bootstrapBundle() error {
	bundle, err := f.graphRepo.FetchElasticBundleFromGraph()
	if err != nil {
		return err
	}
	_, err = f.bundleRepo.Save(bundle)
	if err != nil {
		return err
	}
	return nil
}

func (f *migrationServer) bootstrapFactory() error {

	factory, err := f.graphRepo.FetchElasticFactoryFromGraph()
	if err != nil {
		return err
	}
	_, err = f.factoryRepo.Save(factory)
	if err != nil {
		return err
	}
	return nil
}

func (f *migrationServer) bootstrapPools() error {
	pools, err := f.graphRepo.FetchElasticPoolsFromGraph()
	if err != nil {
		return err
	}
	_, err = f.poolRepo.SaveBatch(pools)
	if err != nil {
		return err
	}
	return nil
}

func (f *migrationServer) bootstrapPositions() error {
	positions, err := f.graphRepo.FetchElasticPositionsFromGraph()
	if err != nil {
		return err
	}
	_, err = f.positionRepo.SaveBatch(positions)
	if err != nil {
		return err
	}
	return nil
}

func (f *migrationServer) bootstrapTokens() error {
	tokens, err := f.graphRepo.FetchElasticTokensFromGraph()

	if err != nil {

		return err
	}

	_, err = f.tokenRepo.SaveBatch(tokens)
	if err != nil {
		return err
	}
	return nil
}

func (f *migrationServer) bootstrapTicks() error {
	ticks, err := f.graphRepo.FetchElasticTicksFromGraph()
	if err != nil {
		return err
	}
	_, err = f.tickRepo.SaveBatch(ticks)
	if err != nil {
		return err
	}
	return nil
}

func (f *migrationServer) bootstrapFarms() error {
	farms, err := f.graphRepo.FetchElasticFarmsFromGraph()
	if err != nil {
		return err
	}
	_, err = f.farmRepo.SaveBatch(farms)
	if err != nil {
		return err
	}
	return nil
}

func (f *migrationServer) bootstrapFarmPools() error {
	farmPools, err := f.graphRepo.FetchElasticFarmPoolsFromGraph()
	if err != nil {
		return err
	}
	_, err = f.farmPoolRepo.SaveBatch(farmPools)
	if err != nil {
		return err
	}
	return nil
}

func (f *migrationServer) bootstrapJoinedPositions() error {
	joinedPositions, err := f.graphRepo.FetchElasticJoinedPositionsFromGraph()
	if err != nil {
		return err
	}
	_, err = f.joinedPositionRepo.SaveBatch(joinedPositions)
	if err != nil {
		return err
	}
	return nil
}
