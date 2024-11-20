all:

copy-exchange-client:
	rm -rf exchange/*
	mkdir -p exchange/health_rpc
	mkdir -p exchange/accounts_rpc
	mkdir -p exchange/auction_rpc
	mkdir -p exchange/campaign_rpc
	mkdir -p exchange/derivative_exchange_rpc
	mkdir -p exchange/exchange_rpc
	mkdir -p exchange/explorer_rpc
	mkdir -p exchange/insurance_rpc
	mkdir -p exchange/meta_rpc
	mkdir -p exchange/oracle_rpc
	mkdir -p exchange/portfolio_rpc
	mkdir -p exchange/spot_exchange_rpc
	mkdir -p exchange/trading_rpc

	cp -r ../injective-indexer/api/gen/grpc/health/pb exchange/health_rpc/pb
	cp -r ../injective-indexer/api/gen/grpc/injective_accounts_rpc/pb exchange/accounts_rpc/pb
	cp -r ../injective-indexer/api/gen/grpc/injective_accounts_rpc/pb exchange/accounts_rpc/pb
	cp -r ../injective-indexer/api/gen/grpc/injective_auction_rpc/pb exchange/auction_rpc/pb
	cp -r ../injective-indexer/api/gen/grpc/injective_campaign_rpc/pb exchange/campaign_rpc/pb
	cp -r ../injective-indexer/api/gen/grpc/injective_derivative_exchange_rpc/pb exchange/derivative_exchange_rpc/pb
	cp -r ../injective-indexer/api/gen/grpc/injective_exchange_rpc/pb exchange/exchange_rpc/pb
	cp -r ../injective-indexer/api/gen/grpc/injective_explorer_rpc/pb exchange/explorer_rpc/pb
	cp -r ../injective-indexer/api/gen/grpc/injective_insurance_rpc/pb exchange/insurance_rpc/pb
	cp -r ../injective-indexer/api/gen/grpc/injective_meta_rpc/pb exchange/meta_rpc/pb
	cp -r ../injective-indexer/api/gen/grpc/injective_oracle_rpc/pb exchange/oracle_rpc/pb
	cp -r ../injective-indexer/api/gen/grpc/injective_portfolio_rpc/pb exchange/portfolio_rpc/pb
	cp -r ../injective-indexer/api/gen/grpc/injective_spot_exchange_rpc/pb exchange/spot_exchange_rpc/pb
	cp -r ../injective-indexer/api/gen/grpc/injective_trading_rpc/pb exchange/trading_rpc/pb

.PHONY: copy-exchange-client tests coverage

copy-chain-types:
	cp ../injective-core/injective-chain/crypto/ethsecp256k1/*.go chain/crypto/ethsecp256k1
	rm -rf chain/crypto/ethsecp256k1/*test.go rm -rf chain/crypto/ethsecp256k1/*gw.go
	cp ../injective-core/injective-chain/codec/types/*.go chain/codec/types
	rm -rf chain/codec/types/*test.go rm -rf chain/codec/types/*gw.go
	cp ../injective-core/injective-chain/modules/auction/types/*.go chain/auction/types
	rm -rf chain/auction/types/*test.go  rm -rf chain/auction/types/*gw.go
	cp ../injective-core/injective-chain/modules/exchange/types/*.go chain/exchange/types
	rm -rf chain/exchange/types/*test.go  rm -rf chain/exchange/types/*gw.go
	cp ../injective-core/injective-chain/modules/exchange/types/v2/*.go chain/exchange/types/v2
	rm -rf chain/exchange/types/v2/*test.go  rm -rf chain/exchange/types/v2/*gw.go
	cp ../injective-core/injective-chain/modules/insurance/types/*.go chain/insurance/types
	rm -rf chain/insurance/types/*test.go  rm -rf chain/insurance/types/*gw.go
	cp ../injective-core/injective-chain/modules/ocr/types/*.go chain/ocr/types
	rm -rf chain/ocr/types/*test.go  rm -rf chain/ocr/types/*gw.go
	cp ../injective-core/injective-chain/modules/oracle/types/*.go chain/oracle/types
	cp -r ../injective-core/injective-chain/modules/oracle/bandchain chain/oracle
	rm -rf chain/oracle/types/*test.go  rm -rf chain/oracle/types/*gw.go
	cp ../injective-core/injective-chain/modules/peggy/types/*.go chain/peggy/types
	rm -rf chain/peggy/types/*test.go  rm -rf chain/peggy/types/*gw.go
	cp ../injective-core/injective-chain/modules/permissions/types/*.go chain/permissions/types
	rm -rf chain/permissions/types/*test.go  rm -rf chain/permissions/types/*gw.go
	cp ../injective-core/injective-chain/modules/tokenfactory/types/*.go chain/tokenfactory/types
	rm -rf chain/tokenfactory/types/*test.go  rm -rf chain/tokenfactory/types/*gw.go
	cp ../injective-core/injective-chain/modules/wasmx/types/*.go chain/wasmx/types
	rm -rf chain/wasmx/types/*test.go  rm -rf chain/wasmx/types/*gw.go
	cp ../injective-core/injective-chain/stream/types/*.go chain/stream/types
	rm -rf chain/stream/types/*test.go  rm -rf chain/stream/types/*gw.go
	cp ../injective-core/injective-chain/stream/types/v2/*.go chain/stream/types/v2
	rm -rf chain/stream/types/v2/*test.go  rm -rf chain/stream/types/v2/*gw.go
	cp ../injective-core/injective-chain/types/*.go chain/types
	rm -rf chain/types/*test.go rm -rf chain/types/*gw.go

	@echo "👉 Replace injective-core/injective-chain/modules with sdk-go/chain"
	@echo "👉 Replace injective-core/injective-chain/codec with sdk-go/chain/codec"
	@echo "👉 Replace injective-core/injective-chain/codec/types with sdk-go/chain/codec/types"
	@echo "👉 Replace injective-core/injective-chain/types with sdk-go/chain/types"
	@echo "👉 Replace injective-core/injective-chain/crypto with sdk-go/chain/crypto"
	@echo "👉 Replace injective-core/injective-chain/stream/types with sdk-go/chain/stream/types"

tests:
	go test -race ./client/... ./ethereum/...
coverage:
	go test -race -coverprofile=coverage.out -covermode=atomic ./client/... ./ethereum/...
