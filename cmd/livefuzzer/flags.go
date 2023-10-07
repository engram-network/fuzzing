package main

import "github.com/urfave/cli/v2"

var (
	seedFlag = &cli.Int64Flag{
		Name:  "seed",
		Usage: "Seed for the RNG, (Default = RandomSeed)",
		Value: 0,
	}

	skFlag = &cli.StringFlag{
		Name:  "sk",
		Usage: "Secret key",
		Value: "1a2f6d731224de48f78b73ee630ef040b8785fdeb3831539fda03863885b8008",
	}

	corpusFlag = &cli.StringFlag{
		Name:  "corpus",
		Usage: "Use additional Corpus",
	}

	noALFlag = &cli.BoolFlag{
		Name:  "no-al",
		Usage: "Disable accesslist creation",
		Value: false,
	}

	countFlag = &cli.IntFlag{
		Name:  "count",
		Usage: "Count of addresses to create",
		Value: 100,
	}

	rpcFlag = &cli.StringFlag{
		Name:  "rpc",
		Usage: "RPC provider",
		Value: "https://engram.tech/testnet",
	}

	txCountFlag = &cli.IntFlag{
		Name:  "txcount",
		Usage: "Number of transactions send per account per block, 0 = best estimate",
		Value: 0,
	}
)
