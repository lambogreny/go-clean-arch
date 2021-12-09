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
`go run .\cmd\main.go`

Acessando o banco sqlite3(sempre colocar o ; no final do comando sql)
`sqlite3 test.db`
