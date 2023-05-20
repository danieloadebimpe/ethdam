package main

import (
	"github.com/chokey2nv/ens-resolve/config"
	"github.com/chokey2nv/ens-resolve/services"
)

func main() {
	config := config.AppConfig()
	services.Server(config)
}
