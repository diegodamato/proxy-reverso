Desafio Concrete Version 1.0 14/06/2018

INFORMAÇÕES GERAIS
-------------------

- Proxy reverso


INFORMAÇÕES DE COMANDOS PARA RODAR O SOFTWARE
----------------------------------------------

- Compilação do projeto
go build -o proxy-reverso cmd/proxy-reverso/main.go

- Executando em desenvolvimento
go run cmd/proxy-reverso/main.go


EXEMPLO DE ARQUIVO DE CONFIGURAÇÃO YML
---------------------------------------

address:
  - location:
    originpath: "/origin1"
    destination: "localhost1"
    destinationpath: "/v1/destino1"
  - location:
    originpath: "/origin2"
    destination: "localhost2"
    destinationpath: "/v1/destino2"
