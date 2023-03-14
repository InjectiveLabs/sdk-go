package common

import (
	ctypes "github.com/InjectiveLabs/sdk-go/chain/types"
	log "github.com/InjectiveLabs/suplog"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"google.golang.org/grpc/credentials"
	"time"
)

func init() {
	// set the address prefixes
	config := sdk.GetConfig()

	// This is specific to Injective chain
	ctypes.SetBech32Prefixes(config)
	ctypes.SetBip44CoinType(config)
}

type ClientOptions struct {
	GasPrices             string
	GasPricesIncrement    sdk.Dec
	GasPriceResetInterval *time.Ticker
	TLSCert               credentials.TransportCredentials
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

func OptionGasPriceIncrement(percent int64) ClientOption {
	return func(opts *ClientOptions) error {
		opts.GasPricesIncrement = sdk.NewDecWithPrec(100+percent, 2)
		return nil
	}
}

func OptionsGasPriceResetInterval(ticker *time.Ticker) ClientOption {
	return func(opts *ClientOptions) error {
		if ticker != nil {
			opts.GasPriceResetInterval = ticker
		} else {
			opts.GasPriceResetInterval = time.NewTicker(time.Minute * 5)
		}
		return nil
	}
}

func OptionTLSCert(tlsCert credentials.TransportCredentials) ClientOption {
	return func(opts *ClientOptions) error {
		if tlsCert == nil {
			log.Infoln("client does not use grpc secure transport")
		} else {
			log.Infoln("succesfully load server TLS cert")
		}
		opts.TLSCert = tlsCert
		return nil
	}
}
