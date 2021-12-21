## Documentando um exemplo de construção de uma etapa do CRM

### Exemplo PRD

- **Criação da entidade da tabela PRD**  
  `entity.crm.prd.prd_ti9`
- **Criação do prd_repository**  
  Implementa os métodos relacionados a prd
- **Criação do crmRepository**  
  Implementa efetivamente as operações no banco de dados
- **Criação do DTO para entidade da PRD**  
  Objetivo de utilização do DTO para manipular a entidade
- **Criação do caso de uso**  
  `process_prd_crm.go`  
  Case de uso para determinada ação
- **Criação do adaptador**  
  `api.controllers.crm.prd`
  Controller que chama os repositórios e casos de uso
  

## Camada de serviço
`services.crm.erp_crm` -> Sai do erp e vai para o crm