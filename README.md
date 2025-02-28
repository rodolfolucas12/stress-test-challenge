# stress-test-challenge

Este projeto é uma ferramenta de teste de estresse para serviços web. Ele permite que você envie um grande número de requisições simultâneas para um URL especificado e gera um relatório detalhado sobre o desempenho do serviço.

## Funcionalidades

- Envio de múltiplas requisições HTTP para um URL especificado.
- Configuração do número total de requisições e do nível de concorrência.
- Geração de relatórios com o tempo total de execução, número de requisições bem-sucedidas e distribuição de códigos de status HTTP.

## Requisitos

- Go 1.22.3 ou superior
- Docker (opcional, para execução em container)

## Instalação

### Localmente

1. Clone o repositório:
    ```sh
    git clone https://github.com/rodolfolucas12/stress-test-challenge.git
    cd stress-test-challenge
    ```

2. Instale as dependências:
    ```sh
    go mod tidy
    ```

3. Compile o projeto:
    ```sh
    go build -o stress-test-challenge ./cmd
    ```

4. Execute o programa:
    ```sh
    ./stress-test-challenge -url="http://example.com" -requests=100 -concurrency=10
    ```

### Usando Docker

1. Construa a imagem Docker:
    ```sh
    docker build -t stress-test-challenge .
    ```

2. Execute o container Docker:
    ```sh
    docker run --rm stress-test-challenge -url=http://example.com -requests=100 -concurrency=10
    ```

## Uso

### Parâmetros

- `-url`: URL do serviço a ser testado.
- `-requests`: Número total de requisições a serem enviadas.
- `-concurrency`: Número de requisições simultâneas.

### Exemplo de Uso

```sh
./stress-test-challenge -url="http://example.com" -requests=100 -concurrency=10