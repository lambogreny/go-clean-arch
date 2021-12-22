package erp_crm

import (
	"fmt"
	crmRepository "github.com/augusto/imersao5-esquenta-go/adapter/repository/crm"
	"github.com/augusto/imersao5-esquenta-go/services/crm"
	"github.com/augusto/imersao5-esquenta-go/usecase/crm/pedido"
	"log"
)

func PedidoService(clientId string) error {
	log.Println("Inicio da transação dos procedimento de integração dos pedidos")

	//Chama a função que retorna as duas conexões
	dbCrmConn, _, _, connError := crm.ServicesDatabases(clientId)

	if connError != nil {
		return connError
	}

	//Criando o repositório
	repoCrm := crmRepository.NewPedidoRepositoryDbErp(dbCrmConn)

	useCaseCrm := pedido.NewProcessPedido(repoCrm)

	err := useCaseCrm.UseCaseSelect()
	fmt.Println(err)

	return nil
}
