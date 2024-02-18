# API Golang para catálogo de produtos

##  Rodar a aplicação

** Requisitos: **
- Docker
- Docker Compose
- Golang

Execute o comando para subir o banco de dados:
```
docker compose up
```

Em outro terminal execute a aplicação:

```
go run cmd/catalog/main.go
```

Existe um arquivo `api.http` na raiz do projecto que pode ser utilizado para testar a API via extensão `Rest Client` do VS CODE
