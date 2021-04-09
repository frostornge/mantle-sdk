package indexer

import mttypes "github.com/terra-project/mantle-sdk/types"

type Querier func(request interface{}, variables mttypes.GraphQLParams) error

type Committer func(entity interface{}) error

type Indexer interface {
	Name() string
	Index(q Querier, c Committer) error
}
