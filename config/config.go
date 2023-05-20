package config

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

//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
type Config struct {
	Client         *ethclient.Client
	Github         *GithubConfig
	OpenseaMetaApi string
}
type GithubConfig struct {
	ApiUrl   string
	AuthUrl  string
	ClientId string
	Secret   string
}

func AppConfig() *Config {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/bbad6dc9dece4ec1b9c0beae4aea0a9d")
	if err != nil {
		panic(err)
	}

	config := Config{
		Client: client,
		Github: &GithubConfig{
			ApiUrl:   "https://api.github.com/",
			AuthUrl:  "https://github.com/login/oauth/",
			Secret:   viper.GetString("GITHUB_CLIENT_SECRET"),
			ClientId: viper.GetString("GITHUB_CLIENT_ID"),
		},
		OpenseaMetaApi: "https://api.opensea.io/api/v1/asset/{contract}/{id}",
	}
	fmt.Println(config)
	fmt.Println(*config.Github)
	return &config
}
