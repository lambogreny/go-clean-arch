package prd

import (
	"github.com/pkg/errors"
	"time"
)

type Prd struct {
	Tipo                string
	Codigo_produto      string
	Descricao_produto   string
	Cod_tipi            string
	Grupo_estoque       string
	Grupo_estoque_n2    string
	Grupo_estoque_n3    string
	Grupo_estoque_n4    string
	Unidade             string
	Marca               string
	Data_cad            time.Time
	Usuario_inclusao    string
	Ultimo_preco_liq    string
	Data_hora_alteracao time.Time
	Usuario_alteracao   string
	Partnumber          string
	Ativo               string
	Peso_liquido        string
	Peso_bruto          string
}

func NewPrd() *Prd {
	return &Prd{}
}

func (t *Prd) CheckTipo() error {
	if t.Tipo == "" {
		return errors.New("O Tipo não pode ser nulo")
	}
	return nil
}
