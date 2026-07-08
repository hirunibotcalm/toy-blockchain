package main

import (
    "toy-blockchain/blockchain"
	"toy-blockchain/cli"
)

func main() {
    chain := blockchain.NewBlockchain()

	c := cli.CLI{
		Chain: chain,
	}

	c.Run()
}