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

type baseRegistry struct {
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

	return &baseRegistry{
		indexers:   []Indexer{},
		models:     []mttypes.Model{},
		kvIndexMap: kvindex.NewKVIndexMap(blockStateKVIndex),
	}
}

func (r baseRegistry) Indexers() []Indexer            { return r.indexers }
func (r baseRegistry) Models() []mttypes.Model        { return r.models }
func (r baseRegistry) KVIndexMap() kvindex.KVIndexMap { return r.kvIndexMap }

func (r *baseRegistry) RegisterIndexer(indexers ...Indexer) {
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
