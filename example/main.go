package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	sample_caller "interface_example/contracts"
)

// 1 0xfA20C84A7fDB1AB2335F2707C3009b138D770611
// 2 0xdFca404C3AD92a8A7989463F8eb4E64700F30fED
// c 0xbA33aBd12B722f1e381C0D4A41bEc501ee6d8324

func main() {
	callerAddress := os.Getenv("CONTRACT")
	implAddress := os.Getenv("IMPL_CONTRACT")
	keyFile := os.Getenv("KEY_FILE")

	bytes, err := ioutil.ReadFile(keyFile)
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.ToECDSA(bytes)
	if err != nil {
		panic(err)
	}

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		panic(err)
	}

	contract, err := sample_caller.NewSampleCaller(common.HexToAddress(callerAddress), client)
	if err != nil {
		panic(err)
	}

	if opts, err := createKeyedTransactor(client, privateKey); err != nil {
		panic(err)
	} else if tx, err := contract.SetSampleAddress(opts, common.HexToAddress(implAddress)); err != nil {
		panic(err)
	} else {
		waitTx(client, tx)
	}

	if opts, err := createKeyedTransactor(client, privateKey); err != nil {
		panic(err)
	} else if tx, err := contract.SetSampleValue(opts, big.NewInt(1)); err != nil {
		panic(err)
	} else {
		waitTx(client, tx)
	}

	sampleValue, err := contract.GetSampleValue(nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("sample value is: %d\n", sampleValue.Int64())
}

func waitTx(client *ethclient.Client, tx *types.Transaction) {
	ticker := time.NewTicker(time.Second)
	for {
		_, pending, err := client.TransactionByHash(context.Background(), tx.Hash())
		if err != nil {
			panic(err)
		}
		if !pending {
			break
		}

		<-ticker.C
	}
	ticker.Stop()
}

func createKeyedTransactor(client *ethclient.Client, privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	addr := crypto.PubkeyToAddress(privateKey.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(), addr)
	if err != nil {
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, err
	}

	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice

	return auth, nil
}
