package process_transaction

type TransactionDtoInput struct {
	//ID        string  `json:"id" binding:"required"`
	AccountID string  `json:"account_id" binding:"required"`
	Amount    float64 `json:"amount"  binding:"required"`
}

type TransactionDtoOutput struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

type DeleteTransactionOutput struct {
	Message string `json:"message"`
}

//type TransactionDtoSelect struct {
//
//}
