package app

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	l "log"

	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tm-db"
	TerraApp "github.com/terra-project/core/app"
	core "github.com/terra-project/core/types"
	wasmconfig "github.com/terra-project/core/x/wasm/config"
	types "github.com/terra-project/mantle/types"
	"github.com/terra-project/mantle/utils"
)

type App struct {
	terra *TerraApp.TerraApp
}

func NewApp(
	db dbm.DB,
	genesis *tmtypes.GenesisDoc,
) *App {
	config := sdk.GetConfig()
	config.SetCoinType(core.CoinType)
	config.SetFullFundraiserPath(core.FullFundraiserPath)
	config.SetBech32PrefixForAccount(core.Bech32PrefixAccAddr, core.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(core.Bech32PrefixValAddr, core.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(core.Bech32PrefixConsAddr, core.Bech32PrefixConsPub)
	config.Seal()
	
	app := TerraApp.NewTerraApp(
		log.NewTMLogger(ioutil.Discard),
		db,
		nil,
		true, // need this so KVStores are set
		0,
		make(map[int64]bool),
		&wasmconfig.Config{BaseConfig: wasmconfig.BaseConfig{
			ContractQueryGasLimit: viper.GetUint64(wasmconfig.FlagContractQueryGasLimit),
			CacheSize:             viper.GetUint64(wasmconfig.FlagCacheSize),
		}},
		fauxMerkleModeOpt, // error
		setPruningOptions(),
	)

	// only init chain on genesis
	if app.LastBlockHeight() == 0 {
		// init chain
		validators := make([]*tmtypes.Validator, len(genesis.Validators))
		for i, val := range genesis.Validators {
			validators[i] = tmtypes.NewValidator(val.PubKey, val.Power)
		}
		validatorSet := tmtypes.NewValidatorSet(validators)
		nextVals := tmtypes.TM2PB.ValidatorUpdates(validatorSet)
		csParams := tmtypes.TM2PB.ConsensusParams(genesis.ConsensusParams)
		ic := abci.RequestInitChain{
			Time:            genesis.GenesisTime,
			ChainId:         genesis.ChainID,
			AppStateBytes:   genesis.AppState,
			ConsensusParams: csParams,
			Validators:      nextVals,
		}

		initChainResponse := app.InitChain(ic)
		initChainResponseJSON, _ := json.Marshal(initChainResponse)
		commitResponse := app.Commit()
		commitResponseJSON, _ := json.Marshal(commitResponse)

		l.Printf("Init chain finished, LastBlockHeight=%d", app.LastBlockHeight())
		l.Printf("== InitChainResponse: %s", string(initChainResponseJSON))
		l.Printf("== CommitResponse: %s", string(commitResponseJSON))
	}


	return &App{
		terra: app,
	}
}

// Pass this in as an option to use a dbStoreAdapter instead of an IAVLStore for simulation speed.
func fauxMerkleModeOpt(bapp *baseapp.BaseApp) {
	bapp.SetFauxMerkleMode()
}

func setPruningOptions() func(*baseapp.BaseApp) {
	// prune nothing
	pruningOptions := sdk.PruningOptions{
		KeepRecent: 0,
		KeepEvery: 0,
		Interval: 10,
	}
	return baseapp.SetPruning(pruningOptions)
}

func (c *App) GetApp() *TerraApp.TerraApp {
	return c.terra
}

func (c *App) GetQueryRouter() sdk.QueryRouter {
	return c.terra.QueryRouter()
}

func (c *App) BeginBlocker(block *types.Block) abci.ResponseBeginBlock {
	var abciHeader = utils.ConvertToABCIHeader(&block.Header)
	var abciRequest = abci.RequestBeginBlock{
		Header: abciHeader,
	}

	abciResponse := c.terra.BeginBlock(abciRequest)

	return abciResponse
}

func (c *App) EndBlocker(block *types.Block) abci.ResponseEndBlock {
	abciRequest := abci.RequestEndBlock{
		Height: block.Header.Height,
	}
	abciResponse := c.terra.EndBlock(abciRequest)
	
	return abciResponse
}

func (c *App) DeliverTxs(txs []string) []abci.ResponseDeliverTx {
	responses := make([]abci.ResponseDeliverTx, len(txs))
	for _, tx := range txs {
		decoded, err := base64.StdEncoding.DecodeString(tx)
		if err != nil {
			panic(err)
		}
		abciRequest := abci.RequestDeliverTx{
			Tx: decoded,
		}
		response := c.terra.DeliverTx(abciRequest)
		responses = append(responses, response)
	}

	return responses
}

func (c *App) Commit(transactional bool) abci.ResponseCommit {
	response := c.terra.Commit()

	// no need to further save app state per block
	if !transactional {
		return response
	}
	// c.terra.ExportAppStateAndValidators()

	return response
}
