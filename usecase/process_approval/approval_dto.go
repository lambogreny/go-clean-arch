package process_approval

type ApprovalDtoInput struct {
	user string `form:"user" binding:"required"`
}

/**
 * Struct de entrada para aprovação, reprovação e devolução
 */
type ApprovalDtoInteractionInput struct {
	Filial            string `json:"filial" binding:"required"`
	Cotacao           string `json:"cotacao" binding:"required"`
	Fornecedor        string `json:"fornecedor" binding:"required"`
	TipoDeAprovacao   string `json:"tipoDeAprovacao" binding:"required"`
	Usuario           string `json:"usuario" binding:"required"`
	StatusDeAprovacao string `json:"statusDeAprovacao" binding:"required"`
	Justificativa     string `json:"justificativa"`
	SeqConcatenada    string `json:"SeqConcatenada" binding:"required"`
}
