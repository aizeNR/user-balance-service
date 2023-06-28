package model

type Paging struct {
	Limit uint64
	Page  uint64
}

func (p Paging) GetOffset() uint64 {
	return (p.Page - 1) * p.Limit
}
