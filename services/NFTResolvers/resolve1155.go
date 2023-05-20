package resolver

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/chokey2nv/ens-resolve/config"
	"github.com/chokey2nv/ens-resolve/services/go_gen/ERC1155"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func Resolve1155(config *config.Config, address common.Address) ([]interface{}, error) {
	client := config.Client
	transferSingle := crypto.Keccak256Hash([]byte("TransferSingle(address,address,address,uint256,uint256)"))

	// blockNumber, _ := client.BlockNumber(context.Background())
	walletAddress := common.BytesToHash(
		common.LeftPadBytes(address.Hash().Bytes(), 32),
	)
	query := &ethereum.FilterQuery{
		// FromBlock: big.NewInt(int64(blockNumber - 2)),
		Topics: [][]common.Hash{
			{transferSingle},
			nil,
			nil,
			{walletAddress},
		},
	}

	logs, err := client.FilterLogs(context.Background(), *query)
	if err != nil {
		log.Fatal(err)
	}
	var nftData []interface{}
	for _, l := range logs {
		//split the byte array into two parts: the first 32 bytes for the token ID, and the last 32 bytes for the value
		tokenIdBytes := l.Data[:32]
		// valueBytes := data[32:]

		// convert the token ID byte array to a big.Int
		tokenBigInt := new(big.Int).SetBytes(tokenIdBytes)

		erc1155, err := ERC1155.NewERC1155(l.Address, client)
		if err != nil {
			return nil, err
		}
		uri, err := erc1155.Uri(nil, tokenBigInt)
		if err != nil {
			fmt.Println(err.Error())
		}
		if uri != "" {
			mData, err := ResolveNFTUrl(uri, tokenIdBytes)
			if err == nil {
				nftData = append(nftData, mData)
			} else {
				log.Fatalln(err)
			}
		}

	}
	return nftData, nil
}
