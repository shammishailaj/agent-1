package main

import (
	"flag"
	"io/ioutil"
	"math/big"
	"os"
	"strings"

	"github.com/opscoin/agent/authorized_keys"
	"github.com/opscoin/agent/eth"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kouhin/envflag"
	log "github.com/sirupsen/logrus"
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
	keyFile := flag.String("key_file", "/root/.ssh/authorized_keys", "Disk location of authorized keys file")

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
	var keys []string
	for i := 0; i < 100; i++ {
		name, err := token.SshPublicKeys(nil, big.NewInt(int64(i)))
		if err != nil {
			if strings.Index(err.Error(), "unmarshalling empty output") == -1 {
				log.Errorf("Failed to retrieve token name: %s \n", err.Error())
			}
			break
		}
		keys = append(keys, name)
		log.Debug("original(%d) %s\n", len(name), name)
	}

	//Read text from keys file
	b, err := ioutil.ReadFile(*keyFile) // just pass the file name
	if err != nil {
		log.Fatal(err)
	}

	output, err := authorized_keys.AddKeysToFile(string(b), keys)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(*keyFile, []byte(output), 0600)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Completed. Exiting")
}
