package entity

import "github.com/jackc/pgtype"

type ElasticFactory struct {
	Id                           string         `gorm:"column:id"`
	Vid                          int            `gorm:"column:vid"`
	Block                        int            `gorm:"column:block"`
	PoolCount                    pgtype.Numeric `gorm:"column:pool_count"`
	TxCount                      pgtype.Numeric `gorm:"column:tx_count"`
	TotalVolumeUsd               float64        `gorm:"column:total_volume_usd"`
	TotalVolumeEth               float64        `gorm:"column:total_volume_eth"`
	TotalFeesUsd                 float64        `gorm:"column:total_fees_usd"`
	TotalFeesEth                 float64        `gorm:"column:total_fees_eth"`
	UntrackedVolumeUsd           float64        `gorm:"column:untracked_volume_usd"`
	TotalValueLockedUsd          float64        `gorm:"column:total_value_locked_usd"`
	TotalValueLockedEth          float64        `gorm:"column:total_value_locked_eth"`
	TotalValueLockedUsdUntracked float64        `gorm:"column:total_value_locked_usd_untracked"`
	TotalValueLockedEthUntracked float64        `gorm:"column:total_value_locked_eth_untracked"`
	Owner                        string         `gorm:"column:owner"`
}
