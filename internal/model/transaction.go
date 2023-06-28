package model

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type Transaction struct {
	ID            uuid.UUID
	UserID        uint64
	Amount        int64
	OperationDate time.Time
	Comment       string
}
