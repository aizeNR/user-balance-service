package v1

type topUpRequest struct {
	UserID uint64 `json:"user_id"`
	Amount uint64 `json:"amount"`
}

type writeOffRequest struct {
	UserID uint64 `json:"user_id"`
	Amount uint64 `json:"amount"`
}

type transferRequest struct {
	RceiverID uint64 `json:"receiver_id"`
	SenderID  uint64 `json:"sender_id"`
	Amount    uint64 `json:"amount"`
}

type getBalanceResponse struct {
	UserID uint64 `json:"user_id"`
	Balance float64 `json:"balance"`
}
