package pedido

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

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

// -------------------------------------------------- CRM -> ERP -----------------------------

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
			utils.LogDatabaseDetails("PEDIDO", "CALL API", "PARSE STRUCT TO JSON", payloadErr.Error(), "")
			return payloadErr
		}

		// -----------------------------------------------------------Montando e executando o request ------------------------------------------------------------//
		payload := strings.NewReader(string(u))

		req, reqError := http.NewRequest("POST", extra.Base_url, payload)

		if reqError != nil {
			log.Println("ERRO NA MONTAGEM DA REQUISIÇÃO")
			utils.LogDatabaseDetails("PEDIDO", "CALL API", "MONTAGEM REQUEST", reqError.Error(), "")
			return reqError
		}

		req.Header.Add("Content-Type", "application/json")

		res, resError := http.DefaultClient.Do(req)

		if resError != nil {
			log.Println("ERRO NA CHAMADA DA API")
			utils.LogDatabaseDetails("PEDIDO", "CALL API", "ERRO REQUEST API", reqError.Error(), "")
			return resError
		}
		defer res.Body.Close()

		body, bodyError := ioutil.ReadAll(res.Body)

		if bodyError != nil {
			log.Println("ERRO NA LEITURA DE RESPOSTA")
			utils.LogDatabaseDetails("PEDIDO", "CALL API", "ERRO LEITURA DA RESPOSTA", reqError.Error(), "")
			return bodyError
		}

		if res.StatusCode == 500 {
			utils.LogDatabaseDetails("PEDIDO", "CALL API", "STATUS 500 API PEDIDo", reqError.Error(), "")
			fmt.Errorf("Internal server error na api de pedidos no ti9")
		}

		if res.StatusCode != 201 {
			utils.LogDatabaseDetails("PEDIDO", "CALL API", "STATUS DIFERENTE DE 201", reqError.Error(), "")
			fmt.Errorf(string(body))
		}

		fmt.Println(res.StatusCode)
		fmt.Println(string(body))

		// -----------------------------------------------------------Deletando o registro da sincroniza ------------------------------------------------------------//

		deleteSincroniza := p.Repository.DeleteSincroniza(extra.Owner, x.NumeroPedPalm)

		if deleteSincroniza != nil {
			return deleteSincroniza
		}

		//break
	}

	return nil

}

//-------------------------------------------------------------- ERP -> CRM ---------------------------
func (p *ProcessPedido) UseCaseSelectCpv() ([]pedido.Cpv, error) {

	cpv, err := p.Repository.SelectCpv()

	if err != nil {
		return nil, err
	}

	return cpv, nil

}

func (p *ProcessPedido) UseCaseCheckUpdateCrm(id string, owner string) (bool, error) {

	resp, err := p.Repository.CheckUpdateCrm(id, owner)

	if err != nil {
		return false, err
	}
	return resp, nil
}

func (p *ProcessPedido) UseCaseUpdate(cpv pedido.Cpv, owner string) error {

	err := p.Repository.UpdateCrm(cpv, owner)
	// log.Println("Retorno do UpdateCrm : ", err)

	if err != nil {
		log.Println("(CAMADA DE USE CASE) Erro ao atualizar o pedido no crm")
		return err
	}

	log.Println("(CAMADA DE USE CASE) Pedido atualizado com sucesso no crm")

	return nil

}

func (p *ProcessPedido) UseCaseInsert(cpv pedido.Cpv, owner string) error {

	err := p.Repository.InsertCrm(cpv, owner)

	if err != nil {
		return err
	}
	return nil
}

func (p *ProcessPedido) UseCaseDelete(id string, tipo string) error {

	err := p.Repository.DeleteErp(id, tipo)

	if err != nil {
		return err
	}
	return nil

}
