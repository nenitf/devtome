# devtome

[![emojicom](https://img.shields.io/badge/emojicom-%F0%9F%90%9B%20%F0%9F%86%95%20%F0%9F%92%AF%20%F0%9F%91%AE%20%F0%9F%86%98%20%F0%9F%92%A4-%23fff)](http://neni.dev/emojicom) [![CI](https://github.com/nenitf/devtome/actions/workflows/ci.yml/badge.svg)](https://github.com/nenitf/devtome/actions/workflows/ci.yml)

Projeto para manter backups dos próprios artigos publicados no [dev.to](https://dev.to/).

## Utilização

1. Crie uma [API Key](https://dev.to/settings/account)
2. Baixe o [executável](https://github.com/nenitf/devtome/releases)
3. Crie/edite o arquivo `config.txt` com referência no arquivo `config.example.txt`
4. Execute o binário

## Desenvolvimento

### Ambiente

#### Primeira inicialização

```sh
docker-compose up -d
docker-compose exec app go mod tidy
```

#### Demais inicializações

```sh
docker-compose up -d
```

### Teste

```sh
docker-compose exec app go test ./...
#docker-compose exec app go test ./pkg/pkgname
#docker-compose exec app go test ./pkg/pkgname/pkgname_test.go
#docker-compose exec app go test ./pkg/pkgname -run "regexOfTest$"
```

### Linting

```sh
docker-compose exec app go fmt ./...
```

## Créditos

### Desenvolvimento

- https://developers.forem.com/api/
- https://dev.to/albertodeago88/learn-golang-basics-by-creating-a-file-counter-50f1
- https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/reading-files
- https://medium.com/@flaqueEau/how-to-cleanly-test-files-in-go-af025bfcd9d3
- https://docs.docker.com/language/golang/run-tests/
- https://firehydrant.io/blog/develop-a-go-app-with-docker-compose/
- https://golang.cafe/blog/golang-httptest-example.html
- https://dev.to/eminetto/testing-apis-in-golang-using-apitest-1860
- https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html
