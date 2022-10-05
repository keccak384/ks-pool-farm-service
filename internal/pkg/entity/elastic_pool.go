package entity

import "github.com/jackc/pgtype"

type ElasticPool struct {
	Id                                    string         `gorm:"column:id"`
	Vid                                   int            `gorm:"column:vid"`
	Block                                 int            `gorm:"column:block"`
	CreatedAtTimestamp                    pgtype.Numeric `gorm:"column:created_at_timestamp"`
	CreatedAtBlockNumber                  pgtype.Numeric `gorm:"column:created_at_block_number"`
	Token0ID                              string         `gorm:"column:token_0"`
	Token1ID                              string         `gorm:"column:token_1"`
	Token0                                ElasticToken   `gorm:"references:Id;foreignKey:Token0ID"`
	Token1                                ElasticToken   `gorm:"references:Id;foreignKey:Token1ID"`
	FeeTier                               pgtype.Numeric `gorm:"column:fee_tier"`
	Liquidity                             pgtype.Numeric `gorm:"column:liquidity"`
	ReinvestL                             pgtype.Numeric `gorm:"column:reinvest_l"`
	ReinvestLLast                         pgtype.Numeric `gorm:"column:reinvest_l_last"`
	SqrtPrice                             pgtype.Numeric `gorm:"column:sqrt_price"`
	TotalSupply                           pgtype.Numeric `gorm:"column:total_supply"`
	FeeGrowthGlobal                       pgtype.Numeric `gorm:"column:fee_growth_global"`
	SecondsPerLiquidityGlobal             pgtype.Numeric `gorm:"column:seconds_per_liquidity_global"`
	LastSecondsPerLiquidityDataUpdateTime pgtype.Numeric `gorm:"column:last_seconds_per_liquidity_data_update_time"`
	Token0Price                           float64        `gorm:"column:token_0_price"`
	Token1Price                           float64        `gorm:"column:token_1_price"`
	Tick                                  pgtype.Numeric `gorm:"column:tick"`
	VolumeToken0                          float64        `gorm:"column:volume_token_0"`
	VolumeToken1                          float64        `gorm:"column:volume_token_1"`
	VolumeUsd                             float64        `gorm:"column:volume_usd"`
	UntrackedVolumeUsd                    float64        `gorm:"column:untracked_volume_usd"`
	FeesUsd                               float64        `gorm:"column:fees_usd"`
	TxCount                               pgtype.Numeric `gorm:"column:tx_count"`
	CollectedFeesToken0                   float64        `gorm:"column:collected_fees_token_0"`
	CollectedFeesToken1                   float64        `gorm:"column:collected_fees_token_1"`
	CollectedFeesUsd                      float64        `gorm:"column:collected_fees_usd"`
	TotalValueLockedToken0                float64        `gorm:"column:total_value_locked_token_0"`
	TotalValueLockedToken1                float64        `gorm:"column:total_value_locked_token_1"`
	TotalValueLockedEth                   float64        `gorm:"column:total_value_locked_eth"`
	TotalValueLockedUsd                   float64        `gorm:"column:total_value_locked_usd"`
	TotalValueLockedUsdUntracked          float64        `gorm:"column:total_value_locked_usd_untracked"`
	LiquidityProviderCount                pgtype.Numeric `gorm:"column:liquidity_provider_count"`
	PositionCount                         pgtype.Numeric `gorm:"column:position_count"`
	ClosedPostionCount                    pgtype.Numeric `gorm:"column:closed_postion_count"`
}

