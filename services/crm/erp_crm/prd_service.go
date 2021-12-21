package erp_crm

import (
	"fmt"
	"github.com/augusto/imersao5-esquenta-go/services/crm"
	"log"
)

func PrdService(clientId string) error {
	fmt.Println("Inicio da transação dos procedimento de integração da PRD")

	//Chama a função que retorna as duas conexões
	//databaseErp, databaseCrm, err := crm.ServicesDatabases("sad")
	dbCrmConn, dbErpConn, err := crm.ServicesDatabases(clientId)
	fmt.Println(dbCrmConn, dbErpConn)

	if err != nil {
		log.Println("Falha ao se conectar com os bancos de dados!")
		return err
	}

	//fmt.Println(databaseErp)
	//fmt.Println(databaseCrm)

	//Criando o repositório
	//repo := crmRepository.NewPrdRepositoryDbErp(dbErpConn)
	//fmt.Println(repo)

	//usecase := prd.NewProcessPrd(repo)
	//
	//data, err := usecase.Select()
	//
	//if err != nil {
	//	return err
	//}
	//
	//for _, x := range data {
	//
	//	//Checando o se deve realizar o update
	//	checkUpdate, err := usecase.CheckUpdateCrm(x.Codigo_produto)
	//
	//	if err != nil {
	//		return err
	//	}
	//
	//	switch checkUpdate {
	//	case true:
	//		fmt.Println("Irei realizar a operação de update!")
	//
	//	case false:
	//		fmt.Println("Irei realizar a operação de insert!")
	//
	//	}
	//
	//}

	return nil
}
