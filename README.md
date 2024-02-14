# Esquilo Aniquilador

![GitHub top language](https://img.shields.io/github/languages/top/kauefraga/esquilo-aniquilador)
![Rinha de Backend](https://img.shields.io/badge/Rinha_de-Backend-8A2BE2)
![GitHub's license](https://img.shields.io/github/license/kauefraga/esquilo-aniquilador)
![GitHub last commit (branch)](https://img.shields.io/github/last-commit/kauefraga/esquilo-aniquilador/main)

> 🐿 Minha API aniquiladora pra segunda edição da Rinha de Backend em Golang. A Rinha de Backend é uma competição muito divertida e, para mim, toda edição é/será de muito aprendizado. Participa aí!

> [!TIP]
> Veja [o repositório da Rinha](https://github.com/zanfranceschi/rinha-de-backend-2024-q1). O prazo final é 10/03/2024.

## 🗺 O que foi implementado

Inicialmente:

- [x] Rota `GET /`, Hello worldzinho 😃
- [x] Rota `POST /clientes/:id/transacoes` com banco de dados em memória (para testes e validação das regras de negócio)
- [ ] Rota `GET /clientes/:id/extrato` com banco de dados em memória

Posteriormente/atualmente:

- [ ] Rota `GET /clientes/:id/extrato` com banco de dados PostgreSQL
- [ ] Rota `POST /clientes/:id/transacoes` com Postgres

## ⬇ Como instalar e botar pra fu...ncionar

1. Clone o repositório
2. Rode `go run cmd/api/main.go`

O segundo passo já deve instalar as dependências. Caso contrário, rode `go mod download` e execute a segunda instrução de novo.

```bash
# (1)
git clone https://github.com/kauefraga/esquilo-aniquilador.git
cd esquilo-aniquilador

# (2)
go run cmd/api/main.go

# (3?)
go mod download
```

## 🧪 Como rodar os testes

### Testes Gatling

Ainda não testei.

### Testes unitários

Ainda não escrevi 🤡.

### Testes para verificar o funcionamento durante o desenvolvimento

Requisitos: [Visual Studio Code](https://code.visualstudio.com).

1. Instale a extensão [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)
2. Abra o arquivo `requests.http`
3. Clique em `Send Request`, faça isso para cada requisição descrita/que aparece

```bash
# (1)
code --install-extension humao.rest-client

# (2)
code . # Esteja dentro do diretório esquilo-aniquilador (raiz), que você clonou
```

## 🧙‍♂️ O que tunei e como

Não tunei nada, mal implementei o necessário 🤣.

## 🌳 Interações

### Redes Sociais

- [Meu Twitter/X](https://twitter.com/rkauefraga)
- [Rinha de Backend](https://twitter.com/rinhadebackend)

### Meus tweets/xweets

- [Início](https://twitter.com/rkauefraga/status/1757072132729639271)
- [Sobre a regra de negócio que retorna 422](https://twitter.com/rkauefraga/status/1757524333629464861)

## 📝 Licença

Este projeto está sob licença do MIT - Veja a [LICENÇA](https://github.com/kauefraga/esquilo-aniquilador/blob/main/LICENSE) para mais informações.

---

Feito com ❤ por Kauê Fraga Rodrigues.
