package entity

type ElasticFarm struct {
	Id    string `gorm:"column:id"`
	Vid   int    `gorm:"column:vid"`
	Block int    `gorm:"column:block"`
}
