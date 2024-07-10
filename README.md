## Go Ethereum For Libplanet Rollup

Forked from [ethereum-optimism/op-geth](https://github.com/ethereum-optimism/op-geth)


## Operating a private network

Building geth

```shell
make geth
```

Creating a directory for storing blockchain datas

```shell
mkdir blockchainData
```

Creating an account for mining and signing

```shell
./geth --datadir ./blockchainData account new
```

Writing a genesis block to blockchain data

```shell
./geth --datadir ./blockchainData init ./genesis.json
```

`genesis.json` file looks like this. You have to change `extradata` and `alloc` property depending on the account you made. And you can change `balance` of `alloc` arbitrary. In this config, we use `clique` for setting up as a PoA network. See https://geth.ethereum.org/docs/fundamentals/private-network#clique-example for more informations.

```shell
{
  "config": {
    "chainId": 12345,
    "homesteadBlock": 0,
    "eip150Block": 0,
    "eip155Block": 0,
    "eip158Block": 0,
    "byzantiumBlock": 0,
    "constantinopleBlock": 0,
    "petersburgBlock": 0,
    "istanbulBlock": 0,
    "berlinBlock": 0,
    "clique": {
      "period": 5,
      "epoch": 30000
    }
  },
  "difficulty": "1",
  "gasLimit": "8000000",
  "extradata": "0x0000000000000000000000000000000000000000000000000000000000000000CE70F2e49927D431234BFc8D439412eef3a6276b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
  "alloc": {
    "0xCE70F2e49927D431234BFc8D439412eef3a6276b": { "balance": "1111111111111111" },
    "0xaA2337b6FC4EDcc99FBDc9dee5973c94849dCEce": { "balance": "2222222222222222" }
  }
}
```

Running geth on the setting. You have to change `--unlock` and `--miner.etherbase` arguments depending on the account you made. See https://geth.ethereum.org/docs/fundamentals/command-line-options for more informations.

```shell
./geth --datadir ./blockchainData --networkid 12345 --http --http.port 8000 --http.api eth,net,web3,miner,personal,admin --allow-insecure-unlock --unlock 0xCE70F2e49927D431234BFc8D439412eef3a6276b --mine --miner.etherbase 0xCE70F2e49927D431234BFc8D439412eef3a6276b
```

You can check the geth with JsonRPC. Check https://geth.ethereum.org/docs/getting-started#interact-with-geth

```shell
./geth attach http://127.0.0.1:8000
```

## License

The go-ethereum library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html),
also included in our repository in the `COPYING.LESSER` file.

The go-ethereum binaries (i.e. all code inside of the `cmd` directory) are licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also
included in our repository in the `COPYING` file.
