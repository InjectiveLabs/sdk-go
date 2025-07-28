all:

clone-injective-indexer:
	git clone https://github.com/InjectiveLabs/injective-indexer.git -b v1.16.54 --depth 1 --single-branch

clone-injective-core:
	git clone https://github.com/InjectiveLabs/injective-core.git -b v1.16.0 --depth 1 --single-branch

copy-exchange-client: clone-injective-indexer
	rm -rf exchange/*
	mkdir -p exchange/event_provider_api/pb
	mkdir -p exchange/health_rpc/pb
	mkdir -p exchange/accounts_rpc/pb
	mkdir -p exchange/auction_rpc/pb
	mkdir -p exchange/campaign_rpc/pb
	mkdir -p exchange/derivative_exchange_rpc/pb
	mkdir -p exchange/exchange_rpc/pb
	mkdir -p exchange/explorer_rpc/pb
	mkdir -p exchange/insurance_rpc/pb
	mkdir -p exchange/meta_rpc/pb
	mkdir -p exchange/oracle_rpc/pb
	mkdir -p exchange/portfolio_rpc/pb
	mkdir -p exchange/spot_exchange_rpc/pb
	mkdir -p exchange/trading_rpc/pb

	cp -r injective-indexer/api/gen/grpc/event_provider_api/pb/*.pb.go exchange/event_provider_api/pb
	cp -r injective-indexer/api/gen/grpc/health/pb/*.pb.go exchange/health_rpc/pb
	cp -r injective-indexer/api/gen/grpc/injective_accounts_rpc/pb/*.pb.go exchange/accounts_rpc/pb
	cp -r injective-indexer/api/gen/grpc/injective_auction_rpc/pb/*.pb.go exchange/auction_rpc/pb
	cp -r injective-indexer/api/gen/grpc/injective_campaign_rpc/pb/*.pb.go exchange/campaign_rpc/pb
	cp -r injective-indexer/api/gen/grpc/injective_derivative_exchange_rpc/pb/*.pb.go exchange/derivative_exchange_rpc/pb
	cp -r injective-indexer/api/gen/grpc/injective_exchange_rpc/pb/*.pb.go exchange/exchange_rpc/pb
	cp -r injective-indexer/api/gen/grpc/injective_explorer_rpc/pb/*.pb.go exchange/explorer_rpc/pb
	cp -r injective-indexer/api/gen/grpc/injective_insurance_rpc/pb/*.pb.go exchange/insurance_rpc/pb
	cp -r injective-indexer/api/gen/grpc/injective_meta_rpc/pb/*.pb.go exchange/meta_rpc/pb
	cp -r injective-indexer/api/gen/grpc/injective_oracle_rpc/pb/*.pb.go exchange/oracle_rpc/pb
	cp -r injective-indexer/api/gen/grpc/injective_portfolio_rpc/pb/*.pb.go exchange/portfolio_rpc/pb
	cp -r injective-indexer/api/gen/grpc/injective_spot_exchange_rpc/pb/*.pb.go exchange/spot_exchange_rpc/pb
	cp -r injective-indexer/api/gen/grpc/injective_trading_rpc/pb/*.pb.go exchange/trading_rpc/pb

	rm -rf injective-indexer

copy-chain-types: clone-injective-core
	cp -r injective-core/injective-chain/codec chain
	mkdir -p chain/crypto/codec && cp injective-core/injective-chain/crypto/codec/*.go chain/crypto/codec
	rm -rf chain/crypto/codec/*test.go rm -rf chain/crypto/codec/*gw.go
	mkdir -p chain/crypto/ethsecp256k1 && cp injective-core/injective-chain/crypto/ethsecp256k1/*.go chain/crypto/ethsecp256k1
	rm -rf chain/crypto/ethsecp256k1/*test.go rm -rf chain/crypto/ethsecp256k1/*gw.go
	mkdir -p chain/crypto/hd && cp injective-core/injective-chain/crypto/hd/*.go chain/crypto/hd
	rm -rf chain/crypto/hd/*test.go rm -rf chain/crypto/hd/*gw.go
	mkdir -p chain/auction/types && \
		cp injective-core/injective-chain/modules/auction/types/*.pb.go chain/auction/types && \
		cp injective-core/injective-chain/modules/auction/types/codec.go chain/auction/types
	mkdir -p chain/erc20/types && \
		cp injective-core/injective-chain/modules/erc20/types/*.pb.go chain/erc20/types && \
		cp injective-core/injective-chain/modules/erc20/types/codec.go chain/erc20/types
	mkdir -p chain/evm/types && \
		cp injective-core/injective-chain/modules/evm/types/*.pb.go chain/evm/types && \
		cp injective-core/injective-chain/modules/evm/types/access_list.go chain/evm/types && \
		cp injective-core/injective-chain/modules/evm/types/access_list_tx.go chain/evm/types && \
		cp injective-core/injective-chain/modules/evm/types/chain_config.go chain/evm/types && \
		cp injective-core/injective-chain/modules/evm/types/codec.go chain/evm/types && \
		cp injective-core/injective-chain/modules/evm/types/dynamic_fee_tx.go chain/evm/types && \
		cp injective-core/injective-chain/modules/evm/types/errors.go chain/evm/types && \
		cp injective-core/injective-chain/modules/evm/types/eth.go chain/evm/types && \
		cp injective-core/injective-chain/modules/evm/types/events.go chain/evm/types && \
		cp injective-core/injective-chain/modules/evm/types/key.go chain/evm/types && \
		cp injective-core/injective-chain/modules/evm/types/legacy_tx.go chain/evm/types && \
		cp injective-core/injective-chain/modules/evm/types/logs.go chain/evm/types && \
		cp injective-core/injective-chain/modules/evm/types/msg.go chain/evm/types && \
		cp injective-core/injective-chain/modules/evm/types/params.go chain/evm/types && \
		cp injective-core/injective-chain/modules/evm/types/storage.go chain/evm/types && \
		cp injective-core/injective-chain/modules/evm/types/tx.go chain/evm/types && \
		cp injective-core/injective-chain/modules/evm/types/tx_data.go chain/evm/types && \
		cp injective-core/injective-chain/modules/evm/types/utils.go chain/evm/types
	mkdir -p chain/exchange/types && \
		cp injective-core/injective-chain/modules/exchange/types/*.go chain/exchange/types && \
		rm -rf chain/exchange/types/*test.go && rm -rf chain/exchange/types/*gw.go
	mkdir -p chain/exchange/types/v2 && \
    		cp injective-core/injective-chain/modules/exchange/types/v2/*.go chain/exchange/types/v2 && \
    		rm -rf chain/exchange/types/v2/*test.go && rm -rf chain/exchange/types/v2/*gw.go
	mkdir -p chain/insurance/types && \
		cp injective-core/injective-chain/modules/insurance/types/*.pb.go chain/insurance/types && \
		cp injective-core/injective-chain/modules/insurance/types/codec.go chain/insurance/types
	mkdir -p chain/ocr/types && \
		cp injective-core/injective-chain/modules/ocr/types/*.pb.go chain/ocr/types && \
		cp injective-core/injective-chain/modules/ocr/types/errors.go chain/ocr/types && \
		cp injective-core/injective-chain/modules/ocr/types/key.go chain/ocr/types && \
		cp injective-core/injective-chain/modules/ocr/types/params.go chain/ocr/types && \
		cp injective-core/injective-chain/modules/ocr/types/proposal.go chain/ocr/types && \
		cp injective-core/injective-chain/modules/ocr/types/types.go chain/ocr/types && \
		cp injective-core/injective-chain/modules/ocr/types/codec.go chain/ocr/types
	mkdir -p chain/oracle/types && \
		cp injective-core/injective-chain/modules/oracle/types/*.pb.go chain/oracle/types && \
		cp injective-core/injective-chain/modules/oracle/types/codec.go chain/oracle/types && \
		cp injective-core/injective-chain/modules/oracle/types/errors.go chain/oracle/types && \
		cp injective-core/injective-chain/modules/oracle/types/msgs.go chain/oracle/types && \
		cp injective-core/injective-chain/modules/oracle/types/oracle.go chain/oracle/types && \
		cp injective-core/injective-chain/modules/oracle/types/params.go chain/oracle/types && \
		cp injective-core/injective-chain/modules/oracle/types/proposal.go chain/oracle/types && \
		cp injective-core/injective-chain/modules/oracle/types/stork_oracle.go chain/oracle/types && \
		cp -r injective-core/injective-chain/modules/oracle/bandchain chain/oracle
	mkdir -p chain/peggy/types && \
		cp injective-core/injective-chain/modules/peggy/types/*.pb.go chain/peggy/types && \
		cp injective-core/injective-chain/modules/peggy/types/abi_json.go chain/peggy/types && \
		cp injective-core/injective-chain/modules/peggy/types/codec.go chain/peggy/types && \
		cp injective-core/injective-chain/modules/peggy/types/ethereum.go chain/peggy/types && \
		cp injective-core/injective-chain/modules/peggy/types/ethereum_signer.go chain/peggy/types && \
		cp injective-core/injective-chain/modules/peggy/types/errors.go chain/peggy/types && \
		cp injective-core/injective-chain/modules/peggy/types/key.go chain/peggy/types && \
		cp injective-core/injective-chain/modules/peggy/types/msgs.go chain/peggy/types && \
		cp injective-core/injective-chain/modules/peggy/types/params.go chain/peggy/types && \
		cp injective-core/injective-chain/modules/peggy/types/types.go chain/peggy/types
	mkdir -p chain/permissions/types && \
		cp injective-core/injective-chain/modules/permissions/types/*.pb.go chain/permissions/types && \
		cp injective-core/injective-chain/modules/permissions/types/codec.go chain/permissions/types
	mkdir -p chain/tokenfactory/types && \
		cp injective-core/injective-chain/modules/tokenfactory/types/*.pb.go chain/tokenfactory/types && \
		cp injective-core/injective-chain/modules/tokenfactory/types/codec.go chain/tokenfactory/types
	mkdir -p chain/txfees/types && \
		cp injective-core/injective-chain/modules/txfees/types/*.pb.go chain/txfees/types && \
		cp injective-core/injective-chain/modules/txfees/types/codec.go chain/txfees/types
	mkdir -p chain/txfees/osmosis/types && \
		cp injective-core/injective-chain/modules/txfees/osmosis/types/*.pb.go chain/txfees/osmosis/types
	mkdir -p chain/wasmx/types && \
		cp injective-core/injective-chain/modules/wasmx/types/*.pb.go chain/wasmx/types && \
		cp injective-core/injective-chain/modules/wasmx/types/authz.go chain/wasmx/types && \
		cp injective-core/injective-chain/modules/wasmx/types/codec.go chain/wasmx/types && \
		cp injective-core/injective-chain/modules/wasmx/types/custom_execution.go chain/wasmx/types && \
		cp injective-core/injective-chain/modules/wasmx/types/errors.go chain/wasmx/types && \
		cp injective-core/injective-chain/modules/wasmx/types/key.go chain/wasmx/types && \
		cp injective-core/injective-chain/modules/wasmx/types/msgs.go chain/wasmx/types && \
		cp injective-core/injective-chain/modules/wasmx/types/params.go chain/wasmx/types && \
		cp injective-core/injective-chain/modules/wasmx/types/proposal.go chain/wasmx/types
	mkdir -p chain/stream/types && \
		cp injective-core/injective-chain/stream/types/*.pb.go chain/stream/types
	mkdir -p chain/stream/types/v2 && \
    		cp injective-core/injective-chain/stream/types/v2/*.pb.go chain/stream/types/v2
	mkdir -p chain/types && \
		cp injective-core/injective-chain/types/*.pb.go injective-core/injective-chain/types/config.go chain/types && \
		cp injective-core/injective-chain/types/chain_id.go chain/types && \
		cp injective-core/injective-chain/types/codec.go chain/types && \
		cp injective-core/injective-chain/types/errors.go chain/types && \
		cp injective-core/injective-chain/types/int.go chain/types && \
		cp injective-core/injective-chain/types/util.go chain/types && \
		cp injective-core/injective-chain/types/validation.go chain/types

	@find ./chain -type f -name "*.go" -exec sed -i "" -e "s|github.com/InjectiveLabs/injective-core/injective-chain/modules|github.com/InjectiveLabs/sdk-go/chain|g" {} \;
	@find ./chain -type f -name "*.go" -exec sed -i "" -e "s|github.com/InjectiveLabs/injective-core/injective-chain|github.com/InjectiveLabs/sdk-go/chain|g" {} \;

	mkdir -p chain/evm/precompiles/bank && mkdir -p chain/evm/precompiles/exchange && mkdir -p chain/evm/precompiles/staking && \
		cp injective-core/injective-chain/modules/evm/precompiles/bindings/cosmos/precompile/bank/*.go chain/evm/precompiles/bank && \
		cp injective-core/injective-chain/modules/evm/precompiles/bindings/cosmos/precompile/exchange/*.go chain/evm/precompiles/exchange && \
		cp injective-core/injective-chain/modules/evm/precompiles/bindings/cosmos/precompile/staking/*.go chain/evm/precompiles/staking

	rm -rf proto
	cp -r injective-core/proto ./

	rm -rf injective-core
	make extract-message-names

extract-message-names:
	@echo "Extracting message names from tx.pb.go files..."
	@mkdir -p injective_data
	@find ./chain -name "tx.pb.go" -exec grep -h "proto\.RegisterType" {} \; | \
		sed -n 's/.*proto\.RegisterType([^"]*"\([^"]*\)".*/\1/p' | \
		grep -v 'Response$$' | \
		sort -u | \
		jq -R -s 'split("\n")[:-1]' > injective_data/chain_messages_list.json
	@echo "Message names extracted to injective_data/chain_messages_list.json (excluding Response messages)"
	@echo "Total messages found: $$(jq length injective_data/chain_messages_list.json)"

