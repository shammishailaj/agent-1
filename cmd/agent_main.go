package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kouhin/envflag"
	log "github.com/sirupsen/logrus"

	"github.com/opscoin/agent/eth"
)

var (
	TESTNET = "https://ropsten.infura.io/Bd64GbsMs0gTlYpHNPC0"
	MAINNET = "https://mainnet.infura.io/Bd64GbsMs0gTlYpHNPC0"
)

func setLogLevel(logLevel string) {
	if logLevel == "debug" {
		log.SetLevel(log.DebugLevel)
	} else if logLevel == "info" {
		log.SetLevel(log.InfoLevel)
	} else if logLevel == "warn" {
		log.SetLevel(log.WarnLevel)
	} else if logLevel == "error" {
		log.SetLevel(log.ErrorLevel)
	}

}
func getNetwork(network string) string {
	if network == "test" || network == "testnet" {
		return TESTNET
	}
	return MAINNET
}

func main() {
	network := flag.String("network_name", "main", "override to use the testnetwork, i.e test or main")
	logLevel := flag.String("log_level", "debug", "log level")
	contractAddress := flag.String("contract_address", "", "Address where to find your contract that stores ssh keys")

	//	flag.Parse()
	envflag.Parse()

	setLogLevel(*logLevel)
	networkAddress := getNetwork(*network)

	//logfmt format https://brandur.org/logfmt
	log.SetFormatter(&log.TextFormatter{})

	//TODO: add support for syslog
	log.SetOutput(os.Stdout)

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(networkAddress)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// Instantiate the contract and display its name
	token, err := eth.NewEthSSHKey(common.HexToAddress(*contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	for i := 0; i < 10; i++ {
		name, err := token.SshPublicKeys(nil, big.NewInt(int64(i)))
		if err != nil {
			fmt.Printf("Failed to retrieve token name: %v \n", err)
			continue
		}

		fmt.Printf("original(%d): %o\n", len(name), []byte(name))
	}
}
