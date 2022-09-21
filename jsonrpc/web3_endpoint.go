package jsonrpc

import (
	"fmt"

	"github.com/ExzoNetwork/ExzoCoin/helper/hex"
	"github.com/ExzoNetwork/ExzoCoin/helper/keccak"
	"github.com/ExzoNetwork/ExzoCoin/versioning"
)

// Web3 is the web3 jsonrpc endpoint
type Web3 struct{}

// ClientVersion returns the version of the web3 client (web3_clientVersion)
func (w *Web3) ClientVersion() (interface{}, error) {
	return fmt.Sprintf("exzocoin [%s]", versioning.Version), nil
}

// Sha3 returns Keccak-256 (not the standardized SHA3-256) of the given data
func (w *Web3) Sha3(val string) (interface{}, error) {
	v, err := hex.DecodeHex(val)
	if err != nil {
		return nil, NewInvalidRequestError("Invalid hex string")
	}

	dst := keccak.Keccak256(nil, v)

	return hex.EncodeToHex(dst), nil
}
