package types

import (
	"math/big"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

const (
	// DefaultEIP155ChainID
	DefaultEIP155ChainID = 1776
)

// EthereumConfig returns an Ethereum ChainConfig for EVM state transitions.
// All the negative or nil values are converted to nil
func (cc ChainConfig) EthereumConfig() *params.ChainConfig {
	cfg := &params.ChainConfig{
		ChainID:                 cc.EIP155ChainID.BigIntMut(),
		HomesteadBlock:          getBlockValue(cc.HomesteadBlock),
		DAOForkBlock:            getBlockValue(cc.DAOForkBlock),
		DAOForkSupport:          cc.DAOForkSupport,
		EIP150Block:             getBlockValue(cc.EIP150Block),
		EIP155Block:             getBlockValue(cc.EIP155Block),
		EIP158Block:             getBlockValue(cc.EIP158Block),
		ByzantiumBlock:          getBlockValue(cc.ByzantiumBlock),
		ConstantinopleBlock:     getBlockValue(cc.ConstantinopleBlock),
		PetersburgBlock:         getBlockValue(cc.PetersburgBlock),
		IstanbulBlock:           getBlockValue(cc.IstanbulBlock),
		MuirGlacierBlock:        getBlockValue(cc.MuirGlacierBlock),
		BerlinBlock:             getBlockValue(cc.BerlinBlock),
		LondonBlock:             getBlockValue(cc.LondonBlock),
		ArrowGlacierBlock:       getBlockValue(cc.ArrowGlacierBlock),
		GrayGlacierBlock:        getBlockValue(cc.GrayGlacierBlock),
		MergeNetsplitBlock:      getBlockValue(cc.MergeNetsplitBlock),
		TerminalTotalDifficulty: nil,
		Ethash:                  nil,
		Clique:                  nil,
		ShanghaiTime:            getTimeValue(cc.ShanghaiTime),
		CancunTime:              getTimeValue(cc.CancunTime),
		PragueTime:              getTimeValue(cc.PragueTime),
		BlobScheduleConfig:      cc.BlobScheduleConfig.ToEthereumBlobScheduleConfig(),
	}
	return cfg
}

// DefaultChainConfig returns default evm parameters.
func DefaultChainConfig() ChainConfig {
	homesteadBlock := sdkmath.ZeroInt()
	daoForkBlock := sdkmath.ZeroInt()
	eip150Block := sdkmath.ZeroInt()
	eip155Block := sdkmath.ZeroInt()
	eip158Block := sdkmath.ZeroInt()
	byzantiumBlock := sdkmath.ZeroInt()
	constantinopleBlock := sdkmath.ZeroInt()
	petersburgBlock := sdkmath.ZeroInt()
	istanbulBlock := sdkmath.ZeroInt()
	muirGlacierBlock := sdkmath.ZeroInt()
	berlinBlock := sdkmath.ZeroInt()
	londonBlock := sdkmath.ZeroInt()
	arrowGlacierBlock := sdkmath.ZeroInt()
	grayGlacierBlock := sdkmath.ZeroInt()
	mergeNetsplitBlock := sdkmath.ZeroInt()
	shanghaiTime := sdkmath.ZeroInt()
	pragueTime := sdkmath.ZeroInt()
	cancunTime := sdkmath.ZeroInt()
	chainID := sdkmath.NewInt(DefaultEIP155ChainID)
	defaultBlobSchedule := FromEthereumBlobScheduleConfig(params.DefaultBlobSchedule)

	return ChainConfig{
		HomesteadBlock:      &homesteadBlock,
		DAOForkBlock:        &daoForkBlock,
		DAOForkSupport:      true,
		EIP150Block:         &eip150Block,
		EIP150Hash:          common.Hash{}.String(),
		EIP155Block:         &eip155Block,
		EIP158Block:         &eip158Block,
		ByzantiumBlock:      &byzantiumBlock,
		ConstantinopleBlock: &constantinopleBlock,
		PetersburgBlock:     &petersburgBlock,
		IstanbulBlock:       &istanbulBlock,
		MuirGlacierBlock:    &muirGlacierBlock,
		BerlinBlock:         &berlinBlock,
		LondonBlock:         &londonBlock,
		ArrowGlacierBlock:   &arrowGlacierBlock,
		GrayGlacierBlock:    &grayGlacierBlock,
		MergeNetsplitBlock:  &mergeNetsplitBlock,
		ShanghaiTime:        &shanghaiTime,
		CancunTime:          &cancunTime,
		PragueTime:          &pragueTime,
		EIP155ChainID:       &chainID,
		BlobScheduleConfig:  defaultBlobSchedule,
	}
}

func getBlockValue(block *sdkmath.Int) *big.Int {
	if block == nil || block.IsNegative() {
		return nil
	}

	return block.BigIntMut()
}

func getTimeValue(time *sdkmath.Int) *uint64 {
	if time == nil || time.IsNegative() {
		return nil
	}
	t := time.BigIntMut().Uint64()
	return &t
}

// Validate performs a basic validation of the ChainConfig params. The function will return an error
// if any of the block values is uninitialized (i.e nil) or if the EIP150Hash is an invalid hash.
func (cc ChainConfig) Validate() error {
	if err := ValidateBlock(cc.HomesteadBlock); err != nil {
		return errorsmod.Wrap(err, "homesteadBlock")
	}
	if err := ValidateBlock(cc.DAOForkBlock); err != nil {
		return errorsmod.Wrap(err, "daoForkBlock")
	}
	if err := ValidateBlock(cc.EIP150Block); err != nil {
		return errorsmod.Wrap(err, "eip150Block")
	}
	if err := ValidateHash(cc.EIP150Hash); err != nil {
		return err
	}
	if err := ValidateBlock(cc.EIP155Block); err != nil {
		return errorsmod.Wrap(err, "eip155Block")
	}
	if err := ValidateBlock(cc.EIP158Block); err != nil {
		return errorsmod.Wrap(err, "eip158Block")
	}
	if err := ValidateBlock(cc.ByzantiumBlock); err != nil {
		return errorsmod.Wrap(err, "byzantiumBlock")
	}
	if err := ValidateBlock(cc.ConstantinopleBlock); err != nil {
		return errorsmod.Wrap(err, "constantinopleBlock")
	}
	if err := ValidateBlock(cc.PetersburgBlock); err != nil {
		return errorsmod.Wrap(err, "petersburgBlock")
	}
	if err := ValidateBlock(cc.IstanbulBlock); err != nil {
		return errorsmod.Wrap(err, "istanbulBlock")
	}
	if err := ValidateBlock(cc.MuirGlacierBlock); err != nil {
		return errorsmod.Wrap(err, "muirGlacierBlock")
	}
	if err := ValidateBlock(cc.BerlinBlock); err != nil {
		return errorsmod.Wrap(err, "berlinBlock")
	}
	if err := ValidateBlock(cc.LondonBlock); err != nil {
		return errorsmod.Wrap(err, "londonBlock")
	}
	if err := ValidateBlock(cc.ArrowGlacierBlock); err != nil {
		return errorsmod.Wrap(err, "arrowGlacierBlock")
	}
	if err := ValidateBlock(cc.GrayGlacierBlock); err != nil {
		return errorsmod.Wrap(err, "GrayGlacierBlock")
	}
	if err := ValidateBlock(cc.MergeNetsplitBlock); err != nil {
		return errorsmod.Wrap(err, "MergeNetsplitBlock")
	}
	if err := ValidateTime(cc.ShanghaiTime); err != nil {
		return errorsmod.Wrap(err, "ShanghaiTime")
	}
	if err := ValidateTime(cc.CancunTime); err != nil {
		return errorsmod.Wrap(err, "CancunTime")
	}
	if err := ValidateTime(cc.PragueTime); err != nil {
		return errorsmod.Wrap(err, "PragueTime")
	}
	if err := cc.EthereumConfig().CheckConfigForkOrder(); err != nil {
		return errorsmod.Wrap(err, "invalid config fork order")
	}
	return nil
}

func ValidateHash(hex string) error {
	if hex != "" && strings.TrimSpace(hex) == "" {
		return errorsmod.Wrap(ErrInvalidChainConfig, "hash cannot be blank")
	}

	return nil
}

func ValidateBlock(block *sdkmath.Int) error {
	// nil value means that the fork has not yet been applied
	if block == nil {
		return nil
	}

	if block.IsNegative() {
		return errorsmod.Wrapf(
			ErrInvalidChainConfig, "block value cannot be negative: %s", block,
		)
	}

	return nil
}

func ValidateTime(time *sdkmath.Int) error {
	// nil value means that the fork has not yet been applied
	if time == nil {
		return nil
	}

	if time.IsNegative() {
		return errorsmod.Wrapf(
			ErrInvalidChainConfig, "time value cannot be negative: %s", time,
		)
	}

	return nil
}

func FromEthereumBlobScheduleConfig(ebsc *params.BlobScheduleConfig) *BlobScheduleConfig {
	if ebsc == nil {
		return nil
	}

	return &BlobScheduleConfig{
		Cancun: FromEthereumBlobConfig(ebsc.Cancun),
		Prague: FromEthereumBlobConfig(ebsc.Prague),
		Osaka:  FromEthereumBlobConfig(ebsc.Osaka),
		Verkle: FromEthereumBlobConfig(ebsc.Verkle),
	}
}

func (bsc *BlobScheduleConfig) ToEthereumBlobScheduleConfig() *params.BlobScheduleConfig {
	if bsc == nil {
		return nil
	}

	return &params.BlobScheduleConfig{
		Cancun: bsc.Cancun.ToEthereumBlobConfig(),
		Prague: bsc.Prague.ToEthereumBlobConfig(),
		Osaka:  bsc.Osaka.ToEthereumBlobConfig(),
		Verkle: bsc.Verkle.ToEthereumBlobConfig(),
	}
}

func FromEthereumBlobConfig(ebc *params.BlobConfig) *BlobConfig {
	if ebc == nil {
		return nil
	}

	return &BlobConfig{
		Target:                uint64(ebc.Target),
		Max:                   uint64(ebc.Max),
		BaseFeeUpdateFraction: ebc.UpdateFraction,
	}
}

func (bc *BlobConfig) ToEthereumBlobConfig() *params.BlobConfig {
	if bc == nil {
		return nil
	}

	return &params.BlobConfig{
		Target:         int(bc.Target),
		Max:            int(bc.Max),
		UpdateFraction: bc.BaseFeeUpdateFraction,
	}
}
