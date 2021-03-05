# Desafio Starwars Backend - Go

![imagem com animação de estrelas e com um texto escrito](https://github.com/hi-hi-ray/desafio-sw-go/blob/main/images/github-readme-cover.gif)

Tabela de conteúdos
=================
<!--ts-->
   * [Sobre](#sobre)
   * [Solicitação da empresa](#solicitação-da-empresa)
   * [Solução](#solução)
      * [Como usar](#como-usar)
        * [Ambiente Local](#ambiente-local)
      * [Testes](#testes)
      * [Roadmap](#roadmap)
      * [Tecnologias](#tecnologias)
<!--te-->

## Sobre

Esse repositório é um teste técnico feito com Go para B2W. 

## Solicitação da empresa

Nossos associados são aficionados por Star Wars e com isso, queremos criar um jogo com algumas informações da franquia.

Para possibilitar a equipe de front criar essa aplicação, queremos desenvolver uma API que contenha os dados dos planetas.

**Requisitos:**

- A API deve ser REST
- Para cada planeta, os seguintes dados devem ser obtidos do banco de dados da aplicação, sendo inserido manualmente:
    - Nome
    - Clima
    - Terreno
    - Quantidade
- Para cada planeta também devemos ter a quantidade de aparições em filmes, que podem ser obtidas pela API pública do Star Wars: [https://swapi.dev/about](https://swapi.dev/about)

**Funcionalidades desejadas:**

- Adicionar um planeta (com nome, clima e terreno)
- Listar planetas
- Buscar por nome
- Buscar por ID
- Remover planeta

**OBS: A linguagem para realização do desafio será correspondente a do anúncio da vaga.**

**Bancos que usamos:** MongoDB, Cassandra, DynamoDB, Datomic, ELK.

**E lembre-se! Um bom software é um software bem testado.**

-----

## Solução

### Como usar

#### Ambiente Local

1- Baixando o repositório.

Será necessário baixar o repositório dentro do caminho `GOPATH/src/github.com/hi-hi-ray/`, porem para certificar que o caminho existe, caso não rode `mkdir $GOPATH/src/github.com/hi-hi-ray/` para ser criado o caminho. De resto basta executar os comandos abaixo:

``` 
  cd $GOPATH/src/github.com/hi-hi-ray/
  git clone git@github.com:hi-hi-ray/desafio-starwars.git
```

2- Instalando os pacotes.

Os pacotes que estão no arquivo `vendor/vendor.json` são obrigatórios para que o projeto funcione corretamente. Para instalar eles com facilidade, basta rodar o comando a seguir:

``` 
  govendor install vendor/vendor.json
```

3- Preenchendo o arquivo de configuração.

Antes de rodar em qualquer ambiente é necessário copiar o arquivo `config.toml.example` e preencher os campos a seguir:

Tags:
* Database

  ✩ server (Tipo do campo: Texto e Obrigatório)

  ✩ database (Tipo do campo: Texto e Obrigatório)

  ✩ collection (Tipo do campo: Texto e Obrigatório)

  ✩ port (Tipo do campo: Inteiro e Obrigatório)

  ✩ username (Tipo do campo: Texto)

  ✩ password (Tipo do campo: Texto)
  
  ✩ timeout (Tipo do campo: Inteiro e Obrigatório)

* Servers
  
  Deve ser atualizada caso seja desejado mudar a porta da aplicação.

* Swapi
  
  Deve ser atualizada com a informação da API consumida para buscar a quantidade de planetas.

Para clonar esse arquivo com facilidade você pode usar o comando a seguir:

```shell script
cp ./api/config/config.toml.example ./api/config/config.toml
```

4- Executando a aplicação

Para executar a aplicação é necessário estar na pasta `api`. para isso pode-se executar os comandos a seguir:

``` 
  cd api
  go run main.go
```

Após a aplicação está no ar, você pode utilizar a documentação [Swagger](https://app.swaggerhub.com/apis/hi-hi-ray/DesafioStarWarsAPI-GO/1.0.0#/) para ver a utilização da api. 

### Testes

Para executar os testes existentes é aconselhável estar na pasta `api`, o comando que executa os testes seria:

```shell script
go test -v ./...
```

### Roadmap

[x] Adicionar um planeta (com nome, clima e terreno)

[x] Listar planetas

[x] Buscar por nome

[x] Buscar por ID

[x] Remover planeta

[x] Atualizar planeta usando um ID

[] Adicionar vários planetas (com nome, clima e terreno)

### Tecnologias:

As ferramentas e tecnologias para construir essa aplicação.

- [MinGW](https://sourceforge.net/projects/mingw-w64/)
- [Golang v:1.12](https://golang.org/dl/)
- [MongoDB v:4.0.9](https://www.mongodb.com/)
- [GoLand](https://www.jetbrains.com/pt-br/go/)
- [Govendor](https://github.com/kardianos/govendor)
