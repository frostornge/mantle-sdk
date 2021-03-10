package indexer

type TestQuery struct {
	Foo string
	Bar struct {
		Hello  uint64
		Mantle string
	}
}

// single
var _ Indexer = TestIndexer{}

type TestIndexer struct {
	Foo string
	Bar struct {
		Hello  uint64 `model:"index"`
		Mantle string `model:"index"`
	}
}

func (i TestIndexer) Name() string { return "" }
func (i TestIndexer) Index(q Querier, c Committer) error {
	req := TestQuery{}
	if err := q(&req, nil); err != nil {
		return err
	}

	entity := TestIndexer{
		Foo: req.Foo,
		Bar: struct {
			Hello  uint64 `model:"index"`
			Mantle string `model:"index"`
		}{
			Hello:  req.Bar.Hello,
			Mantle: req.Bar.Mantle,
		},
	}
	if err := c(entity); err != nil {
		return err
	}
	return nil
}

// slice
var _ Indexer = TestSliceIndexer{}

type TestSliceIndexer []TestIndexer

func (i TestSliceIndexer) Name() string { return "" }
func (i TestSliceIndexer) Index(q Querier, c Committer) error {
	var req struct {
		TestSliceIndexer []TestQuery
	}
	if err := q(&req, nil); err != nil {
		return err
	}

	var entity TestSliceIndexer
	for _, r := range req.TestSliceIndexer {
		entity = append(
			entity,
			TestIndexer{
				Foo: r.Foo,
				Bar: struct {
					Hello  uint64 `model:"index"`
					Mantle string `model:"index"`
				}{
					Hello:  r.Bar.Hello,
					Mantle: r.Bar.Mantle,
				},
			},
		)
	}
	if err := c(entity); err != nil {
		return err
	}
	return nil
}

// map
var _ Indexer = TestMapIndexer{}

type TestMapIndexer map[string]TestIndexer

func (i TestMapIndexer) Name() string { return "" }
func (i TestMapIndexer) Index(q Querier, c Committer) error {
	var req struct {
		TestMapIndexer map[string]TestQuery
	}
	if err := q(&req, nil); err != nil {
		return err
	}

	entity := make(TestMapIndexer)
	for k, r := range req.TestMapIndexer {
		entity[k] = TestIndexer{
			Foo: r.Foo,
			Bar: struct {
				Hello  uint64 `model:"index"`
				Mantle string `model:"index"`
			}{
				Hello:  r.Bar.Hello,
				Mantle: r.Bar.Mantle,
			},
		}
	}
	if err := c(entity); err != nil {
		return err
	}
	return nil
}
