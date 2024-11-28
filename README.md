# go-clima-cep-otel:  Observabilidade com OpenTelemetry e Zipkin

Este projeto consiste em dois serviços, A e B, para validação de CEP e obtenção de informações meteorológicas com base
na localização do CEP.

## Estrutura do Projeto

- **Serviço A**: Responsável por receber e validar o CEP.
- **Serviço B**: Responsável pela orquestração, validando o CEP, obtendo a localização e retornando informações
  meteorológicas formatadas.


## Iniciando

1. Clone o repositório:
    ```sh
    git clone https://github.com/ramonamorim/go-clima-cep-otel.git
    cd go-clima-cep-otel
    ```

2. Execute o comando abaixo na pasta raiz do projeto para iniciar o ambiente de desenvolvimento:

    ```sh
    docker-compose up -d
    ```

   Para parar os serviços:
    ```sh
    docker-compose down
    ```

3. Ou utilize os comandos make,
    ```sh
    make up
    ```
    Para parar os serviços:
    ```sh
    make stop
    ```




## Endpoints

### Serviço A

O **serviço A** estará rodando no endereço `http://localhost:8080/`, você pode enviar um cep valido no formato JSON. O
arquivo `api/requests.http` contém exemplos de uso.

Comportamento:

- **POST** `/`
    - Request Body:
      ```json
      {
        "cep": "89221220"
      }
      ```
    - Responses:
        - 200: Encaminha para o Serviço B.
        - 422: `invalid zipcode` caso seja inválido.

### Serviço B

Você pode acessar o serviço B em `http://localhost:8081/{cep}`. O arquivo `api/requests.http` contém exemplos de
uso.

- **GET** `/{cep}`
    - Responses:
        - 200: `{ "city": "São Paulo", "temp_C": 28.5, "temp_F": 83.3, "temp_K": 301.65 }`
        - 404: `can not find zipcode` caso não encontre o CEP.
        - 422: `invalid zipcode` caso o CEP seja inválido.

### Zipkin

Para acessar a telemetria use o seguinte endereço do `zipkin` e após realizar uma requisição clique no
botão "`RUN QUERY`":

- `http://localhost:9411/zipkin`


## Testes

Após iniciar o ambiente de desenvolvimento, você pode testar com o cURL de exemplo abaixo ou com o
arquivo `api/requests.http`:

```sh
curl -X POST http://localhost:8080/ -H "Content-Type: application/json" -d '{"cep": "89221220"}'
```

```sh
curl http://localhost:8081/89221220
```

