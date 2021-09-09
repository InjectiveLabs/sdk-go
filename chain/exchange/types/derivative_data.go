package types

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

type PositionState struct {
	Position       *Position
	FundingPayment sdk.Dec
}

type DerivativeOrderStateExpansion struct {
	SubaccountID  common.Hash
	PositionDelta *PositionDelta
	Payout        sdk.Dec

	TotalBalanceDelta     sdk.Dec
	AvailableBalanceDelta sdk.Dec

	AuctionFeeReward            sdk.Dec
	LiquidityMiningRewardPoints sdk.Dec
	FeeRecipientReward          sdk.Dec
	FeeRecipient                common.Address
	LimitOrderFilledDelta       *DerivativeLimitOrderDelta
	MarketOrderFilledDelta      *DerivativeMarketOrderDelta
	OrderHash                   common.Hash
}

func ApplyDeltasAndGetDerivativeOrderBatchEvent(
	isBuy bool,
	executionType ExecutionType,
	market *DerivativeMarket,
	funding *PerpetualMarketFunding,
	stateExpansions []*DerivativeOrderStateExpansion,
	depositDeltas DepositDeltas,
	liquidityMiningRewards LiquidityMiningRewards,
) (batch *EventBatchDerivativeExecution, filledDeltas []*DerivativeLimitOrderDelta) {
	if len(stateExpansions) == 0 {
		return
	}

	trades := make([]*DerivativeTradeLog, 0, len(stateExpansions))

	if !executionType.IsMarket() {
		filledDeltas = make([]*DerivativeLimitOrderDelta, 0, len(stateExpansions))
	}

	for idx := range stateExpansions {
		expansion := stateExpansions[idx]

		feeRecipientSubaccount := EthAddressToSubaccountID(expansion.FeeRecipient)
		if bytes.Equal(feeRecipientSubaccount.Bytes(), common.Hash{}.Bytes()) {
			feeRecipientSubaccount = AuctionSubaccountID
		}

		depositDeltas.ApplyDepositDelta(expansion.SubaccountID, &DepositDelta{
			TotalBalanceDelta:     expansion.TotalBalanceDelta,
			AvailableBalanceDelta: expansion.AvailableBalanceDelta,
		})
		depositDeltas.ApplyUniformDelta(feeRecipientSubaccount, expansion.FeeRecipientReward)
		depositDeltas.ApplyUniformDelta(AuctionSubaccountID, expansion.AuctionFeeReward)

		sender := SubaccountIDToSdkAddress(expansion.SubaccountID)
		if expansion.LiquidityMiningRewardPoints.IsPositive() {
			liquidityMiningRewards.AddPointsForAddress(sender.String(), expansion.LiquidityMiningRewardPoints)
		}

		if !executionType.IsMarket() {
			filledDeltas = append(filledDeltas, expansion.LimitOrderFilledDelta)
		}

		if expansion.PositionDelta != nil {
			tradeLog := &DerivativeTradeLog{
				SubaccountId:  expansion.SubaccountID.Bytes(),
				PositionDelta: expansion.PositionDelta,
				Payout:        expansion.Payout,
				Fee:           expansion.FeeRecipientReward.Add(expansion.AuctionFeeReward),
				OrderHash:     expansion.OrderHash.Bytes(),
			}
			trades = append(trades, tradeLog)
		}
	}

	if len(trades) == 0 {
		return nil, filledDeltas
	}

	batch = &EventBatchDerivativeExecution{
		MarketId:      market.MarketId,
		IsBuy:         isBuy,
		IsLiquidation: false,
		ExecutionType: executionType,
		Trades:        trades,
	}
	if funding != nil {
		batch.CumulativeFunding = &funding.CumulativeFunding
	}
	return batch, filledDeltas
}
