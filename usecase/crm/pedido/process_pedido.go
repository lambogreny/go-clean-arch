package pedido

import (
	"encoding/json"
	"fmt"
	"github.com/augusto/imersao5-esquenta-go/entity/crm"
	"github.com/augusto/imersao5-esquenta-go/entity/crm/pedido"
	"github.com/augusto/imersao5-esquenta-go/utils/helpers"
	"log"
)

type ProcessPedido struct {
	Repository crm.PedidoRepository
}

func NewProcessPedido(repository crm.PedidoRepository) *ProcessPedido {
	return &ProcessPedido{Repository: repository}
}

func (p *ProcessPedido) UseCaseSelect(owner string, extra helpers.ExtraInfo) ([]pedido.PedidoPayload, error) {

	//Faz select no pedido
	quote, err := p.Repository.SelectQuote(owner)

	if err != nil {
		return nil, err
	}

	//Slice de struct de resposta
	var payloads []pedido.PedidoPayload

	//Percorrendo todos os dados recuperados no `SelectQuote`
	for _, x := range quote {

		payload := pedido.PedidoPayload{}

		items, errQI := p.Repository.SelectQuoteItem(owner, x.Id)

		if errQI != nil {

			return nil, errQI
		}

		//Tratantando as linhas do `SelectQuote`
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

		//Tratando as linhas dos items do pedido
		var itemsPayload []pedido.ItemsPayload

		//Percorrendo todos os dados recuperados no `SelectQuoteItem`
		for _, y := range items {
			itemPayload := pedido.ItemsPayload{}

			//Tratantando as linhas do `SelectQuoteItem`
			itemPayload.ItemPv = helpers.String(y.ItemPv)
			itemPayload.ProductId = helpers.String(y.ProductId)
			itemPayload.Almoxarifado = helpers.String(y.Almoxarifado)
			itemPayload.DataEntregaItem = helpers.String(y.DataEntregaItem)
			itemPayload.Finalidade = helpers.String(y.Finalidade)
			itemPayload.QtdePedida = helpers.String(y.QtdePedida)
			itemPayload.UnitPrice = helpers.String(y.UnitPrice)
			itemPayload.Discount = helpers.String(y.Discount)
			itemPayload.DescValor = helpers.Int(y.DescValor)
			itemPayload.CodigoCentro = helpers.String(y.CodigoCentro)

			itemsPayload = append(itemsPayload, itemPayload)
		}

		//Colocando todos os itens no pedido
		payload.Items = itemsPayload

		//Fazendo o append no slice de resposta
		payloads = append(payloads, payload)

	}

	return payloads, nil
}

func (p *ProcessPedido) UseCaseCallApi(payloads []pedido.PedidoPayload, extra helpers.ExtraInfo) error {
	fmt.Println("O url base da api do pedido é : ", extra.Base_url)
	fmt.Println("Total de payloads  : ", len(payloads))

	//Percorrendos todos os payloads
	for _, x := range payloads {
		//Transformando o payload struct para json
		u, payloadErr := json.Marshal(x)

		if payloadErr != nil {
			log.Println("PAYLOAD ERROR")
			return payloadErr
		}

		fmt.Printf(string(u))

		// -----------------------------------------------------------Montando e executando o request ------------------------------------------------------------//
		//payload := strings.NewReader(string(u))
		//
		//req, reqError := http.NewRequest("POST", extra.Base_url, payload)
		//
		//if reqError != nil {
		//	log.Println("ERRO NA MONTAGEM DA REQUISIÇÃO")
		//	return reqError
		//}
		//
		//req.Header.Add("Content-Type", "application/json")
		//
		//res, resError := http.DefaultClient.Do(req)
		//
		//if resError != nil {
		//	log.Println("ERRO NA CHAMADA DA API")
		//	return resError
		//}
		//defer res.Body.Close()
		//
		//body, bodyError := ioutil.ReadAll(res.Body)
		//
		//if bodyError != nil {
		//	log.Println("ERRO NA LEITURA DE RESPOSTA")
		//	return bodyError
		//}
		//
		//if res.StatusCode == 500 {
		//	fmt.Errorf("Internal server error na api de pedidos no ti9")
		//}
		//
		//if res.StatusCode != 201 {
		//	fmt.Errorf(string(body))
		//}
		//
		//fmt.Println(res.StatusCode)
		//fmt.Println(string(body))

		// -----------------------------------------------------------Deletando o registro da sincroniza ------------------------------------------------------------//

		//#TODO aqui chamar o delete na sincroniza
		deleteSincroniza := p.Repository.DeleteSincroniza(extra.Owner, x.NumeroPedPalm)

		if deleteSincroniza != nil {
			return deleteSincroniza
		}

		//break
	}

	return nil

}
