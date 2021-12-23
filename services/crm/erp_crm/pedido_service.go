package erp_crm

import (
	crmRepository "github.com/augusto/imersao5-esquenta-go/adapter/repository/crm"
	"github.com/augusto/imersao5-esquenta-go/services/crm"
	"github.com/augusto/imersao5-esquenta-go/usecase/crm/pedido"
	"log"
)

func PedidoService(clientId string) error {
	log.Println("Inicio da transação dos procedimento de integração dos pedidos")

	//Chama a função que retorna as duas conexões
	dbCrmConn, _, ownerCrm, connError := crm.ServicesDatabases(clientId)

	if connError != nil {
		return connError
	}

	//Criando o repositório
	repoCrm := crmRepository.NewPedidoRepositoryDbErp(dbCrmConn)

	useCaseCrm := pedido.NewProcessPedido(repoCrm)

	//data, err := useCaseCrm.UseCaseSelect(ownerCrm)
	_, err := useCaseCrm.UseCaseSelect(ownerCrm)

	if err != nil {
		return err
	}

	//fmt.Println("Esse são o dados selecionados:")
	//fmt.Println(data)

	//Criar caso de uso para fazer a execução dos requests

	return nil
}
