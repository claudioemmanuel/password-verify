
## Tecnologia
#### `Back-end`
- [Go](https://go.dev)
#### Para executar a aplicação será necessário o [Docker](https://www.docker.com)
#### Para os testes na API REST será necessário o [Insomnia](https://insomnia.rest/download)

## Preparação do projeto

Clone o projeto
```bash
$ git clone https://github.com/claudioemmanuel/password-verify.git
```
Instalação do projeto
```bash
$ docker-compose up --build
```
Se tudo correr bem, o projeto deverá estar rodando em **localhost:8080**

## Testes
Para execução dos testes unitários, após o projeto estar rodando, basta entrar com o comando  

```bash
$ go test ./pkg/tests
```
e você deverá ver um retorno como ```ok  command-line-arguments  0.221s``` indicando que todos os testes rodaram com sucesso.

## Insomnia
Para os testes há um payload de exemplo no diretório **docs/payload.json**
Basta importá-lo no **Insomnia** para utilizar.

## Endpoint
### /verify ```POST``` 

Requisição REST contendo a senha e uma lista de regra
```bash
{
  "password": "TesteSenhaForte!123&",
  "rules": [
    {
      "rule": "minSize",
      "value": 8
    },
    {
      "rule": "minSpecialChars",
      "value": 2
    },
    {
      "rule": "noRepeted",
      "value": 0
    },
    {
      "rule": "minDigit",
      "value": 4
    }
  ]
}
```
A resposta deverá retornar um mapa com duas chaves:

```verify```: que deve retornar um ```boolean``` dizendo se a senha foi validada por todas as regras
```noMatch```: que deve retornar uma lista de ```strings``` contendo quais as regras a senha não passou, caso a senha atenda todas as regras uma lista vazia será retornada. 
#### Exemplo:
```
{
  "verify": false,
  "noMatch": [
    "minDigit"
  ]
}
```
## Senha válida

**Problema**: Dada uma palavra contínua, e um conjunto de regras, o programa precisa verificar se a senha é válida baseada nas regras pedidas. As regras possíveis são: 
```minSize```: ter pelo menos ```X``` caracteres 
```minUppercase```: ter pelo menos ```X``` caracteres maiúsculos
```minLowercase```: ter pelo menos ```X``` caracteres minúsculos
```minDigit```: ter pelo menos ```X``` dígitos (0-9)
```minSpecialChars```: ter pelo menos ```X``` caracteres especiais. Os caracteres especiais são: ```!@#$%^&*()-+\/{}[]```
```noRepeted```: não tenha nenhum caractere repetido em sequência. Exemplo: ```aab``` viola esta condição, mas  ```aba``` não

**Solução**: Baseado nas regras informadas foi desenvolvido um endpoint para validação da senha enviada pelo usuário.

## 📙 Licença
> Com base nos termos de [MIT LICENSE](https://opensource.org/licenses/MIT)

##### Feito por Claudio Emmanuel com ❤️
