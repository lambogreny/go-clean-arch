## Fluxo de criação de um novo endpoint

 - Criação de uma entidade
    - Escrever testes antes de criar;
    - Criar funções de validação para os componentes da entidade
 - Escrever os métodos da interface
 - asdsa

### Exemplo da criação da feature de Consulta de cotações

- Criação do payload da entrada da api  
  **package** : ``query_info.QueryInfo``
- Criação do adpter do endpoint
  - Registrar no gin router
    **package** : `api.NewRouter.`  
      `queryInfo := new(controllers.QueryInfoController)`
      `v1.POST("/queryInfo/cards", queryInfo.GetCards)`
  - Criar o controller
    **package** : `controllers.GetCards.`  
  - 
