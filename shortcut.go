package prompt

import "context"

func NoopExecutor(in string) {}

// Input get the input data from the user and return it.
func Input(opts ...Option) string {
	pt := New(NoopExecutor)
	pt.renderer.prefixTextColor = DefaultColor

	for _, opt := range opts {
		if err := opt(pt); err != nil {
			panic(err)
		}
	}
	return pt.Input()
}

// InputCtx get the input data from the user and return it.
func InputCtx(ctx context.Context, opts ...Option) (string, error) {
	pt := New(NoopExecutor)
	pt.renderer.prefixTextColor = DefaultColor

	for _, opt := range opts {
		if err := opt(pt); err != nil {
			panic(err)
		}
	}
	return pt.InputCtx(ctx)
}
