package pedido

import (
	"fmt"
	"github.com/augusto/imersao5-esquenta-go/entity/crm"
	"github.com/augusto/imersao5-esquenta-go/entity/crm/pedido"
	"github.com/augusto/imersao5-esquenta-go/utils"
	"github.com/augusto/imersao5-esquenta-go/utils/helpers"
)

type ProcessPedido struct {
	Repository crm.PedidoRepository
}

func NewProcessPedido(repository crm.PedidoRepository) *ProcessPedido {
	return &ProcessPedido{Repository: repository}
}

func (p *ProcessPedido) UseCaseSelect(owner string) ([]pedido.Quote, error) { //Mudar aqui
	//Faz select na quote
	quote, err := p.Repository.SelectQuote(owner)

	if err != nil {
		return nil, err
	}

	var payloads []pedido.PedidoPayload

	for _, x := range quote {

		payload := pedido.PedidoPayload{}

		items, errQI := p.Repository.SelectQuoteItem(owner, x.Id)

		if errQI != nil {
			utils.LogDatabase("PEDIDO", "CRITICAL", x.Id, false, errQI.Error())
			return nil, errQI
		}

		//Tratantando as linhas do pedido payload
		payload.Filial = helpers.String(x.Filial)
		payload.Cliente = helpers.String(x.Ti9Codigo)
		payload.CodOperacao = helpers.String(x.CodigoOperacao)
		payload.DataEntrada = helpers.String(x.DataEntrada)
		payload.CondPagamento = helpers.String(x.CondPagamento)
		payload.IndPres = helpers.String(x.Indpres)
		payload.OrigemRegistro = helpers.String(x.OrigemRegistro)
		payload.ObsSimples = helpers.String(x.ObsSimples)
		payload.OutrasInfoCom = helpers.String(x.OutrasInfoCom)
		payload.Moeda = helpers.Int(x.Moeda)
		payload.TipoFrete = helpers.String(x.TipoFrete)
		payload.Transp = helpers.String(x.Transp)
		payload.DescFinanc = helpers.String(x.DescFinanc)
		payload.CodUsuarioInclusao = "009249" //Usuário chapado
		payload.Finalidade = helpers.String(x.Finalidade)
		payload.NumeroPedPalm = x.Id

		//Tratando as linhas dos items do pedido -> Loop na resposta do SelectQuoteItem
		var itemsPayload []pedido.ItemsPayload

		for _, y := range items {
			itemPayload := pedido.ItemsPayload{}

			//Tratantando as linhas dos items do payload
			itemPayload.ItemPv = helpers.String(y.ItemPv)
		}

		//payload.Items = items

		fmt.Println(itemsPayload)

		//Fazendo o append no slice
		payloads = append(payloads, payload)

	}

	//Printando o json
	//u, err := json.Marshal(payloads)
	//fmt.Println(string(u))

	return quote, nil
}
