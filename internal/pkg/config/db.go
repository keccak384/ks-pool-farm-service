package config

import "fmt"

type DBConfig struct {
	Type         string `default:"postgres"`
	Host         string `default:"localhost"`
	Port         int    `default:"5432"`
	DBName       string `default:"ks-data-prepare"`
	User         string `default:"user"`
	Password     string `default:"password"`
	ConnLifeTime int    `default:"300"`
	ConnTimeOut  int    `default:"30"`
	MaxIdleConns int    `default:"10"`
	MaxOpenConns int    `default:"80"`
	LogLevel     int    `default:"1"`
}

func (c *DBConfig) DNS() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		c.User, c.Password, c.Host, c.Port, c.DBName,
	)
}

func (c *DBConfig) MigrationSource() string {
	fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", c.User, c.Password, c.Host, c.Port, c.DBName)
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", c.User, c.Password, c.Host, c.Port, c.DBName)
}
