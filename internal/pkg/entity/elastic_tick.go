package entity

import "github.com/jackc/pgtype"

type ElasticTick struct {
	Id                         string         `gorm:"column:id"`
	Vid                        int            `gorm:"column:vid"`
	Block                      int            `gorm:"column:block"`
	PoolAddress                string         `gorm:"column:pool_address"`
	TickIdx                    pgtype.Numeric `gorm:"column:tick_idx"`
	Pool                       string         `gorm:"column:pool"`
	LiquidityGross             pgtype.Numeric `gorm:"column:liquidity_gross"`
	LiquidityNet               pgtype.Numeric `gorm:"column:liquidity_net"`
	Price0                     float64        `gorm:"column:price_0"`
	Price1                     float64        `gorm:"column:price_1"`
	VolumeToken0               float64        `gorm:"column:volume_token_0"`
	VolumeToken1               float64        `gorm:"column:volume_token_1"`
	VolumeUsd                  float64        `gorm:"column:volume_usd"`
	UntrackedVolumeUsd         float64        `gorm:"column:untracked_volume_usd"`
	FeesUsd                    float64        `gorm:"column:fees_usd"`
	CollectedFeesToken0        float64        `gorm:"column:collected_fees_token_0"`
	CollectedFeesToken1        float64        `gorm:"column:collected_fees_token_1"`
	CollectedFeesUsd           float64        `gorm:"column:collected_fees_usd"`
	CreatedAtTimestamp         pgtype.Numeric `gorm:"column:created_at_timestamp"`
	CreatedAtBlockNumber       pgtype.Numeric `gorm:"column:created_at_block_number"`
	LiquidityProviderCount     pgtype.Numeric `gorm:"column:liquidity_provider_count"`
	FeeGrowthOutside           pgtype.Numeric `gorm:"column:fee_growth_outside"`
	SecondsPerLiquidityOutside pgtype.Numeric `gorm:"column:seconds_per_liquidity_outside"`
}
