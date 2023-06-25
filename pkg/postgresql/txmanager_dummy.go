package postgresql

import "context"

type DummyManager struct{}

func NewDumyTxManager() *DummyManager {
	return &DummyManager{}
}

func (s *DummyManager) RunTx(ctx context.Context, do func(ctx context.Context) error) error {
	return do(ctx)
}
