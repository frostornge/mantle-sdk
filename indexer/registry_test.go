package indexer

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRegistry(t *testing.T) {
	Convey("init test", t, func() {
		Convey("#register", func() {
			testCases := map[string]struct {
				Name    string
				Indexer Indexer
			}{
				"should register single entity": {
					Name:    "TestIndexer",
					Indexer: TestIndexer{},
				},
				"should register slice entity": {
					Name:    "TestSliceIndexer",
					Indexer: TestSliceIndexer{},
				},
				"should register map entity": {
					Name:    "TestMapIndexer",
					Indexer: TestMapIndexer{},
				},
			}
			for title, testCase := range testCases {
				Convey(title, func() {
					registry := NewRegistry()
					registry.RegisterIndexer(testCase.Indexer)

					So(registry.Indexers(), ShouldHaveLength, 1)
					So(registry.Models(), ShouldHaveLength, 1)
					So(registry.KVIndexMap(), ShouldHaveLength, 2) // with BaseState

					kvi, ok := registry.KVIndexMap()[testCase.Name]
					So(ok, ShouldBeTrue)

					cases := map[string]string{
						"Hello":  "uint64",
						"Mantle": "string",
					}
					for indexName, indexType := range cases {
						entry, ok := kvi.Entry(indexName)
						So(ok, ShouldBeTrue)
						So(entry, ShouldNotBeNil)
						So(entry.Type().Name(), ShouldEqual, indexType)
					}
				})
			}
		})
	})
}
