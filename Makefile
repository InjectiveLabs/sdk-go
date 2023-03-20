all:

copy-exchange-client:
	rm -rf exchange/*
	mkdir -p exchange/meta_rpc
	mkdir -p exchange/exchange_rpc
	mkdir -p exchange/accounts_rpc
	mkdir -p exchange/auction_rpc
	mkdir -p exchange/oracle_rpc
	mkdir -p exchange/insurance_rpc
	mkdir -p exchange/explorer_rpc
	mkdir -p exchange/spot_exchange_rpc
	mkdir -p exchange/derivative_exchange_rpc
	mkdir -p exchange/portfolio_rpc

	cp -r ../injective-indexer/api/gen/grpc/injective_meta_rpc/pb exchange/meta_rpc/pb
	cp -r ../injective-indexer/api/gen/grpc/injective_exchange_rpc/pb exchange/exchange_rpc/pb
	cp -r ../injective-indexer/api/gen/grpc/injective_accounts_rpc/pb exchange/accounts_rpc/pb
	cp -r ../injective-indexer/api/gen/grpc/injective_auction_rpc/pb exchange/auction_rpc/pb
	cp -r ../injective-indexer/api/gen/grpc/injective_oracle_rpc/pb exchange/oracle_rpc/pb
	cp -r ../injective-indexer/api/gen/grpc/injective_insurance_rpc/pb exchange/insurance_rpc/pb
	cp -r ../injective-indexer/api/gen/grpc/injective_explorer_rpc/pb exchange/explorer_rpc/pb
	cp -r ../injective-indexer/api/gen/grpc/injective_spot_exchange_rpc/pb exchange/spot_exchange_rpc/pb
	cp -r ../injective-indexer/api/gen/grpc/injective_derivative_exchange_rpc/pb exchange/derivative_exchange_rpc/pb
	cp -r ../injective-indexer/api/gen/grpc/injective_portfolio_rpc/pb exchange/portfolio_rpc/pb

.PHONY: copy-exchange-client

copy-chain-types:
	cp ../injective-core/injective-chain/modules/auction/types/*.go chain/auction/types
	rm -rf chain/auction/types/*test.go  rm -rf chain/auction/types/*gw.go
	cp ../injective-core/injective-chain/modules/exchange/types/*.go chain/exchange/types
	rm -rf chain/exchange/types/*test.go  rm -rf chain/exchange/types/*gw.go
	cp ../injective-core/injective-chain/modules/insurance/types/*.go chain/insurance/types
	rm -rf chain/insurance/types/*test.go  rm -rf chain/insurance/types/*gw.go
	cp ../injective-core/injective-chain/modules/ocr/types/*.go chain/ocr/types
	rm -rf chain/ocr/types/*test.go  rm -rf chain/ocr/types/*gw.go
	cp ../injective-core/injective-chain/modules/oracle/types/*.go chain/oracle/types
	rm -rf chain/oracle/types/*test.go  rm -rf chain/oracle/types/*gw.go
	cp ../injective-core/injective-chain/modules/peggy/types/*.go chain/peggy/types
	rm -rf chain/peggy/types/*test.go  rm -rf chain/peggy/types/*gw.go
	cp ../injective-core/injective-chain/modules/wasmx/types/*.go chain/wasmx/types
	rm -rf chain/wasmx/types/*test.go  rm -rf chain/wasmx/types/*gw.go
	cp ../injective-core/injective-chain/modules/tokenfactory/types/*.go chain/tokenfactory/types
	rm -rf chain/tokenfactory/types/*test.go  rm -rf chain/tokenfactory/types/*gw.go

	echo "👉 Replace injective-core/injective-chain/modules with sdk-go/chain"
	echo "👉 Replace injective-core/injective-chain/types with sdk-go/chain/types"
