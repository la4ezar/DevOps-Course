// Making requests to the server with custom client
package main

import (
	"github.com/la4ezar/devops/DevOps-Course/app/restapi/internal/config"
	"github.com/la4ezar/devops/DevOps-Course/app/restapi/internal/crypto"
	"github.com/la4ezar/devops/DevOps-Course/app/restapi/pkg/client"
	"github.com/la4ezar/devops/DevOps-Course/app/restapi/pkg/log"
)

func main() {
	cfg, err := config.NewDefaultClientConfig()
	if err != nil {
		log.D().WithError(err).Error()
	}

	if err := cfg.Validate(); err != nil {
		log.D().WithError(err).Error()
	}

	// Init Client
	c := client.New(cfg.Client)

	// Health check
	c.HealthCheck()

	// GET requests
	c.GetCryptos()
	c.GetCrypto("BTC")
	c.GetCrypto("LTC")

	//POST request
	cryptocurrency := client.Cryptocurrency{
		Name:     "LachoCoin",
		CryptoID: "LCN",
		Price:    1.54,
		Authors:  client.CryptoAuthors{crypto.Author{Firstname: "Lachezar", Lastname: "Bogomilov"}},
	}
	c.PostCrypto(cryptocurrency)

	//PUT request
	updatedCrypto := client.Cryptocurrency{
		Name:     "Bitcoin",
		CryptoID: "BTC",
		Price:    45000.3,
		Authors:  client.CryptoAuthors{crypto.Author{Firstname: "Satoshi", Lastname: "Nakamoto"}},
	}
	c.PutCrypto(updatedCrypto)

	// DELETE request
	c.DeleteCrypto("LCN")
}
