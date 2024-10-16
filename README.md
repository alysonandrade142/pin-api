# Projeto de API com Go, Testify e Scripts Python

Este projeto é uma API desenvolvida em Go, utilizando o GORM para a interação com o banco de dados, e com testes unitários escritos com a biblioteca Testify. Também contém scripts Python para manipulação de dados e integração com o banco de dados.

## Requisitos

Certifique-se de ter os seguintes itens instalados no seu ambiente:

- [Docker](https://www.docker.com/get-started) - Para rodar o banco de dados e outros serviços.
- [Python 3.x](https://www.python.org/downloads/) - Para rodar scripts de manipulação de dados.
- [Golang](https://golang.org/doc/install) - Para rodar e testar a API.

## Passo a passo para execução do projeto

### 1. Inicialize o ambiente com Docker

O primeiro passo é subir os serviços necessários (como o banco de dados PostgreSQL) utilizando Docker. Para isso, execute o seguinte comando no terminal:

```bash

docker-compose up -d

```

Com o Docker rodando normalmente, será necessário criar e importar os dados na base utilizando um script python. 

```bash

pip install -r scripts/requirements.txt

py scripts/main.py

```

No final, podemos rodar nossos testes e verificar se eles estão de acordo.

```bash

go test ./tests/

```

Para acessar o conteúdo dos serviços, basta utilizar o postman que foi disponibilizado na pasta docs, lembre-se que é importante que os serviços e a base estejam criados.