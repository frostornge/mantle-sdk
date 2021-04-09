package indexer

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/terra-project/mantle-sdk/graph"
	"github.com/terra-project/mantle-sdk/graph/generate"
	"github.com/terra-project/mantle-sdk/types"
	mttypes "github.com/terra-project/mantle-sdk/types"
	"golang.org/x/sync/errgroup"
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
	g := new(errgroup.Group)
	for _, indexer := range instance.indexers {
		g.Go(instance.buildIndexerRunner(indexer))
	}

	if err := g.Wait(); err != nil {
		panic(err)
	}
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

func (instance baseInstance) buildIndexerRunner(i Indexer) func() error {
	return func() error {
		q := buildQuerier(
			instance.querier,
			[]mttypes.Model{reflect.ValueOf(i).Type()},
		)
		c := buildCommitter(instance.committer)

		if err := i.Index(q, c); err != nil {
			return err
		}
		return nil
	}
}
