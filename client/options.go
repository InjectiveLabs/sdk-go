package client

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"
	"google.golang.org/grpc/credentials"
)

type clientOptions struct {
	GasPrices string
	TLSCert   credentials.TransportCredentials
}

type clientOption func(opts *clientOptions) error

func defaultClientOptions() *clientOptions {
	return &clientOptions{}
}

func OptionGasPrices(gasPrices string) clientOption {
	return func(opts *clientOptions) error {
		_, err := sdk.ParseDecCoins(gasPrices)
		if err != nil {
			err = errors.Wrapf(err, "failed to ParseDecCoins %s", gasPrices)
			return err
		}

		opts.GasPrices = gasPrices
		return nil
	}
}

func OptionTLSCert(tlsCert credentials.TransportCredentials) clientOption {
	return func(opts *clientOptions) error {
		if tlsCert == nil {
			log.Infoln("Client does not use grpc secure transport")
		} else {
			log.Infoln("Succesfully load server TLS cert")
		}
		opts.TLSCert = tlsCert
		return nil
	}
}