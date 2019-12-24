package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/ethereum/go-ethereum/common"
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"

	"github.com/maticnetwork/heimdall/checkpoint"
	"github.com/maticnetwork/heimdall/helper"
	"github.com/maticnetwork/heimdall/staking"
	"github.com/maticnetwork/heimdall/types"
	hmRest "github.com/maticnetwork/heimdall/types/rest"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec) {
	r.HandleFunc(
		"/checkpoint/buffer",
		checkpointBufferHandlerFn(cdc, cliCtx),
	).Methods("GET")

	r.HandleFunc("/checkpoint/count",
		checkpointCountHandlerFn(cdc, cliCtx),
	).Methods("GET")

	r.HandleFunc(
		"/checkpoint/headers/{headerBlockIndex}",
		checkpointHeaderHandlerFn(cdc, cliCtx),
	).Methods("GET")

	r.HandleFunc("/checkpoint/latest-checkpoint",
		latestCheckpointHandlerFunc(cdc, cliCtx),
	).Methods("GET")

	r.HandleFunc("/checkpoint/{checkpointNumber}",
		checkpointByNumberHandlerFunc(cdc, cliCtx),
	).Methods("GET")

	r.HandleFunc("/checkpoint/{start}/{end}",
		checkpointHandlerFn(cdc, cliCtx),
	).Methods("GET")

	r.HandleFunc("/checkpoint/last-no-ack",
		noackHandlerFn(cdc, cliCtx)).Methods("GET")

	r.HandleFunc("/overview",
		overviewHandlerFunc(cdc, cliCtx)).Methods("GET")

	helper.InitHeimdallConfig("")
}

