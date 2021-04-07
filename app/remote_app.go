package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	core "github.com/terra-project/core/types"
	"github.com/terra-project/mantle-sdk/committer"
	"github.com/terra-project/mantle-sdk/db"
	"github.com/terra-project/mantle-sdk/depsresolver"
	"github.com/terra-project/mantle-sdk/graph"
	"github.com/terra-project/mantle-sdk/graph/schemabuilders"
	"github.com/terra-project/mantle-sdk/indexer"
	"github.com/terra-project/mantle-sdk/querier"
	reg "github.com/terra-project/mantle-sdk/registry"
	"github.com/terra-project/mantle-sdk/types"
)

type RemoteMantle struct {
	db                   db.DB
	registry             *reg.Registry
	gqlInstance          *graph.RemoteGraphQLInstance
	depsResolverInstance depsresolver.DepsResolver
	committerInstance    committer.Committer
	indexerInstance      *indexer.IndexerBaseInstance
	baseMantleEndpoint   string
}

type RemoteSyncConfiguration struct {
	SyncUntil   uint64
	SyncFrom    uint64
	Reconnect   bool
	OnWSError   func(err error)
	OnInjectErr func(err error)
}

// NewRemoteMantle creates a mantle instance
// where there is no mantlemint
func NewRemoteMantle(
	db db.DB,
	baseMantleEndpoint string,
	indexers ...types.IndexerRegisterer,
) (mantleApp *RemoteMantle) {
	// create registry of indexers
	registry := reg.NewRegistry(indexers)

	// initialize deps resolver -- still required for inter-indexer sync
	depsResolverInstance := depsresolver.NewDepsResolver()

	querierInstance := querier.NewQuerier(db, registry.KVIndexMap)

	// create gql instance w/ remote deps resolver and only the injected indexers
	gqlInstance := graph.NewRemoteGraphQLInstance(
		depsResolverInstance,
		querierInstance,
		baseMantleEndpoint,
		schemabuilders.CreateRemoteModelSchemaBuilder(baseMantleEndpoint),
		schemabuilders.CreateModelSchemaBuilder(registry.KVIndexMap, registry.Models...),
	)

	// initializer committer
	committerInstance := committer.NewCommitter(db, registry.KVIndexMap)

	// initialize indexer
	indexerInstance := indexer.NewIndexerBaseInstance(
		registry.Indexers,
		registry.IndexerOutputs,
		gqlInstance.QueryInternal,
		gqlInstance.Commit,
	)

	// do init here, since we don't have mantlemint nor a core
	config := sdk.GetConfig()
	config.SetCoinType(core.CoinType)
	config.SetFullFundraiserPath(core.FullFundraiserPath)
	config.SetBech32PrefixForAccount(core.Bech32PrefixAccAddr, core.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(core.Bech32PrefixValAddr, core.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(core.Bech32PrefixConsAddr, core.Bech32PrefixConsPub)
	config.Seal()

	// create remote mantle
	mantleApp = &RemoteMantle{
		db:                   db,
		registry:             &registry,
		gqlInstance:          gqlInstance,
		depsResolverInstance: depsResolverInstance,
		committerInstance:    committerInstance,
		indexerInstance:      indexerInstance,
		baseMantleEndpoint:   baseMantleEndpoint,
	}

	return
}

var LastSyncedHeightQuery = "query{LastSyncedHeight}"

// Sync starts mantle as a remote mode.
// In remote mode the chain data is not synced (i.e. no Inject happens)
// and only indexers are run
type LastSyncedHeightResponse struct {
	Data struct {
		LastSyncedHeight uint64 `json:"LastSyncedHeight"`
	} `json:"data"`
}

func (rmantle *RemoteMantle) Sync(config RemoteSyncConfiguration) {
	var lastKnownHeight uint64 = 0
	// listen to LastSyncHeight change
	// trigger indexer
	// TODO: refactor this into gql subscription & reconnect logic
	for {
		// poll every 200ms
		time.Sleep(200 * time.Millisecond)

		// get currentHeight
		currentHeight := getLastSyncedHeight(rmantle.baseMantleEndpoint)
		if currentHeight <= lastKnownHeight {
			continue
		}

		// time
		tStart := time.Now()

		rmantle.indexerInstance.RunIndexerRound()
		indexerOutputs := rmantle.depsResolverInstance.GetState()

		// dispose deps resolver
		rmantle.depsResolverInstance.Dispose()

		// convert indexer outputs to slice
		var commitTargets = make([]interface{}, len(indexerOutputs))
		var i = 0
		for _, output := range indexerOutputs {
			commitTargets[i] = output
			i++
		}

		// set db critical zones
		rmantle.db.SetCriticalZone()

		// commit
		if commitErr := rmantle.committerInstance.Commit(currentHeight, commitTargets...); commitErr != nil {
			panic(commitErr)
		}

		// time end
		tEnd := time.Now()

		log.Printf(
			"[mantle] Indexing finished for block(%d), processed in %dms",
			currentHeight,
			tEnd.Sub(tStart).Milliseconds(),
		)

		lastKnownHeight = currentHeight

		// release db lock
		releaseErr := rmantle.db.ReleaseCriticalZone()
		if releaseErr != nil {
			panic(releaseErr)
		}
	}
}

func (rmantle *RemoteMantle) Server(port int) {
	go rmantle.gqlInstance.ServeHTTP(port)
}

func getLastSyncedHeight(baseMantleEndpoint string) uint64 {
	response, err := http.Get(fmt.Sprintf(
		"%s?query=%s",
		baseMantleEndpoint,
		LastSyncedHeightQuery,
	))

	if err != nil {
		panic(err)
	}

	bz, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	sync := LastSyncedHeightResponse{}
	if err := json.Unmarshal(bz, &sync); err != nil {
		panic(err)
	}

	return sync.Data.LastSyncedHeight
}
