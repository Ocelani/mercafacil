# Processo Seletivo Backend Mercafácil

Solução para integração entre sistemas.

## Especificações

O problema consiste em receber 1 ou mais contatos de celulares através de uma API Rest e adicioná-los ao
banco de dados do cliente Macapá ou do cliente VareJão.

### Fluxo de Ações

- [x] A API receberá um JSON via POST contendo o nome e celular;
- [x] O cliente deverá estar autenticado para inserir o contato na base
- O contato deverá ser inserido no banco de dados do cliente seguindo as regras de cada cliente

### Especificações da API:

- [x] A autenticação será através de um token JWT no Authorization Header
- [x] Cada cliente tem 1 uma chave única
- A lista de contatos que será inserido em cada cliente está no arquivo contato.json

### Especificações do Cliente Macapá:

- [x] Banco de dados Mysql
- [ ] Formato do Nome é somente maiúsculas
- [ ] O formato de telefone segue o padrão +55 (41) 93030-6905
- Em anexo está o sql de criação da tabela

### Especificações do Cliente Varejão:

- [x] Banco de dados Postgresql
- [x] Formato do Nome é livre
- [ ] O formato de telefone segue o padrão 554130306905
- Em anexo está o sql de criação da tabela

## Sugestões

A criação de um ambiente de testes usando Docker para simular o banco de dados do cliente é altamente recomendada.

A solução poderá ser desenvolvida em Golang ou Node.js.

Fique livre para desenhar a solução da maneira que achar mais conveniente e supor qualquer cenário que
não foi abordado nas especificações acima.

Se, por qualquer motivo, você não consiga completar este teste,
recomendamos que nos encaminhe o que foi desenvolvido de qualquer maneira.

A falta de cumprimento de alguns dos requisitos aqui descritos não implica necessariamente na desconsideração do candidato.
