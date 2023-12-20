package common

import (
	ctypes "github.com/InjectiveLabs/sdk-go/chain/types"
	log "github.com/InjectiveLabs/suplog"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"google.golang.org/grpc/credentials"
)

func init() {
	// set the address prefixes
	config := sdk.GetConfig()

	// This is specific to Injective chain
	ctypes.SetBech32Prefixes(config)
	ctypes.SetBip44CoinType(config)
}

type ClientOptions struct {
	GasPrices string
	TLSCert   credentials.TransportCredentials
	TxFactory *tx.Factory
}

type ClientOption func(opts *ClientOptions) error

func DefaultClientOptions() *ClientOptions {
	return &ClientOptions{}
}

func OptionGasPrices(gasPrices string) ClientOption {
	return func(opts *ClientOptions) error {
		_, err := sdk.ParseDecCoins(gasPrices)
		if err != nil {
			err = errors.Wrapf(err, "failed to ParseDecCoins %s", gasPrices)
			return err
		}

		opts.GasPrices = gasPrices
		return nil
	}
}

func OptionTLSCert(tlsCert credentials.TransportCredentials) ClientOption {
	return func(opts *ClientOptions) error {
		if tlsCert == nil {
			log.Infoln("client does not use grpc secure transport")
		} else {
			log.Infoln("successfully load server TLS cert")
		}
		opts.TLSCert = tlsCert
		return nil
	}
}

func OptionTxFactory(txFactory *tx.Factory) ClientOption {
	return func(opts *ClientOptions) error {
		opts.TxFactory = txFactory
		return nil
	}
}
