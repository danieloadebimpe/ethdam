package resolver

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/chokey2nv/ens-resolve/config"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func FilterNFT(config *config.Config) {
	client := config.Client

	// TransferSingle := crypto.Keccak256Hash([]byte("TransferSingle(address,address,address,uint256,uint256,bytes)"))
	// TransferBatch := crypto.Keccak256Hash([]byte("TransferBatch(address,address,address,uint256[],uint256[],bytes)"))
	// SafeTransferFrom := crypto.Keccak256Hash([]byte("safeTransferFrom(address,address,uint256,uint256,bytes)"))
	// SafeBatchTransferFrom := crypto.Keccak256Hash([]byte("safeBatchTransferFrom(address,address,uint256[],uint256[],bytes)"))
	transferSignature := crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
	// approvalSignature := crypto.Keccak256Hash([]byte("Approval(address,address,uint256)"))
	// approvalForAllSignature := crypto.Keccak256Hash([]byte("ApprovalForAll(address,address,bool)"))
	// fmt.Println(transferSignature, approvalForAllSignature, approvalSignature)

	// block, _ := client.BlockNumber(context.Background())

	// FilterQuery to get all ERC721 contract addresses
	walletAddress := common.BytesToHash(
		common.LeftPadBytes(common.HexToHash(
			strings.ToLower("0x88e4519e2Baa513Ed92B0Ae4c788D7E5c5B03Ea4"),
		).Bytes(), 32),
	)
	fmt.Println(walletAddress, transferSignature)
	query := &ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(15498629)),
		ToBlock:   big.NewInt(int64(15498629)),
		// Addresses: []common.Address{common.HexToAddress("0x6761dc899f300a16d6070fe1139a4ee037b057d3")},
		Topics: [][]common.Hash{
			{transferSignature},
			nil,
			{walletAddress},
		},
	}

	logs, err := client.FilterLogs(context.Background(), *query)
	if err != nil {
		log.Fatal(err)
	}

	for _, l := range logs {
		fmt.Println("==================================")
		fmt.Printf("Address: %s\n", l.Address.Hex())
		fmt.Printf("TxHash: %s\n", l.TxHash)
		fmt.Printf("Topics: %v\n", l.Topics)
		fmt.Printf("Data: %v\n", l.BlockNumber)
		fmt.Println("==================================")
		
	}
}
