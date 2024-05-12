# Sistema de Reservas

## Descrição

Este é um sistema de reservas desenvolvido em Go, utilizando o framework Gin para roteamento e middleware, o ORM GORM
para interação com o banco de dados. O sistema permite aos usuários criar, modificar, e cancelar reservas, além de
autenticar usuários e gerenciar sessões com JWT.

## Funcionalidades

- Autenticação de usuário e gestão de sessões com JWT.
- Criação, modificação e cancelamento de reservas.
- Verificação de disponibilidade de recursos.

## Tecnologias Utilizadas

- **Go**: Linguagem de programação.
- **Gin**: Framework web usado para roteamento e middleware.
- **GORM**: ORM para interação com banco de dados.
- **JWT**: Tokens de autenticação para gerenciar sessões de usuários.
- **Postgres**: Banco de dados utilizado para desenvolvimento.

## Instalação

### Pré-requisitos

- Docker

### Configuração do Ambiente

1. Clone o repositório.

```bash
git clone https://github.com/marcos-nsantos/reservation-api.git
```

2. Acesse o diretório do projeto.

```bash
cd reservation-api
```

3. Crie o arquivo `.env` a partir do arquivo de exemplo.

```bash
cp .env.example .env
```

4. Execute o comando `docker-compose up` para iniciar a aplicação com o banco de dados.

```bash
docker-compose up
```

Por padrão, a aplicação estará disponível em `http://localhost:8080`.