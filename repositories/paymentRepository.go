package repositories

type PaymentRepository interface {
	TransferBalance(from, to string, amount uint) (bool, error)
}