#gen: gen-proto
#
#gen-proto: clone-all copy-proto
#	buf generate --template buf.gen.chain.yaml
#	buf generate --template buf.gen.indexer.yaml
#	rm -rf local_proto
#	$(call clean_repos)
#
#define clean_repos
#	rm -Rf injective-indexer
#endef
#
#clean-all:
#	$(call clean_repos)
#
#clone-injective-indexer:
#	git clone https://github.com/InjectiveLabs/injective-indexer.git -b v1.13.4 --depth 1 --single-branch
#
#clone-all: clone-injective-indexer
#
#copy-proto:
#	rm -rf local_proto
#	mkdir -p local_proto
#	find ./injective-indexer/api/gen/grpc -type f -name "*.proto" | while read -r file; do \
#		dest="local_proto/$$(basename $$(dirname $$(dirname "$$file")))/$$(basename $$(dirname "$$file"))"; \
#		mkdir -p "$$dest"; \
#		cp "$$file" "$$dest"; \
#	done

tests:
	go test -race ./client/... ./ethereum/...
coverage:
	go test -race -coverprofile=coverage.out -covermode=atomic ./client/... ./ethereum/...

lint: export GOPROXY=direct
lint:
	golangci-lint run --timeout=15m -v --new-from-rev=dev

lint-last-commit: export GOPROXY=direct
lint-last-commit:
	golangci-lint run --timeout=15m -v --new-from-rev=HEAD~

lint-master: export GOPROXY=direct
lint-master:
	golangci-lint run --timeout=15m -v --new-from-rev=master

lint-all: export GOPROXY=direct
lint-all:
	golangci-lint run --timeout=15m -v

.PHONY: copy-exchange-client tests coverage lint lint-last-commit lint-master lint-all extract-message-names
