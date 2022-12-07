package jsonrpc

import (
	"fmt"

	"github.com/ExzoNetwork/ExzoCoin/helper/keccak"
	"github.com/ExzoNetwork/ExzoCoin/versioning"
)

// Web3 is the web3 jsonrpc endpoint
type Web3 struct {
	chainID   uint64
	chainName string
}

var clientVersionTemplate = "%s [chain-id: %d] [version: %s]"

// ClientVersion returns the version of the web3 client (web3_clientVersion)
func (w *Web3) ClientVersion() (interface{}, error) {
	return fmt.Sprintf(
		clientVersionTemplate,
		w.chainName,
		w.chainID,
		versioning.Version,
	), nil
}

// Sha3 returns Keccak-256 (not the standardized SHA3-256) of the given data
func (w *Web3) Sha3(v argBytes) (interface{}, error) {
	dst := keccak.Keccak256(nil, v)

	return argBytes(dst), nil
}
