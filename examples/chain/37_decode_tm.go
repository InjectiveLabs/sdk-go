package main

import (
	"encoding/base64"
	"fmt"
	cosmostxpb "github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

var parserMap = map[string]interface{}{
	"/cosmos.bank.v1beta1.MsgSend":      banktypes.MsgSend{},
	"/cosmos.bank.v1beta1.MsgMultiSend": banktypes.MsgMultiSend{},
}

func main() {
	//multisend := "Cr0BCroBCiEvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dNdWx0aVNlbmQSlAEKSAoqaW5qMTd2eXRkd3FjenF6NzJqNjVzYXVrcGxya3RkNGd5Zm1lNWFnZjZjEhoKA2luahITMTAwMDAwMDAwMDAwMDAwMDAwMBJICippbmoxN3Z5dGR3cWN6cXo3Mmo2NXNhdWtwbHJrdGQ0Z3lmbWU1YWdmNmMSGgoDaW5qEhMxMDAwMDAwMDAwMDAwMDAwMDAwEn8KYApUCi0vaW5qZWN0aXZlLmNyeXB0by52MWJldGExLmV0aHNlY3AyNTZrMS5QdWJLZXkSIwohA5Bh/pUwQH5Sgsniw2eFD5lZHswzGUxzuaaH8g0xKzO2EgQKAggBGPm7CBIbChUKA2luahIONjAzMDQ1MDAwMDAwMDAQoa4HGkFuew+TP4HtsHcXFtFcg33d5QeJZCLmT3glrpgbY+NKbCQ9IGiIzMke1kql9DTEdKFqyMPXvfUQ4bUTXq2tCV7HAA=="
	send := "CtQBCpQBChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEnQKKmluajFybGU4eXJ5bmx0cnVtNG1uN2VmbjBtNTJmNXFzajNobTM4eHdydxIqaW5qMWNwc3hldTNzczdyODNtM2EwYTM2YzJzM3hqdTR3NDJyM3NhMjh2GhoKA2luahITNTQ5OTAwMDAwMDAwMDAwMDAwMBjLvYwM+j81Ci8vaW5qZWN0aXZlLnR5cGVzLnYxYmV0YTEuRXh0ZW5zaW9uT3B0aW9uc1dlYjNUeBICCAESfgpeClQKLS9pbmplY3RpdmUuY3J5cHRvLnYxYmV0YTEuZXRoc2VjcDI1NmsxLlB1YktleRIjCiECAVMpGRBIgvdIfgYwsSJyaE3W1+j1FIcVtXdaNsj+qbISBAoCCH8YBBIcChYKA2luahIPMjAwMDAwMDAwMDAwMDAwEIC1GBpBSpQwr6FfXa7KWLz4Eousx1VrPuVMHgMAUHkRVeNK/UpFQztyTAeISB6vjl3Dx2sRKWkqy/hQ7G3lICoajar1HRw="
	bytes, _ := base64.StdEncoding.DecodeString(send)
	rawTx := cosmostxpb.TxRaw{}
	rawTx.Unmarshal(bytes)

	txBody := cosmostxpb.TxBody{}
	txBody.Unmarshal(rawTx.BodyBytes)

	for _, msg := range txBody.Messages {
		switch parserMap[msg.TypeUrl].(type) {
		case banktypes.MsgSend:
			var result banktypes.MsgSend
			result.XXX_Unmarshal(msg.Value)
			fmt.Println(result.FromAddress)
			fmt.Println(result.ToAddress)
			fmt.Println(result.Amount)

		case banktypes.MsgMultiSend:
			var result banktypes.MsgMultiSend
			result.XXX_Unmarshal(msg.Value)
			inputs := result.Inputs
			outputs := result.Outputs

			fmt.Println(inputs)
			fmt.Println(outputs)

		default:
			fmt.Println("Unexpected Type:", msg.TypeUrl)
		}
	}
}
