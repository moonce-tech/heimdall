package staking

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/maticnetwork/heimdall/checkpoint"
	"github.com/tendermint/tendermint/crypto"
)

type Keeper struct {
	storeKey         sdk.StoreKey
	cdc              *codec.Codec
	checkpointKeeper checkpoint.Keeper

	// codespace
	codespace sdk.CodespaceType
}

var (
	ValidatorsKey = []byte{0x01} // prefix for each key to a validator
)

type Validator struct {
	ValidatorAddress common.Address
	StartEpoch       int64
	EndEpoch         int64
	Pubkey           crypto.PubKey
	Power            int64 // aka Amount
	Signer           common.Address
}

// NewKeeper creates new keeper for staking
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, codespace sdk.CodespaceType) Keeper {
	keeper := Keeper{
		storeKey:  key,
		cdc:       cdc,
		codespace: codespace,
	}
	return keeper
}

func (k Keeper) AddValidator(ctx sdk.Context) {

}
