package command

import "github.com/ExzoNetwork/ExzoCoin/server"

const (
	DefaultGenesisFileName = "genesis.json"
	DefaultChainName       = "exzocoin"
	DefaultChainID         = 1229
	DefaultPremineBalance  = "0xCECB8F27F4200F3A000000" // 250 million units of native network currency
	DefaultConsensus       = server.IBFTConsensus
	DefaultGenesisGasUsed  = 458752  // 0x70000
	DefaultGenesisGasLimit = 5242880 // 0x500000
)

const (
	JSONOutputFlag  = "json"
	GRPCAddressFlag = "grpc-address"
	JSONRPCFlag     = "jsonrpc"
)

// GRPCAddressFlagLEGACY Legacy flag that needs to be present to preserve backwards
// compatibility with running clients
const (
	GRPCAddressFlagLEGACY = "grpc"
)
