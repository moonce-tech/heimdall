package staking

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hmCommon "github.com/maticnetwork/heimdall/common"
	"github.com/maticnetwork/heimdall/helper"
	hmTypes "github.com/maticnetwork/heimdall/types"
)

func NewHandler(k hmCommon.Keeper, contractCaller helper.IContractCaller) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgValidatorJoin:
			return HandleMsgValidatorJoin(ctx, msg, k, contractCaller)
		case MsgValidatorExit:
			return HandleMsgValidatorExit(ctx, msg, k, contractCaller)
		case MsgSignerUpdate:
			return HandleMsgSignerUpdate(ctx, msg, k)
		default:
			return sdk.ErrTxDecode("Invalid message in checkpoint module").Result()
		}
	}
}

func HandleMsgValidatorJoin(ctx sdk.Context, msg MsgValidatorJoin, k hmCommon.Keeper, contractCaller helper.IContractCaller) sdk.Result {
	hmCommon.StakingLogger.Debug("Handing new validator join", "Msg", msg)
	//fetch validator from mainchain
	validator, err := contractCaller.GetValidatorInfo(msg.ID)
	if err != nil || bytes.Equal(validator.Signer.Bytes(), helper.ZeroAddress.Bytes()) {
		hmCommon.StakingLogger.Error(
			"Unable to fetch validator from rootchain",
			"error", err,
			"msgValidator", msg.ID,
			"mainChainSigner", validator.Signer.String(),
		)
		return hmCommon.ErrNoValidator(k.Codespace).Result()
	}
	hmCommon.StakingLogger.Debug("Fetched validator from rootchain successfully", "validator", validator.String())

	// Generate PubKey from Pubkey in message and signer
	pubkey := msg.SignerPubKey
	signer := pubkey.Address()

	// check validator ID in message corresponds
	if !bytes.Equal(signer.Bytes(), validator.Signer.Bytes()) || msg.StartEpoch != validator.StartEpoch {
		hmCommon.StakingLogger.Error(
			"Signer Address or startEpoch or doesn't match",
			"msgValidator", msg.SignerPubKey.Address().String(),
			"mainchainValidator", validator.Signer.String(),
			"msgStartEpoch", msg.StartEpoch,
			"mainchainStartEpoch", validator.StartEpoch,
		)
		return hmCommon.ErrNoValidator(k.Codespace).Result()
	}

	// Check if validator has been validator before
	if _, ok := k.GetSignerFromValidatorID(ctx, msg.ID); ok {
		hmCommon.StakingLogger.Error("Validator has been validator before, cannot join with same address", "presentValidator", msg.ID)
		return hmCommon.ErrValidatorAlreadyJoined(k.Codespace).Result()
	}

	// create new validator
	newValidator := hmTypes.Validator{
		ID:    msg.ID,
		StartEpoch: msg.StartEpoch,
		EndEpoch:   msg.EndEpoch,
		Power:      msg.GetPower(),
		PubKey:     pubkey,
		Signer:     signer,
	}

	// add validator to store
	hmCommon.StakingLogger.Debug("Adding new validator to state", "validator", newValidator.String())
	err = k.AddValidator(ctx, newValidator)
	if err != nil {
		hmCommon.StakingLogger.Error("Unable to add validator to state", "error", err, "validator", newValidator.String())
		return hmCommon.ErrValidatorSave(k.Codespace).Result()
	}

	return sdk.Result{}
}

// Handle signer update message
func HandleMsgSignerUpdate(ctx sdk.Context, msg MsgSignerUpdate, k hmCommon.Keeper) sdk.Result {
	hmCommon.StakingLogger.Debug("Handling signer update", "Validator", msg.ID, "Signer", msg.NewSignerPubKey.Address())

	// pull val from store
	validator, ok := k.GetValidatorFromValID(ctx, msg.ID)
	if !ok {
		hmCommon.StakingLogger.Error("Fetching of validator from store failed", "validatorAddress", msg.ID)
		return hmCommon.ErrNoValidator(k.Codespace).Result()
	}

	newPubKey := msg.NewSignerPubKey
	newSigner := newPubKey.Address()

	// check if updating signer
	if !bytes.Equal(newSigner.Bytes(), validator.Signer.Bytes()) {
		// Update signer in prev Validator
		oldSigner := validator.Signer
		validator.Signer = newSigner
		validator.PubKey = newPubKey

		hmCommon.StakingLogger.Debug("Updating new signer", "signer", newSigner.String(), "oldSigner", oldSigner.String(), "validatorID", msg.ID)
	}

	// TODO add old signer sig check

	// power change
	if msg.NewAmount != "" && validator.Power != msg.GetNewPower() {
		// set new power
		validator.Power = msg.GetNewPower()
		hmCommon.StakingLogger.Debug("Updating power", "newPower", validator.Power, "validatorID", msg.ID)
	}

	// save validator
	err := k.AddValidator(ctx, validator)
	if err != nil {
		hmCommon.StakingLogger.Error("Unable to update signer", "error", err, "ValidatorID", validator.ID)
		return hmCommon.ErrSignerUpdateError(k.Codespace).Result()
	}

	return sdk.Result{}
}

func HandleMsgValidatorExit(ctx sdk.Context, msg MsgValidatorExit, k hmCommon.Keeper, contractCaller helper.IContractCaller) sdk.Result {
	hmCommon.StakingLogger.Info("Handling validator exit", "ValidatorID", msg.ID)

	validator, ok := k.GetValidatorFromValID(ctx, msg.ID)
	if !ok {
		hmCommon.StakingLogger.Error("Fetching of validator from store failed", "validatorID", msg.ID)
		return hmCommon.ErrNoValidator(k.Codespace).Result()
	}
	hmCommon.StakingLogger.Debug("validator in store", "validator", validator)
	// check if validator deactivation period is set
	if validator.EndEpoch != 0 {
		hmCommon.StakingLogger.Error("Validator already unbonded")
		return hmCommon.ErrValUnbonded(k.Codespace).Result()
	}

	// get validator from mainchain
	updatedVal, err := contractCaller.GetValidatorInfo(validator.ID)
	if err != nil {
		hmCommon.StakingLogger.Error("Cannot fetch validator info while unstaking", "Error", err, "validatorID", validator.ID)
		return hmCommon.ErrNoValidator(k.Codespace).Result()
	}

	// Add deactivation time for validator
	if err := k.AddDeactivationEpoch(ctx, validator, updatedVal); err != nil {
		hmCommon.StakingLogger.Error("Error while setting deactivation epoch to validator", "error", err, "validatorID", validator.ID)
		return hmCommon.ErrValidatorNotDeactivated(k.Codespace).Result()
	}

	return sdk.Result{}
}
