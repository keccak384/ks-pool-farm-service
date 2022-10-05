package entity

import "github.com/jackc/pgtype"

type ElasticPosition struct {
	Id                  string         `gorm:"column:id"`
	Vid                 int            `gorm:"column:vid"`
	Block               int            `gorm:"column:block"`
	Owner               []byte         `gorm:column:owner`
	Pool                string         `gorm:"column:pool"`
	Token0              string         `gorm:"column:token_0"`
	Token1              string         `gorm:"column:token_1"`
	TickLower           string         `gorm:"column:tick_lower"`
	TickUpper           string         `gorm:"column:tick_upper"`
	Liquidity           pgtype.Numeric `gorm:"column:liquidity"`
	Transaction         string         `gorm:"column:transaction"`
	FeeGrowthInsideLast pgtype.Numeric `gorm:"column:fee_growth_inside_last"`
}
