package errx

const (
	NotEnoughtMoney Code = "user.balance.money.not_enought"
	BalanceNotFound Code = "user.balance.not_found"
)

type ErrNotEnoughtMoney struct{}

func (e *ErrNotEnoughtMoney) Error() string {
	return "not enought money"
}

func (e *ErrNotEnoughtMoney) Code() Code {
	return NotEnoughtMoney
}

func (e ErrNotEnoughtMoney) Description() string {
	return "User does not have enough money!"
}

type ErrBalanceNotFound struct{}

func (e *ErrBalanceNotFound) Error() string {
	return "balance not found"
}

func (e *ErrBalanceNotFound) Code() Code {
	return BalanceNotFound
}

func (e ErrBalanceNotFound) Description() string {
	return "User does not have balance!"
}
