# Go Gin Boilerplate
> A boilerplate for RESTful API with Golang, Gin and PostgresSQL

[![Go Version][go-image]][go-url]
[![License][license-image]][license-url]

Golang Gin boilerplate with Postgres database resource. Supports multiple configuration environments.

![](golang.png)

## Installation
First, you need to implement your own environment by copying `.example.env` into `.env`.
Then replacing with your own database's information.
```sh
cp .env.example .env
```

Run in your terminal: 
```sh
go run main.go
```

## Usage example
Example call: `http://localhost:8080/api/auth/login`

## Release History

* version 0.0.1
    * Configuration by environment, 
    * Authentication 
    * Entities: User, Token, Question, Score board
   

[go-image]: https://img.shields.io/badge/Go--version-1.18-blue.svg
[go-url]: https://golang.org/doc/go1.18
[license-image]: https://img.shields.io/badge/License-MIT-blue.svg
[license-url]: https://github.com/tienanh129902/go-rest-api/blob/master/LICENSE


