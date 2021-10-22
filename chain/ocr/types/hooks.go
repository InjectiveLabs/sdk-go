package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type OcrHooks interface {
	AfterSetFeedConfig(ctx sdk.Context, feedConfig *FeedConfig)
	AfterTransmit(ctx sdk.Context, feedId string, answer sdk.Dec, timestamp int64)
	AfterFundFeedRewardPool(ctx sdk.Context, feedId string, newPoolAmount sdk.Coin)
}

var _ OcrHooks = MultiOcrHooks{}

type MultiOcrHooks []OcrHooks

func NewMultiOcrHooks(hooks ...OcrHooks) MultiOcrHooks {
	return hooks
}

func (h MultiOcrHooks) AfterSetFeedConfig(ctx sdk.Context, feedConfig *FeedConfig) {
	for i := range h {
		h[i].AfterSetFeedConfig(ctx, feedConfig)
	}
}

func (h MultiOcrHooks) AfterTransmit(ctx sdk.Context, feedId string, answer sdk.Dec, timestamp int64) {
	for i := range h {
		h[i].AfterTransmit(ctx, feedId, answer, timestamp)
	}
}

func (h MultiOcrHooks) AfterFundFeedRewardPool(ctx sdk.Context, feedId string, newPoolAmount sdk.Coin) {
	for i := range h {
		h[i].AfterFundFeedRewardPool(ctx, feedId, newPoolAmount)
	}
}
