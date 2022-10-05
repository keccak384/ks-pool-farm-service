package migration

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	cfg "ks-data-prepare/internal/pkg/config"
)

type Migration interface {
	MigrateElasticUp(up int) error
	MigrateElasticDown(down int) error
	MigrateClassicUp(up int) error
	MigrateClassicDown(down int) error
	Close() (error, error)
}

type dbMigration struct {
	mElastic *migrate.Migrate
	mClassic *migrate.Migrate
}

func NewMigration(dir string) (Migration, error) {
	db, errr := sql.Open("postgres", cfg.Instance().DB.MigrationSource())

	fmt.Println("===", errr)
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	mElastic, err := migrate.NewWithDatabaseInstance(
		"file://./migrations/postgres/elastic",
		"postgres",
		driver,
	)

	if err != nil {
		return nil, err
	}

	mClassic, err := migrate.NewWithDatabaseInstance(
		"file://./migrations/postgres/classic",
		"postgres",
		driver,
	)

	if err != nil {
		return nil, err
	}

	return &dbMigration{
		mElastic,
		mClassic,
	}, nil

}

func (t *dbMigration) MigrateElasticUp(up int) error {
	if up == 0 {
		return t.mElastic.Up()
	} else {
		return t.mElastic.Steps(up)
	}
}

func (t *dbMigration) MigrateElasticDown(down int) error {
	if down == 0 {
		return t.mElastic.Down()
	} else {
		return t.mElastic.Steps(-down)
	}
}

func (t *dbMigration) MigrateClassicUp(up int) error {
	if up == 0 {
		return t.mClassic.Up()
	} else {
		return t.mClassic.Steps(up)
	}
}

func (t *dbMigration) MigrateClassicDown(down int) error {
	if down == 0 {
		return t.mClassic.Down()
	} else {
		return t.mClassic.Steps(-down)
	}
}

func (t *dbMigration) Close() (error, error) {
	t.mElastic.Close()
	t.mClassic.Close()
	return nil, nil
}
