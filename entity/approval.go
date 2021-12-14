package entity

type Approval struct {
	Msg         string
	Tipo        string
	Responsavel string
	Validacao   string
	Valor       float64
	Campo1      string
	Campo2      string
	Campo4      string
	Campo5      string
}

//Construtor
func NewApproval() *Approval {
	return &Approval{}
}
