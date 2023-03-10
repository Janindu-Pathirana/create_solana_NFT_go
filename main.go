package main

import (
	"go/types"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/rpc"
)

type Metadata struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Attributes  []string `json:"attributes,omitempty"`
}

type Contract struct {
	client      *client.Client
	programID   types
	wallet      solana.Signer
	token       solana.PublicKey
	tokenAmount uint64
}

func main() {
	// create a RPC client
	c := client.NewClient(rpc.DevnetRPCEndpoint)

}
