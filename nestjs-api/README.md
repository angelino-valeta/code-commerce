# API NEST.JS PARA ORDEM DE PEDIDOS

## Rodar a aplicação

** Requisitos: **

- Docker
- Docker Compose
- Node.js


Execute o comando para subir o banco de dados e rabbitmq:

```
docker compose up
```

Em outro terminal execute os comandos para subir a aplicação:

```
npm install
npm run fixture
npm run start:dev
```

Existe um arquivo `api.http` na raiz do projeto que pode ser utilizado para testar a API via extensão `Rest Client` do VSCode.