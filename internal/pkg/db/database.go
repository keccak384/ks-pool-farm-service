package db

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"ks-data-prepare/internal/pkg/config"
	cfg "ks-data-prepare/internal/pkg/config"
)

var db *gorm.DB
var dbGraphElastic *gorm.DB

func InitDB() error {
	if db != nil && dbGraphElastic != nil {
		return nil
	}

	internalDB, err := initDB(cfg.Instance().DB, db)
	if err != nil {
		return err
	}

	db = internalDB

	externalDB, err := initDB(cfg.Instance().Graph, dbGraphElastic)
	if err != nil {
		return err
	}

	dbGraphElastic = externalDB

	return nil
}
func initDB(dbCfg config.DBConfig, db *gorm.DB) (*gorm.DB, error) {
	gormDB, err := gorm.Open(
		postgres.Open(dbCfg.DNS()),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.LogLevel(dbCfg.LogLevel)),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to open connection, err: %v", err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get *sql.db, err: %v", err)
	}

	sqlDB.SetMaxIdleConns(dbCfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbCfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(dbCfg.ConnLifeTime) * time.Second)

	if err = sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database, err: %v", err)
	}

	return gormDB, nil

}
func Instance() *gorm.DB {
	return db
}

func InstanceGraphElastic() *gorm.DB {
	return dbGraphElastic
}
