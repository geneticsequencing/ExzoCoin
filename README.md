## Exzo Network Proof-of-Stake Blockchain Core
* Name - Exzo Token
* Symbol - XZO
* Supply - 800 million
* Blocktime - 3 seconds
* Consensus - PoS
* P2P Port - 1478
* JSON-RPC Port - 8545 
* EVM Compatible

## Official Links
* Website - https://exzo.network
* Main Explorer - https://explorer.exzo.technology
* Test Explorer - https://testnet.exzo.technology
* Twitter - https://twitter.com/exzo_network
* Telegram - https://t.me/exzocoin

### Build from Source (Ubuntu 20.04)
Requirements - `Go >=1.17`

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
wget https://go.dev/dl/go1.17.13.linux-amd64.tar.gz
sudo tar -xvf go1.17.13.linux-amd64.tar.gz
sudo mv go /usr/local && rm go1.17.13.linux-amd64.tar.gz
```
Check that it's installed
```
go version
```
You should see something like this:
```
go version go1.17.13 linux/amd64
```

#### Build Exzocoin
```
git clone https://github.com/ExzoNetwork/ExzoCoin.git
cd ExzoCoin/
go build -o exzocoin main.go
```

#### Running a Full Validating Node
After you have [downloaded](https://github.com/ExzoNetwork/ExzoCoin/releases/latest) the binaries or [built from source](https://github.com/ExzoNetwork/ExzoCoin#build-from-source), enter the `ExzoCoin` directory and run the following:
```
./exzocoin server --data-dir ~/.exzocoin --chain mainnet-genesis.json --seal --max-slots 40960 --grpc 0.0.0.0:9632 --libp2p 0.0.0.0:1478 --jsonrpc 0.0.0.0:8545 --enable-graphql --graphql-address 0.0.0.0:9898 --max-inbound-peers 128 --max-outbound-peers 16
```

#### Running a non-Validating node
```
server --data-dir ~/.exzocoin --chain mainnet-genesis.json --libp2p 0.0.0.0:1478 --nat<public_or_private_ip>
```


---
```
Copyright 2022 Polygon Technology
Copyright 2022 Exzocoin

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
