package types_test

import (
	"fmt"

	. "github.com/InjectiveLabs/injective-core/injective-chain/modules/exchange/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
)

var _ = Describe("Msgs", func() {

	Describe("Validate MsgInstantSpotMarketLaunch", func() {

		var msg MsgInstantSpotMarketLaunch

		BeforeEach(func() {
			msg = MsgInstantSpotMarketLaunch{
				Sender:     "inj1wfawuv6fslzjlfa4v7exv27mk6rpfeyvhvxchc",
				Ticker:     "INJ / ATOM",
				BaseDenom:  "inj",
				QuoteDenom: "atom",
			}
		})

		Context("Without Sender field", func() {
			It("should be invalid msg", func() {
				msg.Sender = ""
				err := msg.ValidateBasic()
				Expect(err).NotTo(BeNil())
			})
		})

		Context("Without Ticker field", func() {
			It("should be invalid msg", func() {
				msg.Ticker = ""
				err := msg.ValidateBasic()
				Expect(err).NotTo(BeNil())
			})
		})

		Context("Without BaseDenom field", func() {
			It("should be invalid msg", func() {
				msg.BaseDenom = ""
				err := msg.ValidateBasic()
				Expect(err).NotTo(BeNil())
			})
		})

		Context("Without QuoteDenom field", func() {
			It("should be invalid msg", func() {
				msg.QuoteDenom = ""
				err := msg.ValidateBasic()
				Expect(err).NotTo(BeNil())
			})
		})
	})

	Describe("Validate MsgInstantPerpetualMarketLaunch", func() {

		var msg MsgInstantPerpetualMarketLaunch

		BeforeEach(func() {
			msg = MsgInstantPerpetualMarketLaunch{
				Sender:     "inj1wfawuv6fslzjlfa4v7exv27mk6rpfeyvhvxchc",
				Ticker:     "INJ / ATOM",
				QuoteDenom: "inj",
				Oracle:     "oracle",
			}
		})

		Context("Without Sender field", func() {
			It("should be invalid msg", func() {
				msg.Sender = ""
				err := msg.ValidateBasic()
				Expect(err).NotTo(BeNil())
			})
		})

		Context("Without Ticker field", func() {
			It("should be invalid msg", func() {
				msg.Ticker = ""
				err := msg.ValidateBasic()
				Expect(err).NotTo(BeNil())
			})
		})

		Context("Without QuoteDenom field", func() {
			It("should be invalid msg", func() {
				msg.QuoteDenom = ""
				err := msg.ValidateBasic()
				Expect(err).NotTo(BeNil())
			})
		})

		Context("Without Oracle field", func() {
			It("should be invalid msg", func() {
				msg.Oracle = ""
				err := msg.ValidateBasic()
				Expect(err).NotTo(BeNil())
			})
		})
	})

	Describe("Validate MsgInstantExpiryFuturesMarketLaunch", func() {

		var msg MsgInstantExpiryFuturesMarketLaunch

		BeforeEach(func() {
			msg = MsgInstantExpiryFuturesMarketLaunch{
				Sender:     "inj1wfawuv6fslzjlfa4v7exv27mk6rpfeyvhvxchc",
				Ticker:     "INJ / ATOM",
				QuoteDenom: "inj",
				Oracle:     "oracle",
				Expiry:     10000,
			}
		})

		Context("Without Sender field", func() {
			It("should be invalid msg", func() {
				msg.Sender = ""
				err := msg.ValidateBasic()
				Expect(err).NotTo(BeNil())
			})
		})

		Context("Without Ticker field", func() {
			It("should be invalid msg", func() {
				msg.Ticker = ""
				err := msg.ValidateBasic()
				Expect(err).NotTo(BeNil())
			})
		})

		Context("Without QuoteDenom field", func() {
			It("should be invalid msg", func() {
				msg.QuoteDenom = ""
				err := msg.ValidateBasic()
				Expect(err).NotTo(BeNil())
			})
		})

		Context("Without Oracle field", func() {
			It("should be invalid msg", func() {
				msg.Oracle = ""
				err := msg.ValidateBasic()
				Expect(err).NotTo(BeNil())
			})
		})

		Context("Without Expiry field", func() {
			It("should be invalid msg", func() {
				msg.Expiry = 0
				err := msg.ValidateBasic()
				Expect(err).NotTo(BeNil())
			})
		})
	})

	Describe("Validate MsgCreateSpotLimitOrder", func() {

		var msg MsgCreateSpotLimitOrder

		BeforeEach(func() {
			// Pass the private key to generate the address
			var privateKey = "DA6336E19777F2CB5EC1CAD901BB858ADEAAB371344FCD491BEA3E2A15708DC1"
			pk, err := crypto.HexToECDSA(privateKey)
			if err != nil {
				err = errors.Wrap(err, "failed to hex-decode cosmos account privkey")
			}

			ethAddress := crypto.PubkeyToAddress(pk.PublicKey)

			dec := sdk.OneDec()
			msg = MsgCreateSpotLimitOrder{
				Sender: ethAddress.String(),
				Order: SpotOrder{
					MarketId: "0xb0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3773",
					OrderInfo: OrderInfo{
						SubaccountId: "CA6A7F8C75B5EEACFDA204309CF5823CE4185673000000000000000000000001",
						FeeRecipient: "inj1dzqd00lfd4y4qy2pxa0dsdwzfnmsu27hgttswz",
						Price:        sdk.NewDec(137),
						Quantity:     sdk.NewDec(24),
					},
					OrderType:    0,
					TriggerPrice: dec,
					Salt:         12,
				},
			}

			_, err = msg.Order.ComputeOrderHash()
			if err != nil {
				fmt.Println("ERROR: ", err)
			}

		})

		Context("Without Sender field", func() {
			It("should be invalid with invalid address error", func() {
				msg.Sender = ""
				err := msg.ValidateBasic()

				// Construct empty address alongside with error message
				expectedError := msg.Sender + ": " + sdkerrors.ErrInvalidAddress.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("Without MarketId field", func() {
			It("should be invalid with key not found error", func() {
				msg.Order.MarketId = ""
				err := msg.ValidateBasic()

				expectedError := msg.Order.MarketId + ": " + sdkerrors.ErrKeyNotFound.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("Without SubaccountId field", func() {
			It("should be invalid with key not found error", func() {
				msg.Order.OrderInfo.SubaccountId = ""
				err := msg.ValidateBasic()

				expectedError := msg.Order.OrderInfo.SubaccountId + ": " + sdkerrors.ErrKeyNotFound.Error()
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
				msg.Order.OrderInfo.Price = sdk.NewDec(0)
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
				msg.Order.OrderInfo.Quantity = sdk.NewDec(0)
				err := msg.ValidateBasic()

				expectedError := msg.Order.OrderInfo.Quantity.String() + ": " + sdkerrors.ErrInvalidCoins.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With invalid OrderType field", func() {
			It("should be invalid with key not found error when OrderType is smaller than 0", func() {
				msg.Order.OrderType = -1
				err := msg.ValidateBasic()

				expectedError := string(msg.Order.OrderType) + ": " + sdkerrors.ErrKeyNotFound.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
			It("should be invalid with key not found error when OrderType is greater than 5", func() {
				msg.Order.OrderType = 6
				err := msg.ValidateBasic()

				expectedError := string(msg.Order.OrderType) + ": " + sdkerrors.ErrKeyNotFound.Error()
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
			// Pass the private key to generate the address
			var privateKey = "DA6336E19777F2CB5EC1CAD901BB858ADEAAB371344FCD491BEA3E2A15708DC1"
			pk, err := crypto.HexToECDSA(privateKey)
			if err != nil {
				err = errors.Wrap(err, "failed to hex-decode cosmos account privkey")
			}

			ethAddress := crypto.PubkeyToAddress(pk.PublicKey)

			dec := sdk.OneDec()
			msg = MsgCreateSpotMarketOrder{
				Sender: ethAddress.String(),
				Order: SpotOrder{
					MarketId: "0xb0057716d5917badaf911b193b12b910811c1497b5bada8d7711f758981c3773",
					OrderInfo: OrderInfo{
						SubaccountId: "CA6A7F8C75B5EEACFDA204309CF5823CE4185673000000000000000000000001",
						FeeRecipient: "inj1dzqd00lfd4y4qy2pxa0dsdwzfnmsu27hgttswz",
						Price:        sdk.NewDec(0),
						Quantity:     sdk.NewDec(24),
					},
					OrderType:    0,
					TriggerPrice: dec,
					Salt:         12,
				},
			}

			_, err = msg.Order.ComputeOrderHash()
			if err != nil {
				fmt.Println("ERROR: ", err)
			}

		})

		Context("Without Sender field", func() {
			It("should be invalid with invalid address error", func() {
				msg.Sender = ""
				err := msg.ValidateBasic()

				// Construct empty address alongside with error message
				expectedError := msg.Sender + ": " + sdkerrors.ErrInvalidAddress.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("Without MarketId field", func() {
			It("should be invalid with key not found error", func() {
				msg.Order.MarketId = ""
				err := msg.ValidateBasic()

				expectedError := msg.Order.MarketId + ": " + sdkerrors.ErrKeyNotFound.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("Without SubaccountId field", func() {
			It("should be invalid with key not found error", func() {
				msg.Order.OrderInfo.SubaccountId = ""
				err := msg.ValidateBasic()

				expectedError := msg.Order.OrderInfo.SubaccountId + ": " + sdkerrors.ErrKeyNotFound.Error()
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
				msg.Order.OrderInfo.Quantity = sdk.NewDec(0)
				err := msg.ValidateBasic()

				expectedError := msg.Order.OrderInfo.Quantity.String() + ": " + sdkerrors.ErrInvalidCoins.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
		})

		Context("With invalid OrderType field", func() {
			It("should be invalid with key not found error when OrderType is smaller than 0", func() {
				msg.Order.OrderType = -1
				err := msg.ValidateBasic()

				expectedError := string(msg.Order.OrderType) + ": " + sdkerrors.ErrKeyNotFound.Error()
				Expect(err.Error()).To(Equal(expectedError))
			})
			It("should be invalid with key not found error when OrderType is greater than 5", func() {
				msg.Order.OrderType = 6
				err := msg.ValidateBasic()

				expectedError := string(msg.Order.OrderType) + ": " + sdkerrors.ErrKeyNotFound.Error()
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

})
