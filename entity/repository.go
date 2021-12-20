package entity

/**
 * Interface que implementa os métodos relacionados a entidade transaction
 */
type TransactionRepository interface {
	Insert(accountId string, amount float64, status string, errorMessage string) error
	Select(id string) ([]Transaction, error)
	DeleteTransaction(id string) error
}

//#TODO implementar a checagem de permissão, Resgatar fornecedores,se é a ultima aprovação,próximo email
/**
 * Implementa as ações do usuário nas aprovações
 */
type ApprovalRepository interface {
	Select(user string) ([]Approval, error)
	CheckPermission(usuario string) (bool, error)
	RecoverAllProviders(filial string, cotacao string) //#TODO (transformar em endpoint)
	Interact(filial string, cotacao string, fornecedor string, tipoDeAprovacao string, usuario string, statusDeAprovacao string, justificativa string, seqConcatenada string) error
}

//type QueryInfoRepository interface {
//	GetCards()
//}
