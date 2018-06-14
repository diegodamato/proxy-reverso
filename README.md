Desafio Globo.com Version 1.0 14/06/2018

INFORMAÇÕES GERAIS
-------------------

- Proxy reverso


INSTALAÇÃO
----------

- O projeto utiliza o GoVendor como ferramenta de gerenciamento de dependência. Para instalar o GoVendor, segue comando abaixo:
go get -u github.com/kardianos/govendor

- Execute o seguinte comando após clonar o repositório do projeto:
govendor sync


CONFIGURANDO O ARQUIVO DE PROPRIEDADE YML
-----------------------------------------

O arquivo de configuração possui as seguintes propriedades configuráveis:

- location: contém o path de origem, o host de destino e o path de destino.
- originpath: path de origem
- destination: host de destino
- destinationpath: path de destino


COMANDOS PARA RODAR O SOFTWARE
----------------------------------------------

- Compilação do projeto
go build -o proxy-reverso cmd/proxy-reverso/main.go

- Executando em desenvolvimento
go run cmd/proxy-reverso/main.go
