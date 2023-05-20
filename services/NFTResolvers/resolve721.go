package resolver

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/chokey2nv/ens-resolve/config"
	"github.com/chokey2nv/ens-resolve/services/go_gen/ERC721"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func Resolve721(config *config.Config, address common.Address) ([]interface{}, error) {
	client := config.Client

	transferSignature := crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))

	walletAddress := common.BytesToHash(
		common.LeftPadBytes(address.Hash().Bytes(), 32),
	)
	query := &ethereum.FilterQuery{
		Topics: [][]common.Hash{
			{transferSignature},
			nil,
			{walletAddress},
		},
	}

	logs, err := client.FilterLogs(context.Background(), *query)
	if err != nil {
		return nil, err
	}
	var nftMetadata []interface{}
	for _, l := range logs {
		tokenId := l.Topics[3]
		erc721, err := ERC721.NewERC721(l.Address, client)
		if err != nil {
			panic(err)
		}
		data, _ := erc721.TokenURI(nil, tokenId.Big())
		if data != "" {
			nData, err := ResolveNFTUrl(data, tokenId.Bytes())
			if err != nil {
				return nil, err
			}
			nftMetadata = append(nftMetadata, nData)
		}
	}
	return nftMetadata, nil
}

func ResolveNFTUrl(url string, tokenIdBytes []byte) (interface{}, error) {
	if strings.Contains(url, "data:application/json") {
		url = strings.Split(url, ",")[1]
		return DecodeBase6(url)
	}
	tokenId := common.Bytes2Hex(tokenIdBytes)
	// if strings.Contains(url, "api.opensea.io") {
	// tokenId = new(big.Int).SetBytes(tokenIdBytes).String()
	// }
	url = strings.Replace(url, "ipfs://", "https://ipfs.io/ipfs/", 1)
	url = strings.Replace(url, "{id}", tokenId, 1)

	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var nftData interface{}
	json.Unmarshal(body, &nftData)
	if d, ok := nftData.(map[string]any); ok {
		// Update the `image` property
		var img string
		if img, ok = d["image"].(string); ok {
			img = strings.Replace(img, "ipfs://", "https://ipfs.io/ipfs/", 1)
		}
		d["image"] = img
		// Reassign the updated concrete type back to the interface
		nftData = d
	}
	return nftData, nil
}

func DecodeBase6(base64Str string) (interface{}, error) {

	decoded, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		log.Fatal("Error decoding base64 string:", err)
	}

	var data interface{}
	err = json.Unmarshal(decoded, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
