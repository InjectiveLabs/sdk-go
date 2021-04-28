all:

copy-exchange-client:
	mkdir -p exchange/exchange_rpc
	mkdir -p exchange/accounts_rpc
	mkdir -p exchange/spot_exchange_rpc
	mkdir -p exchange/derivative_exchange_rpc

	cp -r ../injective-exchange/api/gen/grpc/injective_exchange_rpc/pb exchange/exchange_rpc/pb
	cp -r ../injective-exchange/api/gen/grpc/injective_accounts_rpc/pb exchange/accounts_rpc/pb
	cp -r ../injective-exchange/api/gen/grpc/injective_spot_exchange_rpc/pb exchange/spot_exchange_rpc/pb
	cp -r ../injective-exchange/api/gen/grpc/injective_derivative_exchange_rpc/pb exchange/derivative_exchange_rpc/pb

.PHONY: copy-exchange-client
