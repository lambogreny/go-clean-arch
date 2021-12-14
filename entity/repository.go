package entity

type TransactionRepository interface {
	Insert(accountId string, amount float64, status string, errorMessage string) error
	Select() ([]Transaction, error)
}

type ApprovalRepository interface {
	Select() ([]Approval, error)
}
