all:

copy-exchange-client:
	cp -r ../injective-exchange/api/gen/grpc/injective_exchange_rpc exchange/exchange_rpc
	cp -r ../injective-exchange/api/gen/grpc/injective_accounts_rpc exchange/accounts_rpc
	cp -r ../injective-exchange/api/gen/grpc/injective_spot_exchange_rpc exchange/spot_exchange_rpc
	cp -r ../injective-exchange/api/gen/grpc/injective_derivative_exchange_rpc exchange/derivative_exchange_rpc
	rm -rf exchange/exchange_rpc/server
	rm -rf exchange/accounts_rpc/server
	rm -rf exchange/spot_exchange_rpc/server
	rm -rf exchange/derivative_exchange_rpc/server

.PHONY: copy-exchange-client
