# Exzo Network - Proof-of-Stake Blockchain Core
**Chain Specs**
* Native Token Name - Exzo
* Symbol - XZO
* Supply - 250 Million
* Blocktime - 3 seconds
* Consensus - IBFT Proof-of-Stake (PoS)
* P2P Port - 1478
* JSON-RPC Port - 8545 
* ChainID Mainnet - 1229
* ChainID Testnet - 2370
* Ethereum Virtual Machine (EVM) Compatible

**Official Links**
* Website - https://exzo.network
* Main Explorer - https://exzoscan.io
* Test Explorer - https://testnet.exzoscan.io
* Twitter - https://twitter.com/exzo_network
* Telegram - https://t.me/Exzo_Network
* Telegram - https://t.me/exzoc

## About the Exzo Network 
* Exzo Network uses a modular architecture and provides Ethereum compatibility.
* Users able to interact with industry standard wallets through JSON-RPC.
* Develop with Solidity/Vyper, full EVM support.
* Build using widely-adopted Ethereum tooling, libraries and development tools.
* Streamlined UX when doing cross-chain operations.
* Go beyond Ethereum’s Smart Contracts with Runtime plugins.

The Exzo Network Token (XZO) features the following utilities, and the value of XZO token will accrue with the increased usage of the network and revenue from stability fees and liquidation penalties.

* As Network Utility Token: to pay for network fees and stability fees.
* As Governance Token: to vote for/against risk parameters and network change proposals.
* As Economic Capital: in case of liquidation without sufficient collaterals.
* As fee reducing token on decentralized applications and wallets upon the Exzo Network.
* For staking to help secure the Exzo Network in a decentralized fashion.

Cross-Chain Communication
* Completely trustless and decentralized built-in Ethereum Bridge solution.
* Transfer assets to and from any EVM compatible network, most notably Exzo Network and Ethereum mainnets.
* Transfer ERC-20 tokens, NFTs or wrapped native currencies.
* Customize the bridge functionality using Bridge plugins.

## Testing Framework References
* https://www.researchgate.net/publication/340890880_Meta-model_for_Stress_Testing_on_Blockchain_Nodes
* 

**Note: This section is still work in progress, we will be adding more information as we progress.**

### Build from Source (Ubuntu 20.04)
Requirements - `Go >=1.18.x`

#### Setup Go Path
```
sudo nano ~/.profile
```
Paste this into the bottom of the file
```
export GOPATH=$HOME/work
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```
```
source ~/.profile
```

#### Install Go
```
wget https://go.dev/dl/go1.18.7.linux-amd64.tar.gz
sudo tar -xvf go1.18.7.linux-amd64.tar.gz
sudo mv go /usr/local && rm go1.18.7.linux-amd64.tar.gz
```
Check that it's installed
```
go version
```
You should see something like this:
```
go version go1.18.7 linux/amd64
```

#### Build Exzo Network
```
git clone https://github.com/ExzoNetwork/ExzoCoin.git
cd ExzoCoin/
go build -o exzocoin main.go
```

#### Setup a Validator Node (optional)
You can setup your node as a validator to help secure the network. You can follow [this guide](ValidatorGuide.md) to setting one up. 

#### Running a Non-validating Node
After you have [downloaded](https://github.com/ExzoNetwork/Exzo-Network-Core/releases/latest) the binaries or [built from source](https://github.com/ExzoNetwork/ExzoCoin#build-from-source), enter the `ExzoCoin` directory and run the following:
```
mkdir ~/.exzocoin
```
```
./exzocoin server --data-dir ~/.exzocoin --chain mainnet-genesis.json --libp2p 0.0.0.0:1478 --nat <public_or_private_ip>
```


---
```
Copyright 2022 Exzo Network

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
