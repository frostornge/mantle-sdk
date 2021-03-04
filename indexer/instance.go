package indexer

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/terra-project/mantle-sdk/graph"
	"github.com/terra-project/mantle-sdk/graph/generate"
	"github.com/terra-project/mantle-sdk/types"

	mttypes "github.com/terra-project/mantle-sdk/types"
)

type Instance interface {
	RunRound()
}

type baseInstance struct {
	indexers  []Indexer
	committer mttypes.GraphQLCommitter
	querier   mttypes.GraphQLQuerier
}

func NewBaseInstance(
	indexers []Indexer,
	querier mttypes.GraphQLQuerier,
	committer mttypes.GraphQLCommitter,
) Instance {
	return baseInstance{
		indexers:  indexers,
		committer: committer,
		querier:   querier,
	}
}

func (instance baseInstance) RunRound() {
	// create wait group for ALL indexers
	wg := sync.WaitGroup{}
	wg.Add(len(instance.indexers))

	for _, indexer := range instance.indexers {
		go func(i Indexer) {
			defer wg.Done()

			q := buildQuerier(
				instance.querier,
				[]mttypes.Model{reflect.ValueOf(i).Type()},
			)
			c := buildCommitter(instance.committer)

			if err := i.Index(q, c); err != nil {
				panic(err)
			}
		}(indexer)
	}

	wg.Wait()
}

func buildQuerier(querier mttypes.GraphQLQuerier, output []mttypes.Model) Querier {
	return func(query interface{}, variables mttypes.GraphQLParams) error {
		qs := generate.GenerateQuery(query, variables)
		result := querier(qs, variables, output)

		if result.HasErrors() {
			var errorsString = make([]string, len(result.Errors))
			for i, e := range result.Errors {
				errorsString[i] = e.Error()
			}

			return fmt.Errorf(
				"graphql query resulted in errors: %s",
				strings.Join(errorsString, " "),
			)
		}

		return graph.UnmarshalInternalQueryResult(result, query)
	}
}

func buildCommitter(committer types.GraphQLCommitter) Committer {
	return func(entity interface{}) error {
		return committer(entity)
	}
}
