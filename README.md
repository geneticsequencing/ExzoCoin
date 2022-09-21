## Exzocoin
* Name - EXZO
* Symbol - XZO
* Supply - 800 million
* Blocktime - 3 seconds
* Consensus - PoS
* P2P Port - 1478
* JSON-RPC Port - 8545 
* EVM Compatible

## Official Links
* Website - https://exzocoin.com
* Main Explorer - https://explorer.exzocoin.com
* Test Explorer - https://testnet.exzocoin.com
* Twitter - https://twitter.com/exzocoin
* Telegram - https://t.me/exzocoin

### Build from Source
Requirements - `Go >=1.17`
```
git clone https://github.com/ExzoNetwork/ExzoCoin.git
cd ExzoCoin/
go build -o exzocoin main.go
```

#### Running a Full Node
After you have [downloaded](https://github.com/ExzoNetwork/ExzoCoin/releases/latest) the binaries or [built from source](https://github.com/ExzoNetwork/ExzoCoin#build-from-source), enter the `ExzoCoin` directory and run the following:
```
./exzocoin server --data-dir /data/exzocoin --chain mainnet-genesis.json --seal --max-slots 40960 --grpc 0.0.0.0:9632 --libp2p 0.0.0.0:1478 --jsonrpc 0.0.0.0:8545 --enable-graphql --graphql-address 0.0.0.0:9898 --max-inbound-peers 128 --max-outbound-peers 16
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