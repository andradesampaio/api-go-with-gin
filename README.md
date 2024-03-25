# Curso Go Alura 
Go e Gin: Criando API rest com simplicidade

# Faça esse curso de GoLang e:
- Crie uma API do zero com Go e Gin</br>
- Integre sua API Go com um banco de dados sendo executado no Docker</br>
- Aprenda a utilizar o GORM, o ORM mais famoso do Go</br>
- Saiba como criar a buscas de recursos com base nos campos de sua struct</br>
- Saiba como implementar sua própria API rest com Gin</br>

## Requisitos

- Docker
- Docker Compose
- Go

## Instruções de Uso

1. Clone este repositório para a sua máquina local.
2. Navegue até o diretório do projeto.
3. Antes de iniciar a aplicação, é necessário criar o banco de dados. Para isso, navegue até a pasta `infra` e execute `docker-compose up`.
4. O script de inicialização irá criar a tabela `personalidade` e inserir 10 registros.

## Acessando o Banco de Dados

Para acessar o banco de dados em execução, use o seguinte comando:

``bash 
docker exec -it <container_id_or_name> /bin/bash

# Projeto com Docker e PostgreSQL

Este projeto utiliza Docker e PostgreSQL para criar um banco de dados.

4. Para testar a aplicação, você pode usar o seguinte comando curl:

```bash
curl --request POST \
  --url http://localhost:8080/alunos/ \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/8.6.1' \
  --data '{
	"nome": "Maria João",
	"CPF": "34267543501",
	"RG": "45.456.879-0",
	"Email" : "teste@gmail.com"
}'

