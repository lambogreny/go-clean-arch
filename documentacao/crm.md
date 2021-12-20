## Documentando um exemplo de construção de uma etapa do CRM

### PRD

- **Criação da entidade da tabela PRD**
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
  