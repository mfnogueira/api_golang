# API em Go com Gin Framework

Esta é uma API simples construída com o **Gin Framework** em Go, que permite gerenciar itens (criação e listagem) e realizar validação de dados, como email, ao criar um novo item.

## Estrutura do Projeto

A estrutura do projeto segue o padrão MVC (Model-View-Controller) para organizar a lógica da aplicação e separar responsabilidades. Abaixo está a estrutura básica do projeto:

```bash
api_golang/
├── controllers/              # Contém a lógica de controle para as rotas da API
│   ├── item_controller.go    # Controlador para itens
│   └── controller.go         # Outros controladores (se necessário)
├── routes/                   # Contém a configuração das rotas da API
│   ├── routes.go             # Definição das rotas principais
│   └── setup.go              # Remover este arquivo (duplicado de routes.go)
├── utils/                    # Funções auxiliares, como validação e log
│   ├── logging.go            # Funções de log
│   └── validation.go         # Funções de validação (como validação de email)
├── main.go                   # Arquivo principal para iniciar o servidor
├── go.mod                    # Dependências do projeto
└── go.sum                    # Hashes das dependências

```
```bash

## Como Rodar o Projeto Localmente

1. **Clone o repositório**:
   ```bash
   git clone https://github.com/seu-usuario/api_golang.git
   cd api_golang
```
2. **Instale as dependências**: O projeto usa o Gin Framework. Para instalar as dependências, execute o seguinte comando:
```bash
go mod tidy
```
3. **Execute a aplicação**: Para rodar a API, use o comando abaixo:
```bash
go run main.go
```
Isso iniciará o servidor na porta `8080`.

## Endpoints da API
1. Criar um Item
    - URL: _POST /items_

    - Descrição: Cria um novo item, realizando a validação do email.

    - Corpo da Requisição (JSON):
    ```bash
        {
    "name": "Nome do Item",
    "email": "exemplo@dominio.com"
    }
    ```
    - **Resposta** (Sucesso - Status 201):
    ```bash
    {
    "message": "Item criado com sucesso",
    "item": {
        "name": "Nome do Item",
        "email": "exemplo@dominio.com"
    }
    }
    ```
    - **Resposta** (Erro - Status 400 - Email inválido):
    ```bash
    {
    "error": "Email inválido"
    }
    ```bash

2. Listar Itens
    - URL: GET /items
    - Descrição: Lista todos os itens disponíveis.
    - Resposta (Sucesso - Status 200):
    ```bash
        {
        "items": ["Item1", "Item2", "Item3"]
        }
    ```

## Exemplo de Requisição no Postman
1. Criar Item
    1. - Abra o Postman e selecione POST.
    2. - Insira a URL: `http://localhost:8080/items`.
    3. - No corpo da requisição, adicione o JSON abaixo:

    ```bash
    {
    "name": "Novo Item",
    "email": "email@dominio.com"
    }
    ```

    4 - Envie a requisição e verifique a resposta com o status 201 Created.

2. Listar Itens
    1. - No Postman, selecione GET.
    2. - Insira a URL: http://localhost:8080/items.
    3. - Envie a requisição e verifique a resposta com o status `200 OK` e o corpo contendo a lista de itens.


## Funções Utilizadas

    Validação de Email: A função ValidateEmail é usada para garantir que o email fornecido ao criar um item seja válido.
    Logging: O projeto utiliza um sistema de logging para registrar informações importantes e erros no processo de execução da API.

## Como Validar o Email

O campo de email é validado pela função ValidateEmail na pasta utils/validation.go. Ela verifica se o email segue o padrão nome@dominio.com. Caso o email não seja válido, a API retorna um erro.

