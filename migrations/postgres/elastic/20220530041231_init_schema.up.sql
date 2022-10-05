CREATE TABLE "elastic_bundles" (
  "id" text PRIMARY KEY,
  "vid" int8 NOT NULL,
  "block" int4 NOT NULL,
  "eth_price_usd" numeric  NOT NULL
);

CREATE TABLE "elastic_factories" (
  "id" text PRIMARY KEY,
  "vid" int8 NOT NULL,
  "block" int4 NOT NULL,
  "pool_count" numeric  NOT NULL,
  "tx_count" numeric  NOT NULL,
  "total_volume_usd" numeric  NOT NULL,
  "total_volume_eth" numeric  NOT NULL,
  "total_fees_usd" numeric  NOT NULL,
  "total_fees_eth" numeric  NOT NULL,
  "untracked_volume_usd" numeric  NOT NULL,
  "total_value_locked_usd" numeric  NOT NULL,
  "total_value_locked_eth" numeric  NOT NULL,
  "total_value_locked_usd_untracked" numeric  NOT NULL,
  "total_value_locked_eth_untracked" numeric  NOT NULL,
  "owner" text  NOT NULL
);

CREATE TABLE "elastic_pools" (
  "id" text PRIMARY KEY,
  "vid" int8 NOT NULL,
  "block" int4 NOT NULL,
  "created_at_timestamp" numeric NOT NULL,
  "created_at_block_number" numeric NOT NULL,
  "token_0" text NOT NULL,
  "token_1" text NOT NULL,
  "fee_tier" numeric NOT NULL,
  "liquidity" numeric NOT NULL,
  "reinvest_l" numeric NOT NULL,
  "reinvest_l_last" numeric NOT NULL,
  "sqrt_price" numeric NOT NULL,
  "total_supply" numeric NOT NULL,
  "fee_growth_global" numeric NOT NULL,
  "seconds_per_liquidity_global" numeric NOT NULL,
  "last_seconds_per_liquidity_data_update_time" numeric NOT NULL,
  "token_0_price" numeric NOT NULL,
  "token_1_price" numeric NOT NULL,
  "tick" numeric,
  "volume_token_0" numeric NOT NULL,
  "volume_token_1" numeric NOT NULL,
  "volume_usd" numeric NOT NULL,
  "untracked_volume_usd" numeric NOT NULL,
  "fees_usd" numeric  NOT NULL,
  "tx_count" numeric  NOT NULL,
  "collected_fees_token_0" numeric  NOT NULL,
  "collected_fees_token_1" numeric  NOT NULL,
  "collected_fees_usd" numeric  NOT NULL,
  "total_value_locked_token_0" numeric  NOT NULL,
  "total_value_locked_token_1" numeric  NOT NULL,
  "total_value_locked_eth" numeric  NOT NULL,
  "total_value_locked_usd" numeric  NOT NULL,
  "total_value_locked_usd_untracked" numeric  NOT NULL,
  "liquidity_provider_count" numeric  NOT NULL,      
  "position_count" numeric  NOT NULL,      
  "closed_postion_count" numeric  NOT NULL       
);

CREATE TABLE "elastic_positions" (
  "id" text PRIMARY KEY,
  "vid" int8 NOT NULL,
  "block" int4 NOT NULL,
  "owner" bytea NOT NULL,
  "pool" text NOT NULL,
  "token_0" text NOT NULL,
  "token_1" text NOT NULL,
  "tick_lower" text NOT NULL,
  "tick_upper" text NOT NULL,
  "liquidity" numeric NOT NULL,
  "transaction" text NOT NULL,
  "fee_growth_inside_last" numeric NOT NULL
);

CREATE TABLE "elastic_tokens" (
  "id" text PRIMARY KEY,
  "vid" int8 NOT NULL,
  "block" int4 NOT NULL,
  "symbol" text NOT NULL,
  "name" text NOT NULL,
  "decimals" numeric NOT NULL,
  "total_supply" numeric NOT NULL,
  "volume" numeric NOT NULL,
  "volume_usd" numeric NOT NULL,
  "untracked_volume_usd" numeric NOT NULL,
  "fees_usd" numeric NOT NULL,
  "tx_count" numeric NOT NULL,
  "pool_count" numeric NOT NULL,
  "total_value_locked" numeric NOT NULL,
  "total_value_locked_usd" numeric NOT NULL,
  "total_value_locked_usd_untracked" numeric NOT NULL,
  "derived_eth" numeric NOT NULL,
  "whitelist_pools" _text NOT NULL
);

CREATE TABLE "elastic_ticks" (
  "id" text PRIMARY KEY,
  "vid" int8 NOT NULL,
  "block" int4 NOT NULL,
  "pool_address" text NOT NULL,
  "tick_idx" numeric NOT NULL,
  "pool" text NOT NULL,
  "liquidity_gross" numeric NOT NULL,
  "liquidity_net" numeric NOT NULL,
  "price_0" numeric NOT NULL,
  "price_1" numeric NOT NULL,
  "volume_token_0" numeric NOT NULL,
  "volume_token_1" numeric NOT NULL,
  "volume_usd" numeric NOT NULL,
  "untracked_volume_usd" numeric NOT NULL,
  "fees_usd" numeric NOT NULL,
  "collected_fees_token_0" numeric NOT NULL,
  "collected_fees_token_1" numeric NOT NULL,
  "collected_fees_usd" numeric NOT NULL,
  "created_at_timestamp" numeric NOT NULL,
  "created_at_block_number" numeric NOT NULL,
  "liquidity_provider_count" numeric NOT NULL,
  "fee_growth_outside" numeric NOT NULL,
  "seconds_per_liquidity_outside" numeric NOT NULL
);


CREATE TABLE "elastic_farms" (
  "id" text PRIMARY KEY,
  "vid" int8 NOT NULL,
  "block" int4 NOT NULL
);

CREATE TABLE "elastic_farm_pools" (
  "id" text PRIMARY KEY,
  "vid" int8 NOT NULL,
  "block" int4 NOT NULL,
  "pid" numeric NOT NULL,
  "start_time" numeric NOT NULL,
  "end_time" numeric NOT NULL,
  "fee_target" numeric NOT NULL,
  "vesting_duration" numeric NOT NULL,
  "pool" text NOT NULL,
  "fair_launch" text NOT NULL,
  "reward_tokens" _text NOT NULL,
  "total_reward_amounts" text NOT NULL
);

CREATE TABLE "elastic_joined_positions" (
  "id" text PRIMARY KEY,
  "vid" int8 NOT NULL,
  "block" int4 NOT NULL,
  "pid" numeric NOT NULL,
  "nft_id" numeric NOT NULL,
  "liquidity" numeric NOT NULL,
  "pool" text NOT NULL,
  "position" text NOT NULL,
  "fair_launch" text NOT NULL
);

ALTER TABLE "elastic_pools" ADD FOREIGN KEY ("token_0") REFERENCES "elastic_tokens" ("id") ON DELETE CASCADE;
ALTER TABLE "elastic_pools" ADD FOREIGN KEY ("token_1") REFERENCES "elastic_tokens" ("id") ON DELETE CASCADE;
ALTER TABLE "elastic_farm_pools" ADD FOREIGN KEY ("pool") REFERENCES "elastic_pools" ("id") ON DELETE CASCADE;