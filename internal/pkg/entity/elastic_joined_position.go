package entity

import "github.com/jackc/pgtype"

type ElasticJoinedPosition struct {
	Id         string         `gorm:"column:id"`
	Vid        int            `gorm:"column:vid"`
	Block      int            `gorm:"column:block"`
	Pid        pgtype.Numeric `gorm:"column:pid"`
	NftId      pgtype.Numeric `gorm:"column:nft_id"`
	Liquidity  pgtype.Numeric `gorm:"column:liquidity"`
	Pool       string         `gorm:"column:pool"`
	Position   string         `gorm:"column:position"`
	FairLaunch string         `gorm:"column:fair_launch"`
}
