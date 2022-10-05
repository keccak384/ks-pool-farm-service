package entity

type ElasticBundle struct {
	Id              string  `gorm:"column:id"`
	Vid             int     `gorm:"column:vid"`
	Block           int     `gorm:"column:block"`
	ElasticPriceUsd float64 `gorm:"column:eth_price_usd"`
}
