package errx

const (
	NotEnoughtMoney Code = "user.balance.money.not_enought"
)

type ErrNotEnoughtMoney struct {}

func (e *ErrNotEnoughtMoney) Error() string {
	return "not enought money"
}

func (e *ErrNotEnoughtMoney) Code() Code {
	return NotEnoughtMoney
} 

func (e ErrNotEnoughtMoney) Description() string {
	return "User does not have enough money!"
}

