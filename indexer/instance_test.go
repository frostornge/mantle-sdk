package indexer

import (
	"reflect"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/terra-project/mantle-sdk/types"
)

func recoverAndErr(fn func()) string {
	msg := ""
	func() {
		defer func() {
			if err := recover(); err != nil {
				msg = err.(error).Error()
			}
		}()
		fn()
	}()
	return msg
}

func TestInstance(t *testing.T) {
	Convey("init test", t, func() {
		indexers := []Indexer{
			TestIndexer{},
			TestSliceIndexer{},
			TestMapIndexer{},
		}
		defaultQuerier := types.GraphQLQuerier(
			func(q string, v types.GraphQLParams, d []types.Model) *graphql.Result {
				return &graphql.Result{
					Data: map[string]interface{}{},
				}
			},
		)
		defaultCommitter := types.GraphQLCommitter(
			func(e interface{}) error {
				return nil
			},
		)

		Convey("#RunRound", func() {
			testCases := map[string]struct {
				Name    string
				Indexer Indexer
				Result  *graphql.Result
			}{
				"should process single index": {
					Name:    "TestIndexer",
					Indexer: TestIndexer{},
					Result: &graphql.Result{
						Data: map[string]interface{}{
							"Foo": "hello",
							"Bar": map[string]interface{}{
								"Hello":  64,
								"Mantle": "World",
							},
						},
					},
				},
				"should process slice index": {
					Name:    "TestSliceIndexer",
					Indexer: TestSliceIndexer{},
					Result: &graphql.Result{
						Data: map[string]interface{}{
							"TestSliceIndexer": []map[string]interface{}{{
								"Foo": "hello",
								"Bar": map[string]interface{}{
									"Hello":  64,
									"Mantle": "World",
								},
							}},
						},
					},
				},
				"should process map index": {
					Name:    "TestMapIndexer",
					Indexer: TestMapIndexer{},
					Result: &graphql.Result{
						Data: map[string]interface{}{
							"TestMapIndexer": map[string]map[string]interface{}{
								"test": {
									"Foo": "hello",
									"Bar": map[string]interface{}{
										"Hello":  64,
										"Mantle": "World",
									},
								},
							},
						},
					},
				},
			}
			for title, testCase := range testCases {
				Convey(title, func() {
					instance := NewBaseInstance(
						[]Indexer{testCase.Indexer},
						func(q string, v types.GraphQLParams, d []types.Model) *graphql.Result {
							return testCase.Result
						},
						func(e interface{}) error {
							name := reflect.ValueOf(e).Type().Name()
							if name != testCase.Name {
								return errors.Errorf("%s, %s", name, testCase.Name)
							}
							return nil
						},
					)
					instance.RunRound()
				})
			}
			Convey("should panic if querier returns error", func() {
				instance := NewBaseInstance(
					indexers,
					func(q string, v types.GraphQLParams, d []types.Model) *graphql.Result {
						return &graphql.Result{
							Errors: []gqlerrors.FormattedError{{
								Message: "TestError",
							}},
						}
					},
					defaultCommitter,
				)

				msg := recoverAndErr(instance.RunRound)
				So(msg, ShouldNotBeEmpty)
				So(msg, ShouldContainSubstring, "TestError")
			})
			Convey("should panic if committer returns error", func() {
				instance := NewBaseInstance(
					indexers,
					defaultQuerier,
					func(e interface{}) error { return errors.New("TestError") },
				)

				msg := recoverAndErr(instance.RunRound)
				So(msg, ShouldNotBeEmpty)
				So(msg, ShouldContainSubstring, "TestError")
			})
		})
	})
}