func checkpointBufferHandlerFn(
	cdc *codec.Codec,
	cliCtx context.CLIContext,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _, err := cliCtx.QueryStore(checkpoint.BufferCheckpointKey, "checkpoint")
		if err != nil {
			hmRest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		// the query will return empty if there is no data in buffer
		if len(res) == 0 {
			hmRest.WriteErrorResponse(w, http.StatusNoContent, errors.New("no content found for requested key").Error())
			return
		}

		var _checkpoint types.CheckpointBlockHeader
		err = cdc.UnmarshalBinaryBare(res, &_checkpoint)
		if err != nil {
			RestLogger.Error("Unable to unmarshall", "Error", err)
			hmRest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		RestLogger.Debug("Checkpoint fetched", "Checkpoint", _checkpoint.String())
		fmt.Printf("Checkpoint %v", _checkpoint.String())
		result, err := json.Marshal(&_checkpoint)
		if err != nil {
			RestLogger.Error("Error while marshalling response to Json", "error", err)
			hmRest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, result)
	}
}

func checkpointCountHandlerFn(
	cdc *codec.Codec,
	cliCtx context.CLIContext,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		RestLogger.Debug("Fetching number of checkpoints from state")
		res, _, err := cliCtx.QueryStore(checkpoint.ACKCountKey, "checkpoint")
		if err != nil {
			hmRest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		// The query will return empty if there is no data
		if len(res) == 0 {
			hmRest.WriteErrorResponse(w, http.StatusNoContent, errors.New("no content found for requested key").Error())
			return
		}

		ackCount, err := strconv.ParseInt(string(res), 10, 64)
		if err != nil {
			RestLogger.Error("Unable to parse int", "Response", res, "Error", err)
			hmRest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		result, err := json.Marshal(map[string]interface{}{"result": ackCount})
		if err != nil {
			RestLogger.Error("Error while marshalling resposne to Json", "error", err)
			hmRest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, result)
	}
}

func checkpointHeaderHandlerFn(
	cdc *codec.Codec,
	cliCtx context.CLIContext,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		// get header number
		headerNumber, ok := rest.ParseUint64OrReturnBadRequest(w, vars["headerBlockIndex"])
		if !ok {
			return
		}

		res, _, err := cliCtx.QueryStore(checkpoint.GetHeaderKey(uint64(headerNumber)), "checkpoint")
		if err != nil {
			hmRest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		// the query will return empty if there is no data
		if len(res) == 0 {
			hmRest.WriteErrorResponse(w, http.StatusNotFound, errors.New("no content found for requested key").Error())
			return
		}
		var _checkpoint types.CheckpointBlockHeader
		err = cdc.UnmarshalBinaryBare(res, &_checkpoint)
		if err != nil {
			hmRest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		result, err := json.Marshal(&_checkpoint)
		if err != nil {
			RestLogger.Error("Error while marshalling resposne to Json", "error", err)
			hmRest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, result)
	}
}

// HeaderBlockResult represents header block result
type HeaderBlockResult struct {
	Proposer   types.HeimdallAddress `json:"proposer"`
	RootHash   common.Hash           `json:"rootHash"`
	StartBlock uint64                `json:"startBlock"`
	EndBlock   uint64                `json:"endBlock"`
}

func checkpointHandlerFn(
	cdc *codec.Codec,
	cliCtx context.CLIContext,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		start, err := strconv.Atoi(vars["start"])
		if err != nil {
			hmRest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		end, err := strconv.Atoi(vars["end"])
		if err != nil {
			hmRest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		roothash, err := checkpoint.GetHeaders(uint64(start), uint64(end))
		if err != nil {
			RestLogger.Error("Unable to get header", "Start", start, "End", end, "Error", err)
			hmRest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var validatorSet types.ValidatorSet

		_validatorSet, _, err := cliCtx.QueryStore(staking.CurrentValidatorSetKey, "staking")
		if err == nil {
			err := cdc.UnmarshalBinaryBare(_validatorSet, &validatorSet)
			if err != nil {
				hmRest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				RestLogger.Error("Unable to get validator set to form proposer", "Error", err)
				return
			}
		}

		checkpoint := HeaderBlockResult{
			Proposer:   validatorSet.Proposer.Signer,
			StartBlock: uint64(start),
			EndBlock:   uint64(end),
			RootHash:   ethcmn.BytesToHash(roothash),
		}

		result, err := json.Marshal(checkpoint)
		if err != nil {
			RestLogger.Error("Error while marshalling resposne to Json", "error", err)
			hmRest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, result)
	}
}

func noackHandlerFn(
	cdc *codec.Codec,
	cliCtx context.CLIContext,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _, err := cliCtx.QueryStore(checkpoint.LastNoACKKey, "checkpoint")
		if err != nil {
			hmRest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if len(res) == 0 {
			hmRest.WriteErrorResponse(w, http.StatusNoContent, errors.New("no content found for requested key").Error())
			return
		}

		lastAckTime, err := strconv.ParseInt(string(res), 10, 64)
		if err != nil {
			RestLogger.Error("Unable to parse int", "Response", res, "Error", err)
			hmRest.WriteErrorResponse(w, http.StatusNoContent, errors.New("no content found for requested key").Error())
			return
		}

		result, err := json.Marshal(map[string]interface{}{"result": lastAckTime})
		if err != nil {
			RestLogger.Error("Error while marshalling resposne to Json", "error", err)
			hmRest.WriteErrorResponse(w, http.StatusNoContent, errors.New("no content found for requested key").Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, result)
	}
}

type stateDump struct {
	ACKCount         int64                         `json:AckCount`
	CheckpointBuffer types.CheckpointBlockHeader   `json:HeaderBlock`
	ValidatorCount   int                           `json:ValidatorCount`
	ValidatorSet     types.ValidatorSet            `json:ValidatorSet`
	LastNoACK        time.Time                     `json:LastNoACKTime`
	Headers          []types.CheckpointBlockHeader `json:"headers"`
}

// get all state-dump of heimdall
func overviewHandlerFunc(
	cdc *codec.Codec,
	cliCtx context.CLIContext,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// ACk count
		var ackCountInt int64
		ackcount, _, err := cliCtx.QueryStore(checkpoint.ACKCountKey, "checkpoint")
		if err == nil {
			ackCountInt, err = strconv.ParseInt(string(ackcount), 10, 64)
			if err != nil {
				RestLogger.Error("Unable to parse int for getting ack count", "Response", ackcount, "Error", err)
			}
		} else {
			RestLogger.Error("Unable to fetch ack count from store", "Error", err)
		}

		// checkpoint buffer
		var _checkpoint types.CheckpointBlockHeader
		_checkpointBufferBytes, _, err := cliCtx.QueryStore(checkpoint.BufferCheckpointKey, "checkpoint")
		if err == nil {
			if len(_checkpointBufferBytes) != 0 {
				err = cdc.UnmarshalBinaryBare(_checkpointBufferBytes, &_checkpoint)
				if err != nil {
					RestLogger.Error("Unable to unmarshall checkpoint present in buffer", "Error", err, "CheckpointBuffer", _checkpointBufferBytes)
				}
			} else {
				RestLogger.Error("No checkpoint present in buffer")
			}
		} else {
			RestLogger.Error("Unable to fetch checkpoint from buffer", "Error", err)
		}

		// validator count
		var validatorCount int
		var validatorSet types.ValidatorSet

		_validatorSet, _, err := cliCtx.QueryStore(staking.CurrentValidatorSetKey, "staking")
		if err == nil {
			cdc.UnmarshalBinaryBare(_validatorSet, &validatorSet)
		}
		validatorCount = len(validatorSet.Validators)

		// last no ack
		var lastNoACKTime int64
		lastNoACK, _, err := cliCtx.QueryStore(checkpoint.LastNoACKKey, "checkpoint")
		if err == nil {
			lastNoACKTime, err = strconv.ParseInt(string(lastNoACK), 10, 64)
		}

		var headers []types.CheckpointBlockHeader
		storedHeaders, _, err := cliCtx.QuerySubspace(checkpoint.HeaderBlockKey, "checkpoint")
		if err != nil {
			RestLogger.Error("Unable to query subspace for headers", "Error", err)
		}

		for _, kvPair := range storedHeaders {
			var checkpointHeader types.CheckpointBlockHeader
			if cdc.UnmarshalBinaryBare(kvPair.Value, &checkpointHeader); err != nil {
				RestLogger.Error("Unable to unmarshall header", "Error", err, "Value", kvPair.Value)
			}
			headers = append(headers, checkpointHeader)
		}

		state := stateDump{
			ACKCount:         ackCountInt,
			CheckpointBuffer: _checkpoint,
			ValidatorCount:   validatorCount,
			ValidatorSet:     validatorSet,
			LastNoACK:        time.Unix(lastNoACKTime, 0),
			Headers:          headers,
		}
		result, err := json.Marshal(map[string]interface{}{"result": state})
		if err != nil {
			RestLogger.Error("Error while marshalling resposne to Json", "error", err)
			hmRest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, result)
	}
}

// get last checkpoint from store
func latestCheckpointHandlerFunc(
	cdc *codec.Codec,
	cliCtx context.CLIContext,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ackCount, _, err := cliCtx.QueryStore(checkpoint.ACKCountKey, "checkpoint")
		if err != nil {
			hmRest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		ackCount, err := strconv.ParseUint(string(_ackCount), 10, 64)
		if err != nil {
			hmRest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		RestLogger.Debug("ACK Count fetched", "ackCount", ackCount)
		lastCheckpointKey := helper.GetConfig().ChildBlockInterval * ackCount
		RestLogger.Debug("Last checkpoint key generated",
			"lastCheckpointKey", lastCheckpointKey,
			"min", helper.GetConfig().ChildBlockInterval,
		)

		res, _, err := cliCtx.QueryStore(checkpoint.GetHeaderKey(lastCheckpointKey), "checkpoint")
		if err != nil {
			hmRest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// the query will return empty if there is no data
		if len(res) == 0 {
			hmRest.WriteErrorResponse(w, http.StatusNotFound, errors.New("No content found for requested key").Error())
			return
		}

		var _checkpoint types.CheckpointBlockHeader
		err = cdc.UnmarshalBinaryBare(res, &_checkpoint)
		if err != nil {
			hmRest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		RestLogger.Debug("Fetched last checkpoint", "Checkpoint", _checkpoint)
		result, err := json.Marshal(&_checkpoint)
		if err != nil {
			RestLogger.Error("Error while marshalling resposne to Json", "error", err)
			hmRest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, result)
	}
}

// get  checkpoint by checkppint number from store
func checkpointByNumberHandlerFunc(
	cdc *codec.Codec,
	cliCtx context.CLIContext,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)

		// get checkpoint number
		checkpointNumber, ok := rest.ParseUint64OrReturnBadRequest(w, vars["checkpointNumber"])
		if !ok {
			return
		}

		RestLogger.Debug("Get Checkpoint for ", "checkpointNumber", checkpointNumber)
		checkpointKey := helper.GetConfig().ChildBlockInterval * checkpointNumber
		RestLogger.Debug("checkpoint key generated",
			"checkpointKey", checkpointKey,
			"min", helper.GetConfig().ChildBlockInterval,
		)

		res, _, err := cliCtx.QueryStore(checkpoint.GetHeaderKey(checkpointKey), "checkpoint")
		if err != nil {
			hmRest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// the query will return empty if there is no data
		if len(res) == 0 {
			hmRest.WriteErrorResponse(w, http.StatusNotFound, errors.New("No content found for requested key").Error())
			return
		}

		var _checkpoint types.CheckpointBlockHeader
		err = cdc.UnmarshalBinaryBare(res, &_checkpoint)
		if err != nil {
			hmRest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		RestLogger.Debug("Fetched checkpoint", "Checkpoint", _checkpoint)
		result, err := json.Marshal(&_checkpoint)
		if err != nil {
			RestLogger.Error("Error while marshalling resposne to Json", "error", err)
			hmRest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, result)
	}
}
