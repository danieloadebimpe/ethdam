package resolver

import (
	"github.com/chokey2nv/ens-resolve/config"
	"github.com/wealdtech/go-ens/v3"
)

type ResolvedName struct {
	Address string `json:"address"`
	Twitter string `json:"twitter"`
	Github  string `json:"github"`
	Discord string `json:"discord"`
	Email   string `json:"email"`
}

func Resolve(config *config.Config, domain string) (*ResolvedName, error) {
	// Connect to the Ethereum network
	client := config.Client
	// domain := config.Domain1

	// resolvedAddr, _ := ens.Resolve(client, domain)
	// Set up the ENS client
	record, err := ens.NewResolver(client, domain)
	if err != nil {
		panic(err)
	}

	addr, _ := record.Address()

	email, err := record.Text("email")
	if err != nil {
		return nil, err
	}
	twitter, err := record.Text("com.twitter")
	if err != nil {
		return nil, err
	}
	github, err := record.Text("com.github")
	if err != nil {
		return nil, err
	}
	discord, err := record.Text("com.discord")
	if err != nil {
		return nil, err
	}
	return &ResolvedName{
		Email:   email,
		Twitter: twitter,
		Github:  github,
		Discord: discord,
		Address: addr.Hex(),
	}, nil
}
