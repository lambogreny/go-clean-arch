package entity

type TransactionRepository interface {
	Insert(accountId string, amount float64, status string, errorMessage string) error
	Select(id string) ([]Transaction, error)
	DeleteTransaction(id string) error
}

type ApprovalRepository interface {
	Select(user string) ([]Approval, error)
}
