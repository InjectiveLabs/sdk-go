package types_test

import (
	"fmt"
	oracletypes "github.com/InjectiveLabs/sdk-go/chain/oracle/types"

	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	. "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Msgs", func() {

	Describe("Validate Deposit", func() {
		var msg MsgDeposit
		baseDenom := "INJ"

		BeforeEach(func() {
			amountToDeposit := sdk.NewCoin(baseDenom, sdk.NewInt(50))
			subaccountID := "0x90f8bf6a479f320ead074411a4b0e7944ea8c9c1000000000000000000000001"
			sender := sdk.AccAddress(common.FromHex("90f8bf6a479f320ead074411a4b0e7944ea8c9c1"))

			msg = MsgDeposit{
				Sender:       sender.String(),
				SubaccountId: subaccountID,
				Amount:       amountToDeposit,
			}
		})

		Context("With empty subaccountID", func() {
			It("should be valid", func() {
				msg.SubaccountId = ""
				err := msg.ValidateBasic()

				Ω(err).To(BeNil())
			})
		})

		Context("With all valid fields", func() {
			It("should be valid", func() {
				err := msg.ValidateBasic()

				Ω(err).To(BeNil())
			})
		})

		Context("With bad subaccountID", func() {
			It("should be invalid with bad subaccountID error", func() {
				badSubaccountId := "0x90f8bf6a47f320ead074411a4b0e7944ea8c9c1000000000000000000000001" // one less character
				msg.SubaccountId = badSubaccountId
				err := msg.ValidateBasic()

				expectedError := badSubaccountId + ": " + ErrBadSubaccountID.Error()
				Ω(err.Error()).To(Equal(expectedError))
			})
		})

		Context("Without Sender field", func() {
			It("should be invalid with invalid address error", func() {
				msg.Sender = ""
				err := msg.ValidateBasic()

				expectedError := msg.Sender + ": " + sdkerrors.ErrInvalidAddress.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With wrong Sender field", func() {
			It("should be invalid with invalid address error", func() {
				msg.Sender = "0x90f8bf6a479f320ead074411a4b0e79ea8c9c1"
				err := msg.ValidateBasic()

				expectedError := msg.Sender + ": " + sdkerrors.ErrInvalidAddress.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With zero amount", func() {
			It("should be invalid with invalid coin error", func() {
				amountToDeposit := sdk.NewCoin(baseDenom, sdk.ZeroInt())
				msg.Amount = amountToDeposit
				err := msg.ValidateBasic()

				expectedError := sdk.ZeroInt().String() + baseDenom + ": " + sdkerrors.ErrInvalidCoins.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		// For negative deposit amounts or length of two characters or less for coin test is panicking.
	})

	Describe("Validate Withdraw", func() {
		var msg MsgWithdraw
		baseDenom := "INJ"

		BeforeEach(func() {
			amountToWithdraw := sdk.NewCoin(baseDenom, sdk.NewInt(50))
			subaccountID := "0x90f8bf6a479f320ead074411a4b0e7944ea8c9c1000000000000000000000001"
			sender := sdk.AccAddress(common.FromHex("90f8bf6a479f320ead074411a4b0e7944ea8c9c1"))

			msg = MsgWithdraw{
				Sender:       sender.String(),
				SubaccountId: subaccountID,
				Amount:       amountToWithdraw,
			}
		})

		Context("With all valid fields", func() {
			It("should be valid", func() {
				err := msg.ValidateBasic()

				Ω(err).To(BeNil())
			})
		})

		Context("With sender being different from owner of subaccount", func() {
			It("should be invalid with bad subaccountID error", func() {
				wrongSender := sdk.AccAddress(common.FromHex("90f8bf6a479f320ead074411a4b0e7934ea8c9c1"))
				msg.Sender = wrongSender.String()
				err := msg.ValidateBasic()

				expectedError := wrongSender.String() + ": " + ErrBadSubaccountID.Error()
				Ω(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With bad subaccountID", func() {
			It("should be invalid with bad subaccountID error", func() {
				badSubaccountId := "0x90f8bf6a47f320ead074411a4b0e7944ea8c9c1000000000000000000000001" // one less character
				msg.SubaccountId = badSubaccountId
				err := msg.ValidateBasic()

				expectedError := badSubaccountId + ": " + ErrBadSubaccountID.Error()
				Ω(err.Error()).To(Equal(expectedError))
			})
		})

		Context("Without Sender field", func() {
			It("should be invalid with invalid address error", func() {
				msg.Sender = ""
				err := msg.ValidateBasic()

				expectedError := msg.Sender + ": " + sdkerrors.ErrInvalidAddress.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With wrong Sender field", func() {
			It("should be invalid with invalid address error", func() {
				msg.Sender = "0x90f8bf6a479f320ead074411a4b0e79ea8c9c1"
				err := msg.ValidateBasic()

				expectedError := msg.Sender + ": " + sdkerrors.ErrInvalidAddress.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With zero amount", func() {
			It("should be invalid with invalid coin error", func() {
				amountToWithdraw := sdk.NewCoin(baseDenom, sdk.ZeroInt())
				msg.Amount = amountToWithdraw
				err := msg.ValidateBasic()

				expectedError := sdk.ZeroInt().String() + baseDenom + ": " + sdkerrors.ErrInvalidCoins.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})
	})

	Describe("Validate MsgInstantSpotMarketLaunch", func() {

		var msg MsgInstantSpotMarketLaunch

		BeforeEach(func() {
			sender := sdk.AccAddress(common.FromHex("90f8bf6a479f320ead074411a4b0e7944ea8c9c1"))
			msg = MsgInstantSpotMarketLaunch{
				Sender:     sender.String(),
				Ticker:     "INJ / ATOM",
				BaseDenom:  "inj",
				QuoteDenom: "atom",
			}
		})

		Context("With all valid fields", func() {
			It("should be valid", func() {
				err := msg.ValidateBasic()

				Ω(err).To(BeNil())
			})
		})

		Context("Without Sender field", func() {
			It("should be invalid with invalid address error", func() {
				msg.Sender = ""
				err := msg.ValidateBasic()

				expectedError := msg.Sender + ": " + sdkerrors.ErrInvalidAddress.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With wrong Sender field", func() {
			It("should be invalid with invalid address error", func() {
				msg.Sender = "0x90f8bf6a479f320ead074411a4b0e79ea8c9c1"
				err := msg.ValidateBasic()

				expectedError := msg.Sender + ": " + sdkerrors.ErrInvalidAddress.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("Without Ticker field", func() {
			It("should be invalid with invalid ticker error", func() {
				msg.Ticker = ""
				err := msg.ValidateBasic()

				expectedError := "ticker should not be empty: " + ErrInvalidTicker.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("Without BaseDenom field", func() {
			It("should be invalid with invalid base denom error", func() {
				msg.BaseDenom = ""
				err := msg.ValidateBasic()

				expectedError := "base denom should not be empty: " + ErrInvalidBaseDenom.Error()
				Expect(err.Error()).To(Equal(expectedError))

			})
		})

		Context("Without QuoteDenom field", func() {
			It("should be invalid with invalid quote denom error", func() {
				msg.QuoteDenom = ""
				err := msg.ValidateBasic()

				expectedError := "quote denom should not be empty: " + ErrInvalidQuoteDenom.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})
	})

	Describe("Validate MsgInstantPerpetualMarketLaunch", func() {

		var msg MsgInstantPerpetualMarketLaunch

		BeforeEach(func() {
			sender := sdk.AccAddress(common.FromHex("90f8bf6a479f320ead074411a4b0e7944ea8c9c1"))
			msg = MsgInstantPerpetualMarketLaunch{
				Sender:      sender.String(),
				Ticker:      "INJ / ATOM",
				QuoteDenom:  "inj",
				OracleBase:  "inj-band",
				OracleQuote: "atom-band",
				OracleType:  oracletypes.OracleType_Band,
			}
		})

		Context("With all valid fields", func() {
			It("should be valid", func() {
				err := msg.ValidateBasic()

				Ω(err).To(BeNil())
			})
		})

		Context("Without Sender field", func() {
			It("should be invalid with invalid address error", func() {
				msg.Sender = ""
				err := msg.ValidateBasic()

				expectedError := msg.Sender + ": " + sdkerrors.ErrInvalidAddress.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With wrong Sender field", func() {
			It("should be invalid with invalid address error", func() {
				msg.Sender = "0x90f8bf6a479f320ead074411a4b0e79ea8c9c1"
				err := msg.ValidateBasic()

				expectedError := msg.Sender + ": " + sdkerrors.ErrInvalidAddress.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("Without Ticker field", func() {
			It("should be invalid with invalid ticker error", func() {
				msg.Ticker = ""
				err := msg.ValidateBasic()

				expectedError := "ticker should not be empty: " + ErrInvalidTicker.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("Without QuoteDenom field", func() {
			It("should be invalid with invalid quote denom error", func() {
				msg.QuoteDenom = ""
				err := msg.ValidateBasic()

				expectedError := "quote denom should not be empty: " + ErrInvalidQuoteDenom.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("Without Oracle Base field", func() {
			It("should be invalid with invalid oracle error", func() {
				msg.OracleBase = ""
				err := msg.ValidateBasic()

				expectedError := "oracle base should not be empty: " + ErrInvalidOracle.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("Without Oracle Quote field", func() {
			It("should be invalid with invalid oracle error", func() {
				msg.OracleQuote = ""
				err := msg.ValidateBasic()

				expectedError := "oracle quote should not be empty: " + ErrInvalidOracle.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})
	})

	Describe("Validate MsgInstantExpiryFuturesMarketLaunch", func() {

		var msg MsgInstantExpiryFuturesMarketLaunch

		BeforeEach(func() {
			sender := sdk.AccAddress(common.FromHex("90f8bf6a479f320ead074411a4b0e7944ea8c9c1"))
			msg = MsgInstantExpiryFuturesMarketLaunch{
				Sender:      sender.String(),
				Ticker:      "INJ / ATOM",
				QuoteDenom:  "inj",
				OracleBase:  "inj-band",
				OracleQuote: "atom-band",
				OracleType:  oracletypes.OracleType_Band,
				Expiry:      10000,
			}
		})

		Context("With all valid fields", func() {
			It("should be valid", func() {
				err := msg.ValidateBasic()

				Ω(err).To(BeNil())
			})
		})

		Context("Without Sender field", func() {
			It("should be invalid with invalid address error", func() {
				msg.Sender = ""
				err := msg.ValidateBasic()

				expectedError := msg.Sender + ": " + sdkerrors.ErrInvalidAddress.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With wrong Sender field", func() {
			It("should be invalid with invalid address error", func() {
				msg.Sender = "0x90f8bf6a479f320ead074411a4b0e79ea8c9c1"
				err := msg.ValidateBasic()

				expectedError := msg.Sender + ": " + sdkerrors.ErrInvalidAddress.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("Without Ticker field", func() {
			It("should be invalid with invalid ticker error", func() {
				msg.Ticker = ""
				err := msg.ValidateBasic()

				expectedError := "ticker should not be empty: " + ErrInvalidTicker.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("Without QuoteDenom field", func() {
			It("should be invalid with invalid quote denom error", func() {
				msg.QuoteDenom = ""
				err := msg.ValidateBasic()

				expectedError := "quote denom should not be empty: " + ErrInvalidQuoteDenom.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("Without Oracle Base field", func() {
			It("should be invalid with invalid oracle error", func() {
				msg.OracleBase = ""
				err := msg.ValidateBasic()

				expectedError := "oracle base should not be empty: " + ErrInvalidOracle.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("Without Oracle Quote field", func() {
			It("should be invalid with invalid oracle error", func() {
				msg.OracleQuote = ""
				err := msg.ValidateBasic()

				expectedError := "oracle quote should not be empty: " + ErrInvalidOracle.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("Without Expiry field", func() {
			It("should be invalid with invalid expiry date error", func() {
				msg.Expiry = 0
				err := msg.ValidateBasic()

				expectedError := "expiry should not be empty: " + ErrInvalidExpiry.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})
	})

	Describe("Validate MsgCreateSpotLimitOrder", func() {

		var msg MsgCreateSpotLimitOrder

		BeforeEach(func() {
			subaccountID := "0x90f8bf6a479f320ead074411a4b0e7944ea8c9c1000000000000000000000001"
			sender := sdk.AccAddress(common.FromHex("0x90f8bf6a479f320ead074411a4b0e7944ea8c9c1"))

			dec := sdk.OneDec()
			msg = MsgCreateSpotLimitOrder{
				Sender: sender.String(),
				Order: SpotOrder{
					MarketId: "0xb0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3773",
					OrderInfo: OrderInfo{
						SubaccountId: subaccountID,
						FeeRecipient: "inj1dzqd00lfd4y4qy2pxa0dsdwzfnmsu27hgttswz",
						Price:        sdk.NewDec(137),
						Quantity:     sdk.NewDec(24),
					},
					OrderType:    0,
					TriggerPrice: dec,
				},
			}

			_, err := msg.Order.ComputeOrderHash(0)
			if err != nil {
				fmt.Println("ERROR: ", err)
			}

		})

		Context("With all valid fields", func() {
			It("should be valid", func() {
				err := msg.ValidateBasic()

				Ω(err).To(BeNil())
			})
		})

		Context("Without Sender field", func() {
			It("should be invalid with invalid address error", func() {
				msg.Sender = ""
				err := msg.ValidateBasic()

				expectedError := msg.Sender + ": " + sdkerrors.ErrInvalidAddress.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With wrong Sender field", func() {
			It("should be invalid with invalid address error", func() {
				msg.Sender = "0x90f8bf6a479f320ead074411a4b0e79ea8c9c1"
				err := msg.ValidateBasic()

				expectedError := msg.Sender + ": " + sdkerrors.ErrInvalidAddress.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("Without MarketId field", func() {
			It("should be invalid with key not found error", func() {
				msg.Order.MarketId = ""
				err := msg.ValidateBasic()

				expectedError := msg.Order.MarketId + ": " + ErrMarketInvalid.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With empty SubaccountId field", func() {
			It("should be invalid with bad subaccountId error", func() {
				msg.Order.OrderInfo.SubaccountId = ""
				err := msg.ValidateBasic()

				expectedError := msg.Order.OrderInfo.SubaccountId + ": " + ErrBadSubaccountID.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With wrong SubaccountId field", func() {
			It("should be invalid with bad subaccountId error", func() {
				msg.Order.OrderInfo.SubaccountId = "0xCA6A7F8C75B5EEACFDA20430CF5823CE4185673000000000000000000000001"
				err := msg.ValidateBasic()

				expectedError := msg.Order.OrderInfo.SubaccountId + ": " + ErrBadSubaccountID.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With not matching SubaccountId and Sender address", func() {
			It("should be invalid with bad subaccountId error", func() {
				msg.Order.OrderInfo.SubaccountId = "0x90f8bf6a479f320ead074411a4b0e7944ea8d9c1000000000000000000000001"
				err := msg.ValidateBasic()
				senderAddr, _ := sdk.AccAddressFromBech32(msg.Sender)

				expectedError := senderAddr.String() + ": " + ErrBadSubaccountID.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("Without FeeRecipient field", func() {
			It("should be invalid with invalid address error", func() {
				msg.Order.OrderInfo.FeeRecipient = ""
				err := msg.ValidateBasic()

				expectedError := msg.Order.OrderInfo.FeeRecipient + ": " + sdkerrors.ErrInvalidAddress.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With non-positive Price field", func() {
			It("should be invalid with invalid coins error when Price is smaller than 0", func() {
				msg.Order.OrderInfo.Price = sdk.NewDec(-1)
				err := msg.ValidateBasic()

				expectedError := msg.Order.OrderInfo.Price.String() + ": " + sdkerrors.ErrInvalidCoins.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
			It("should be invalid with invalid coins error when Price is equal to 0", func() {
				msg.Order.OrderInfo.Price = sdk.ZeroDec()
				err := msg.ValidateBasic()

				expectedError := msg.Order.OrderInfo.Price.String() + ": " + sdkerrors.ErrInvalidCoins.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With non-positive Quantity field", func() {
			It("should be invalid with invalid coins error when Quantity is smaller than 0", func() {
				msg.Order.OrderInfo.Quantity = sdk.NewDec(-1)
				err := msg.ValidateBasic()

				expectedError := msg.Order.OrderInfo.Quantity.String() + ": " + sdkerrors.ErrInvalidCoins.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
			It("should be invalid with invalid coins error when Quantity is equal to 0", func() {
				msg.Order.OrderInfo.Quantity = sdk.ZeroDec()
				err := msg.ValidateBasic()

				expectedError := msg.Order.OrderInfo.Quantity.String() + ": " + sdkerrors.ErrInvalidCoins.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With invalid OrderType field", func() {
			It("should be invalid with key not found error when OrderType is smaller than 0", func() {
				msg.Order.OrderType = -1
				err := msg.ValidateBasic()

				expectedError := string(msg.Order.OrderType) + ": " + ErrUnrecognizedOrderType.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
			It("should be invalid with key not found error when OrderType is greater than 5", func() {
				msg.Order.OrderType = 6
				err := msg.ValidateBasic()

				expectedError := string(msg.Order.OrderType) + ": " + ErrUnrecognizedOrderType.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With negative TriggerPrice field", func() {
			It("should be invalid with invalid coins error when TriggerPrice is smaller than 0", func() {
				minusDec := sdk.NewDec(-1)
				msg.Order.TriggerPrice = minusDec
				err := msg.ValidateBasic()

				expectedError := msg.Order.TriggerPrice.String() + ": " + sdkerrors.ErrInvalidCoins.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

	})

	Describe("Validate MsgCreateSpotMarketOrder", func() {

		var msg MsgCreateSpotMarketOrder

		BeforeEach(func() {
			subaccountID := "0x90f8bf6a479f320ead074411a4b0e7944ea8c9c1000000000000000000000001"
			sender := sdk.AccAddress(common.FromHex("90f8bf6a479f320ead074411a4b0e7944ea8c9c1"))

			dec := sdk.OneDec()
			msg = MsgCreateSpotMarketOrder{
				Sender: sender.String(),
				Order: SpotOrder{
					MarketId: "0xb0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3773",
					OrderInfo: OrderInfo{
						SubaccountId: subaccountID,
						FeeRecipient: "inj1dzqd00lfd4y4qy2pxa0dsdwzfnmsu27hgttswz",
						Price:        sdk.ZeroDec(),
						Quantity:     sdk.NewDec(24),
					},
					OrderType:    0,
					TriggerPrice: dec,
				},
			}

			_, err := msg.Order.ComputeOrderHash(0)
			if err != nil {
				fmt.Println("ERROR: ", err)
			}

		})

		Context("With all valid fields", func() {
			It("should be valid", func() {
				err := msg.ValidateBasic()

				Ω(err).To(BeNil())
			})
		})

		Context("Without Sender field", func() {
			It("should be invalid with invalid address error", func() {
				msg.Sender = ""
				err := msg.ValidateBasic()

				expectedError := msg.Sender + ": " + sdkerrors.ErrInvalidAddress.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With wrong Sender field", func() {
			It("should be invalid with invalid address error", func() {
				msg.Sender = "0x90f8bf6a479f320ead074411a4b0e79ea8c9c1"
				err := msg.ValidateBasic()

				expectedError := msg.Sender + ": " + sdkerrors.ErrInvalidAddress.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("Without MarketId field", func() {
			It("should be invalid with key not found error", func() {
				msg.Order.MarketId = ""
				err := msg.ValidateBasic()

				expectedError := msg.Order.MarketId + ": " + ErrMarketInvalid.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With empty SubaccountId field", func() {
			It("should be invalid with bad subaccountId error", func() {
				msg.Order.OrderInfo.SubaccountId = ""
				err := msg.ValidateBasic()

				expectedError := msg.Order.OrderInfo.SubaccountId + ": " + ErrBadSubaccountID.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With wrong SubaccountId field", func() {
			It("should be invalid with bad subaccountId error", func() {
				msg.Order.OrderInfo.SubaccountId = "0xCA6A7F8C75B5EEACFDA20430CF5823CE4185673000000000000000000000001"
				err := msg.ValidateBasic()

				expectedError := msg.Order.OrderInfo.SubaccountId + ": " + ErrBadSubaccountID.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With wrong SubaccountId field", func() {
			It("should be invalid with bad subaccountId error", func() {
				msg.Order.OrderInfo.SubaccountId = "0xCA6A7F8C75B5EEACFDA20430CF5823CE4185673000000000000000000000001"
				err := msg.ValidateBasic()

				expectedError := msg.Order.OrderInfo.SubaccountId + ": " + ErrBadSubaccountID.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With not matching SubaccountId and Sender address", func() {
			It("should be invalid with bad subaccountId error", func() {
				msg.Order.OrderInfo.SubaccountId = "0x90f8bf6a479f320ead074411a4b0e7944ea8d9c1000000000000000000000001"
				err := msg.ValidateBasic()
				senderAddr, _ := sdk.AccAddressFromBech32(msg.Sender)

				expectedError := senderAddr.String() + ": " + ErrBadSubaccountID.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("Without FeeRecipient field", func() {
			It("should be invalid with invalid address error", func() {
				msg.Order.OrderInfo.FeeRecipient = ""
				err := msg.ValidateBasic()

				expectedError := msg.Order.OrderInfo.FeeRecipient + ": " + sdkerrors.ErrInvalidAddress.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With non-positive Quantity field", func() {
			It("should be invalid with invalid coins error when Quantity is smaller than 0", func() {
				msg.Order.OrderInfo.Quantity = sdk.NewDec(-1)
				err := msg.ValidateBasic()

				expectedError := msg.Order.OrderInfo.Quantity.String() + ": " + sdkerrors.ErrInvalidCoins.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
			It("should be invalid with invalid coins error when Quantity is equal to 0", func() {
				msg.Order.OrderInfo.Quantity = sdk.ZeroDec()
				err := msg.ValidateBasic()

				expectedError := msg.Order.OrderInfo.Quantity.String() + ": " + sdkerrors.ErrInvalidCoins.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With invalid OrderType field", func() {
			It("should be invalid with key not found error when OrderType is smaller than 0", func() {
				msg.Order.OrderType = -1
				err := msg.ValidateBasic()

				expectedError := string(msg.Order.OrderType) + ": " + ErrUnrecognizedOrderType.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
			It("should be invalid with key not found error when OrderType is greater than 5", func() {
				msg.Order.OrderType = 6
				err := msg.ValidateBasic()

				expectedError := string(msg.Order.OrderType) + ": " + ErrUnrecognizedOrderType.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With negative TriggerPrice field", func() {
			It("should be invalid with invalid coins error when TriggerPrice is smaller than 0", func() {
				minusDec := sdk.NewDec(-1)
				msg.Order.TriggerPrice = minusDec
				err := msg.ValidateBasic()

				expectedError := msg.Order.TriggerPrice.String() + ": " + sdkerrors.ErrInvalidCoins.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})
	})

	Describe("Validate CancelSpotOrder", func() {
		var msg MsgCancelSpotOrder

		BeforeEach(func() {
			sender := sdk.AccAddress(common.FromHex("90f8bf6a479f320ead074411a4b0e7944ea8c9c1"))

			msg = types.MsgCancelSpotOrder{
				Sender:       sender.String(),
				MarketId:     "0xb0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3773",
				IsBuy:        true,
				SubaccountId: "0x90f8bf6a479f320ead074411a4b0e7944ea8c9c1000000000000000000000001",
				OrderHash:    "0x5cf90f9026695a5650035f8a6c92c5294787b18032f08ce45460ee9b6bc63989",
			}
		})

		Context("With all valid fields", func() {
			It("should be valid", func() {
				err := msg.ValidateBasic()

				Ω(err).To(BeNil())
			})
		})

		Context("With bad subaccountID", func() {
			It("should be invalid with bad subaccountID error", func() {
				badSubaccountId := "0x90f8bf6a47f320ead074411a4b0e7944ea8c9c1000000000000000000000001" // one less character
				msg.SubaccountId = badSubaccountId
				err := msg.ValidateBasic()

				expectedError := badSubaccountId + ": " + ErrBadSubaccountID.Error()
				Ω(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With subaccount not owned by the sender", func() {
			It("should be invalid with bad subaccountID error", func() {
				notOwnedSubaccountId := "0x90f8bf6a478f320ead074411a4b0e7944ea8c9c1000000000000000000000001" // one less character
				msg.SubaccountId = notOwnedSubaccountId
				err := msg.ValidateBasic()

				expectedError := msg.Sender + ": " + ErrBadSubaccountID.Error()
				Ω(err.Error()).To(Equal(expectedError))
			})
		})

		Context("Without Sender field", func() {
			It("should be invalid with invalid address error", func() {
				msg.Sender = ""
				err := msg.ValidateBasic()

				expectedError := msg.Sender + ": " + sdkerrors.ErrInvalidAddress.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With wrong Sender field", func() {
			It("should be invalid with invalid address error", func() {
				msg.Sender = "0x90f8bf6a479f320ead074411a4b0e79ea8c9c1"
				err := msg.ValidateBasic()

				expectedError := msg.Sender + ": " + sdkerrors.ErrInvalidAddress.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With wrong Hash field", func() {
			It("should be invalid with invalid hash error", func() {
				msg.OrderHash = "0x90f8bf6a479f320ead074411a4b0e79ea8c9c1"
				err := msg.ValidateBasic()

				expectedError := msg.OrderHash + ": " + ErrOrderHashInvalid.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})
	})

	Describe("Validate SubaccountTransfer", func() {
		var msg MsgSubaccountTransfer
		baseDenom := "INJ"

		BeforeEach(func() {
			amountToTransfer := sdk.NewCoin(baseDenom, sdk.NewInt(50))
			sourceSubaccountID := "0x90f8bf6a479f320ead074411a4b0e7944ea8c9c1000000000000000000000001"
			destinationSubaccountID := "0x90f8bf6a479f320ead074411a4b0e7944ea8c9c1000000000000000000000002"
			sender := sdk.AccAddress(common.FromHex("90f8bf6a479f320ead074411a4b0e7944ea8c9c1"))

			msg = MsgSubaccountTransfer{
				Sender:                  sender.String(),
				SourceSubaccountId:      sourceSubaccountID,
				DestinationSubaccountId: destinationSubaccountID,
				Amount:                  amountToTransfer,
			}
		})

		Context("With all valid fields", func() {
			It("should be valid", func() {
				err := msg.ValidateBasic()

				Ω(err).To(BeNil())
			})
		})

		Context("With bad sourceSubaccountID", func() {
			It("should be invalid with bad subaccountID error", func() {
				badSubaccountId := "0x90f8bf6a47f320ead074411a4b0e7944ea8c9c1000000000000000000000001" // one less character
				msg.SourceSubaccountId = badSubaccountId
				err := msg.ValidateBasic()

				expectedError := badSubaccountId + ": " + ErrBadSubaccountID.Error()
				Ω(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With not owned by sender sourceSubaccountID", func() {
			It("should be invalid with bad subaccountID error", func() {
				notOwnedSubaccountId := "0x90f8bf6a479f320ead074411a4b0e7944ea7c9c1000000000000000000000001"
				msg.SourceSubaccountId = notOwnedSubaccountId
				err := msg.ValidateBasic()

				expectedError := msg.DestinationSubaccountId + ": " + ErrBadSubaccountID.Error()
				Ω(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With bad destinationSubaccountID", func() {
			It("should be invalid with bad subaccountID error", func() {
				badSubaccountId := "0x90f8bf6a47f320ead074411a4b0e7944ea8c9c1000000000000000000000001" // one less character
				msg.DestinationSubaccountId = badSubaccountId
				err := msg.ValidateBasic()

				expectedError := badSubaccountId + ": " + ErrBadSubaccountID.Error()
				Ω(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With not owned by sender destinationSubaccountID", func() {
			It("should be invalid with bad subaccountID error", func() {
				notOwnedSubaccountId := "0x90f8bf6a479f320ead074411a4b0e7944ea7c9c1000000000000000000000002"
				msg.DestinationSubaccountId = notOwnedSubaccountId
				err := msg.ValidateBasic()

				expectedError := notOwnedSubaccountId + ": " + ErrBadSubaccountID.Error()
				Ω(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With subaccounts of the same address that is not the sender", func() {
			It("should be invalid with bad subaccountID error", func() {
				notOwnedSourceSubaccountId := "0x90f8bf6a479f320ead074411a4b0e7944ea7c9c1000000000000000000000001"
				notOwnedDestinationSubaccountId := "0x90f8bf6a479f320ead074411a4b0e7944ea7c9c1000000000000000000000002"
				msg.SourceSubaccountId = notOwnedSourceSubaccountId
				msg.DestinationSubaccountId = notOwnedDestinationSubaccountId
				err := msg.ValidateBasic()

				expectedError := msg.Sender + ": " + ErrBadSubaccountID.Error()
				Ω(err.Error()).To(Equal(expectedError))
			})
		})

		Context("Without Sender field", func() {
			It("should be invalid with invalid address error", func() {
				msg.Sender = ""
				err := msg.ValidateBasic()

				expectedError := msg.Sender + ": " + sdkerrors.ErrInvalidAddress.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With wrong Sender field", func() {
			It("should be invalid with invalid address error", func() {
				msg.Sender = "0x90f8bf6a479f320ead074411a4b0e79ea8c9c1"
				err := msg.ValidateBasic()

				expectedError := msg.Sender + ": " + sdkerrors.ErrInvalidAddress.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With zero amount", func() {
			It("should be invalid with invalid coin error", func() {
				amountToTransfer := sdk.NewCoin(baseDenom, sdk.ZeroInt())
				msg.Amount = amountToTransfer
				err := msg.ValidateBasic()

				expectedError := sdk.ZeroInt().String() + baseDenom + ": " + sdkerrors.ErrInvalidCoins.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})
	})

	Describe("Validate ExternalTransfer", func() {
		var msg MsgExternalTransfer
		baseDenom := "INJ"

		BeforeEach(func() {
			amountToTransfer := sdk.NewCoin(baseDenom, sdk.NewInt(50))
			sourceSubaccountID := "0x90f8bf6a479f320ead074411a4b0e7944ea8c9c1000000000000000000000001"
			destinationSubaccountID := "0x90f8bf6a479f320ead074411a4b0e7944ea8c9c1000000000000000000000002"
			sender := sdk.AccAddress(common.FromHex("90f8bf6a479f320ead074411a4b0e7944ea8c9c1"))

			msg = MsgExternalTransfer{
				Sender:                  sender.String(),
				SourceSubaccountId:      sourceSubaccountID,
				DestinationSubaccountId: destinationSubaccountID,
				Amount:                  amountToTransfer,
			}
		})

		Context("With all valid fields", func() {
			It("should be valid", func() {
				err := msg.ValidateBasic()

				Ω(err).To(BeNil())
			})
		})

		Context("With bad sourceSubaccountID", func() {
			It("should be invalid with bad subaccountID error", func() {
				badSubaccountId := "0x90f8bf6a47f320ead074411a4b0e7944ea8c9c1000000000000000000000001" // one less character
				msg.SourceSubaccountId = badSubaccountId
				err := msg.ValidateBasic()

				expectedError := badSubaccountId + ": " + ErrBadSubaccountID.Error()
				Ω(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With not owned by sender sourceSubaccountID", func() {
			It("should be invalid with bad subaccountID error", func() {
				notOwnedSubaccountId := "0x90f8bf6a479f320ead074411a4b0e7944ea7c9c1000000000000000000000001"
				msg.SourceSubaccountId = notOwnedSubaccountId
				err := msg.ValidateBasic()

				expectedError := msg.Sender + ": " + ErrBadSubaccountID.Error()
				Ω(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With bad destinationSubaccountID", func() {
			It("should be invalid with bad subaccountID error", func() {
				badSubaccountId := "0x90f8bf6a47f320ead074411a4b0e7944ea8c9c1000000000000000000000001" // one less character
				msg.DestinationSubaccountId = badSubaccountId
				err := msg.ValidateBasic()

				expectedError := badSubaccountId + ": " + ErrBadSubaccountID.Error()
				Ω(err.Error()).To(Equal(expectedError))
			})
		})

		Context("Without Sender field", func() {
			It("should be invalid with invalid address error", func() {
				msg.Sender = ""
				err := msg.ValidateBasic()

				expectedError := msg.Sender + ": " + sdkerrors.ErrInvalidAddress.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With wrong Sender field", func() {
			It("should be invalid with invalid address error", func() {
				msg.Sender = "0x90f8bf6a479f320ead074411a4b0e79ea8c9c1"
				err := msg.ValidateBasic()

				expectedError := msg.Sender + ": " + sdkerrors.ErrInvalidAddress.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With zero amount", func() {
			It("should be invalid with invalid coin error", func() {
				amountToTransfer := sdk.NewCoin(baseDenom, sdk.ZeroInt())
				msg.Amount = amountToTransfer
				err := msg.ValidateBasic()

				expectedError := sdk.ZeroInt().String() + baseDenom + ": " + sdkerrors.ErrInvalidCoins.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})
	})
})
