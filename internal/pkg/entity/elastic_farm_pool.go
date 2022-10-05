package entity

import (
	"github.com/jackc/pgtype"
	"github.com/lib/pq"
)

type ElasticFarmPool struct {
	Id                 string         `gorm:"column:id" json:"iddd"`
	Vid                int            `gorm:"column:vid" json:"vid"`
	Block              int            `gorm:"column:block" json:"block"`
	Pid                pgtype.Numeric `gorm:"column:pid" json:"pid"`
	StartTime          pgtype.Numeric `gorm:"column:start_time" json:"start_time"`
	EndTime            pgtype.Numeric `gorm:"column:end_time" json:"end_time"`
	FeeTarget          pgtype.Numeric `gorm:"column:fee_target" json:"fee_target"`
	VestingDuration    pgtype.Numeric `gorm:"column:vesting_duration" json:"vesting_duration"`
	PoolID             string         `gorm:"column:pool" json:"-"`
	FairLaunch         string         `gorm:"column:fair_launch" json:"fair_launch"`
	RewardTokenIDs     pq.StringArray `gorm:"column:reward_tokens;type:text[]" json:"reward_tokens"`
	TotalRewardAmounts string         `gorm:"column:total_reward_amounts" json:"total_reward_amounts"`
	Pool               ElasticPool    `gorm:"references:Id;foreignKey:PoolID" json:"pool"`
}

type ElasticFarmPoolJSON struct {
	Id                 string             `json:"id"`
	Vid                int                `json:"vid"`
	Block              int                `json:"block"`
	Pid                string             `json:"pid"`
	StartTime          string             `json:"start_time"`
	EndTime            string             `json:"end_time"`
	FeeTarget          string             `json:"fee_target"`
	VestingDuration    string             `json:"vesting_duration"`
	PoolID             string             `json:"-"`
	FairLaunch         string             `json:"fair_launch"`
	RewardTokenIDs     []string           `json:"reward_tokens_ids"`
	TotalRewardAmounts string             `json:"total_reward_amounts"`
	Pool               ElasticPoolJSON    `json:"pool"`
	RewardTokens       []ElasticTokenJSON `json:"reward_tokens"`
}

func (item ElasticFarmPool) ToJSON(rewardTokens []ElasticToken) ElasticFarmPoolJSON {
	rewardTokensJSON := make([]ElasticTokenJSON, 0, len(rewardTokens))
	for _, item := range rewardTokens {
		rewardTokensJSON = append(rewardTokensJSON, item.ToJSON())
	}
	return ElasticFarmPoolJSON{
		Id:                 item.Id,
		Vid:                item.Vid,
		Block:              item.Block,
		Pid:                item.Pid.Int.String(),
		StartTime:          item.StartTime.Int.String(),
		EndTime:            item.EndTime.Int.String(),
		FeeTarget:          item.FeeTarget.Int.String(),
		VestingDuration:    item.VestingDuration.Int.String(),
		FairLaunch:         item.FairLaunch,
		RewardTokenIDs:     item.RewardTokenIDs,
		TotalRewardAmounts: item.TotalRewardAmounts,
		Pool:               item.Pool.ToJSON(),
		RewardTokens:       rewardTokensJSON,
	}
}
