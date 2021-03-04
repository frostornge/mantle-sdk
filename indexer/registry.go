package indexer

import (
	"reflect"

	"github.com/terra-project/mantle-sdk/db/kvindex"
	mttypes "github.com/terra-project/mantle-sdk/types"
)

type Registry interface {
	Indexers() []Indexer
	Models() []mttypes.Model
	KVIndexMap() kvindex.KVIndexMap
	RegisterIndexer(indexers ...Indexer)
}

type registry struct {
	indexers   []Indexer
	models     []mttypes.Model
	kvIndexMap kvindex.KVIndexMap
}

func NewRegistry() Registry {
	// add BlockState to kvindexes
	typeBlockState := reflect.TypeOf(mttypes.BlockState{})
	blockStateKVIndex, blockStateKVIndexErr := kvindex.NewKVIndex(typeBlockState, "")
	if blockStateKVIndexErr != nil {
		panic(blockStateKVIndexErr)
	}

	return &registry{
		indexers:   []Indexer{},
		models:     []mttypes.Model{},
		kvIndexMap: kvindex.NewKVIndexMap(blockStateKVIndex),
	}
}

func (r registry) Indexers() []Indexer            { return r.indexers }
func (r registry) Models() []mttypes.Model        { return r.models }
func (r registry) KVIndexMap() kvindex.KVIndexMap { return r.kvIndexMap }

func (r *registry) RegisterIndexer(indexers ...Indexer) {
	for _, i := range indexers {
		r.indexers = append(r.indexers, i)

		model := reflect.ValueOf(i).Type()
		r.models = append(r.models, model)

		kvIndex, kvIndexErr := kvindex.NewKVIndex(model, i.Name())
		if kvIndexErr != nil {
			panic(kvIndexErr)
		}
		r.kvIndexMap[kvIndex.ModelName()] = kvIndex
	}
}
