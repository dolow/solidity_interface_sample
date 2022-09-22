# What is this ?

This is a sample project to show solidity interface behavior.

Following specs can be proved with this project;

- contract can communicate via shared interface with an address of other contract that implements interface
- any contract address can be used as interface implementation

# Prerequisites

- truffle
- golang
- runnning local blockchain network
- private key of local payable account

# Usage

You need to deploy contracts to your network.

Please keep contract addresses, that is used in next step.

```
truffle migrate --network local
```

Run example to ensure dynamic contract address can be used as interface inplementation reference.

You can see every versions of interface implementation have different logic in the same API.

```
cd example
CONTRACT=0x... IMPL_CONTRACT=0x<any version of contract> KEY_FILE=/path/to/private_key go run main.go
CONTRACT=0x... IMPL_CONTRACT=0x<different version of contract> KEY_FILE=/path/to/private_key go run main.go
```

# Note

go package under example/contracts is build with [abigen](https://geth.ethereum.org/docs/dapp/native-bindings)

# License

MIT