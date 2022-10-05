package entity

import (
	"github.com/jackc/pgtype"
	"github.com/lib/pq"
)

type ElasticToken struct {
	Id                           string         `gorm:"column:id"`
	Vid                          int            `gorm:"column:vid"`
	Block                        int            `gorm:"column:block"`
	Symbol                       string         `gorm:"column:symbol"`
	Name                         string         `gorm:"column:name"`
	Decimals                     pgtype.Numeric `gorm:"column:decimals"`
	TotalSupply                  pgtype.Numeric `gorm:"column:total_supply"`
	Volume                       float64        `gorm:"column:volume"`
	VolumeUsd                    float64        `gorm:"column:volume_usd"`
	UntrackedVolumeUsd           float64        `gorm:"column:untracked_volume_usd"`
	FeesUsd                      float64        `gorm:"column:fees_usd"`
	TxCount                      pgtype.Numeric `gorm:"column:tx_count"`
	PoolCount                    pgtype.Numeric `gorm:"column:pool_count"`
	TotalValueLocked             float64        `gorm:"column:total_value_locked"`
	TotalValueLockedUsd          float64        `gorm:"column:total_value_locked_usd"`
	TotalValueLockedUsdUntracked float64        `gorm:"column:total_value_locked_usd_untracked"`
	DerivedEth                   float64        `gorm:"column:derived_eth"`
	WhitelistPools               pq.StringArray `gorm:"column:whitelist_pools;type:text[]"`
}

type ElasticTokenJSON struct {
	Id                           string   `json:"id"`
	Vid                          int      `json:"vid"`
	Block                        int      `json:"block"`
	Symbol                       string   `json:"symbol"`
	Name                         string   `json:"name"`
	Decimals                     string   `json:"decimals"`
	TotalSupply                  string   `json:"total_supply"`
	Volume                       float64  `json:"volume"`
	VolumeUsd                    float64  `json:"volume_usd"`
	UntrackedVolumeUsd           float64  `json:"untracked_volume_usd"`
	FeesUsd                      float64  `json:"fees_usd"`
	TxCount                      string   `json:"tx_count"`
	PoolCount                    string   `json:"pool_count"`
	TotalValueLocked             float64  `json:"total_value_locked"`
	TotalValueLockedUsd          float64  `json:"total_value_locked_usd"`
	TotalValueLockedUsdUntracked float64  `json:"total_value_locked_usd_untracked"`
	DerivedEth                   float64  `json:"derived_eth"`
	WhitelistPools               []string `json:"whitelist_pools"`
}

func (item ElasticToken) ToJSON() ElasticTokenJSON {
	return ElasticTokenJSON{
		Id:                  item.Id,
		Vid:                 item.Vid,
		Block:               item.Block,
		Name:                item.Name,
		Decimals:            item.Decimals.Int.String(),
		TotalSupply:         item.TotalSupply.Int.String(),
		Volume:              item.Volume,
		VolumeUsd:           item.VolumeUsd,
		UntrackedVolumeUsd:  item.UntrackedVolumeUsd,
		FeesUsd:             item.FeesUsd,
		TxCount:             item.TxCount.Int.String(),
		PoolCount:           item.PoolCount.Int.String(),
		TotalValueLocked:    item.TotalValueLocked,
		TotalValueLockedUsd: item.TotalValueLockedUsd,
		DerivedEth:          item.DerivedEth,
		WhitelistPools:      item.WhitelistPools,
	}
}