type ElasticPoolJSON struct {
	Id                                    string           `json:"id"`
	Vid                                   int              `json:"vid"`
	Block                                 int              `json:"block"`
	CreatedAtTimestamp                    string           `json:"created_at_timestamp"`
	CreatedAtBlockNumber                  string           `json:"created_at_block_number"`
	Token0                                ElasticTokenJSON `json:"token_0"`
	Token1                                ElasticTokenJSON `json:"token_1"`
	FeeTier                               string           `json:"fee_tier"`
	Liquidity                             string           `json:"liquidity"`
	ReinvestL                             string           `json:"reinvest_l"`
	ReinvestLLast                         string           `json:"reinvest_l_last"`
	SqrtPrice                             string           `json:"sqrt_price"`
	TotalSupply                           string           `json:"total_supply"`
	FeeGrowthGlobal                       string           `json:"fee_growth_global"`
	SecondsPerLiquidityGlobal             string           `json:"seconds_per_liquidity_global"`
	LastSecondsPerLiquidityDataUpdateTime string           `json:"last_seconds_per_liquidity_data_update_time"`
	Token0Price                           float64          `json:"token_0_price"`
	Token1Price                           float64          `json:"token_1_price"`
	Tick                                  string           `json:"tick"`
	VolumeToken0                          float64          `json:"volume_token_0"`
	VolumeToken1                          float64          `json:"volume_token_1"`
	VolumeUsd                             float64          `json:"volume_usd"`
	UntrackedVolumeUsd                    float64          `json:"untracked_volume_usd"`
	FeesUsd                               float64          `json:"fees_usd"`
	TxCount                               string           `json:"tx_count"`
	CollectedFeesToken0                   float64          `json:"collected_fees_token_0"`
	CollectedFeesToken1                   float64          `json:"collected_fees_token_1"`
	CollectedFeesUsd                      float64          `json:"collected_fees_usd"`
	TotalValueLockedToken0                float64          `json:"total_value_locked_token_0"`
	TotalValueLockedToken1                float64          `json:"total_value_locked_token_1"`
	TotalValueLockedEth                   float64          `json:"total_value_locked_eth"`
	TotalValueLockedUsd                   float64          `json:"total_value_locked_usd"`
	TotalValueLockedUsdUntracked          float64          `json:"total_value_locked_usd_untracked"`
	LiquidityProviderCount                string           `json:"liquidity_provider_count"`
	PositionCount                         string           `json:"position_count"`
	ClosedPostionCount                    string           `json:"closed_postion_count"`
}

func (item ElasticPool) ToJSON() ElasticPoolJSON {
	return ElasticPoolJSON{
		Id:                                    item.Id,
		Vid:                                   item.Vid,
		Block:                                 item.Block,
		CreatedAtTimestamp:                    item.CreatedAtTimestamp.Int.String(),
		CreatedAtBlockNumber:                  item.CreatedAtBlockNumber.Int.String(),
		Token0:                                item.Token0.ToJSON(),
		Token1:                                item.Token1.ToJSON(),
		FeeTier:                               item.FeeTier.Int.String(),
		Liquidity:                             item.Liquidity.Int.String(),
		ReinvestL:                             item.ReinvestL.Int.String(),
		ReinvestLLast:                         item.ReinvestLLast.Int.String(),
		SqrtPrice:                             item.SqrtPrice.Int.String(),
		TotalSupply:                           item.TotalSupply.Int.String(),
		FeeGrowthGlobal:                       item.FeeGrowthGlobal.Int.String(),
		SecondsPerLiquidityGlobal:             item.SecondsPerLiquidityGlobal.Int.String(),
		LastSecondsPerLiquidityDataUpdateTime: item.LastSecondsPerLiquidityDataUpdateTime.Int.String(),
		Token0Price:                           item.Token0Price,
		Token1Price:                           item.Token1Price,
		Tick:                                  item.Tick.Int.String(),
		VolumeToken0:                          item.VolumeToken0,
		VolumeToken1:                          item.VolumeToken1,
		VolumeUsd:                             item.VolumeUsd,
		UntrackedVolumeUsd:                    item.UntrackedVolumeUsd,
		FeesUsd:                               item.FeesUsd,
		TxCount:                               item.TxCount.Int.String(),
		CollectedFeesToken0:                   item.CollectedFeesToken0,
		CollectedFeesToken1:                   item.CollectedFeesToken1,
		CollectedFeesUsd:                      item.CollectedFeesUsd,
		TotalValueLockedToken0:                item.TotalValueLockedToken0,
		TotalValueLockedToken1:                item.TotalValueLockedToken1,
		TotalValueLockedEth:                   item.TotalValueLockedEth,
		TotalValueLockedUsd:                   item.TotalValueLockedUsd,
		TotalValueLockedUsdUntracked:          item.TotalValueLockedUsdUntracked,
		LiquidityProviderCount:                item.LiquidityProviderCount.Int.String(),
		PositionCount:                         item.PositionCount.Int.String(),
		ClosedPostionCount:                    item.ClosedPostionCount.Int.String(),
	}
}
