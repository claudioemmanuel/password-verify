## Password Verify

This repository contains a simple API REST to verify if a password is valid based on a set of rules. This project was developed for job test application.

## Features

- API REST
- Unit tests
- Docker

## Requirements

- Go 1.19
- Docker
- Docker Compose

## Getting Started

1. Clone the project

```bash
$ git clone https://github.com/claudioemmanuel/password-verify.git
cd password-verify
```

2. Build and run the Docker container

```bash
$ docker-compose up --build
```

The API should now be running on `http://localhost:8080`.

## Running Tests

For Insomnia tests, import the **docs/payload.json** file or run the following command:

```bash
go test -v ./...
```

## The problem

Given a continuous word, and a set of rules, the program needs to verify if the password is valid based on the rules requested. The possible rules are:
`minSize`: have at least `X` characters,
`minUppercase`: have at least `X` uppercase characters,
`minLowercase`: have at least `X` lowercase characters,
`minDigit`: have at least `X` digits (0-9),
`minSpecialChars`: have at least `X` special characters. The special characters are: `!@#$%^&*()-+\/{}[]`,
`noRepeted`: do not have any repeated character in sequence. Example: `aab` violates this condition, but `aba` does not.

## The solution

Based on the rules informed, an endpoint was developed to validate the password sent by the user.

## License

This project is licensed under the MIT License - see the [MIT LICENSE](https://opensource.org/licenses/MIT) for details.

## Acknowledgments

- [Golang](https://golang.org/)
- [Gin](https://gin-gonic.com)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Insomnia](https://insomnia.rest/download)

## Author

- [Claudio Emanuel](https://www.linkedin.com/in/claudio-emmanuel/) made with ❤️
