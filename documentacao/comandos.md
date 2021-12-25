## Uteis

Para iniciar um módulo
`go mod init <nome do modulo> (github.com/augusto/imersao5-esquenta-go)`

Para rodar os testes
`go test ./...`

Gerando o mock do repositório
`mockgen -source=C:\Users\augusto.lourencatto\go\src\github.com\augusto\full_cycle_esquenta_go\entity\repository.go --destination=C:\Users\augusto.lourencatto\go\src\github.com\augusto\full_cycle_esquenta_go\entity\mock\mogo
.go
`

Rodando o comando no cmd (roda o server principal)
`go run .\main.go`

Acessando o banco sqlite3(sempre colocar o ; no final do comando sql)
`sqlite3 test.db`

##Testes
Para rodar os testes: `go test -v`  || para rodar em todas as pastas : `go test ./...`  
**obs** : Na classe de teste, começar o nome com 'Test'   
*exemplo* : TestFirstExample

##Docker 
**referência**: https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e
**obs**:Funcionando

Criando um docker file
`docker build . -t go-crm:0.1`

Executando o container
`docker run -p 8080:8080 go-crm:0.1`