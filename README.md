[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

# SKey

## Dependencies

"github.com/stretchr/testify/assert"
"github.com/google/uuid"
## Test

go test -cover ./...
go test ./... -coverprofile docs/testcover.out
go tool cover -html=docs/testcover.out

## Tech Stack

![Go](https://img.shields.io/badge/Go-v1.19-blue)
![Go](https://img.shields.io/badge/Go-Migrate-blue)
![Go](https://img.shields.io/badge/Go-Viper-blue)
![Go](https://img.shields.io/badge/Go-Validator-blue)
![Go](https://img.shields.io/badge/Go-Mock-blue)
![Go](https://img.shields.io/badge/Go-Gin-blue)
![Go](https://img.shields.io/badge/Go-Testify-blue)
![Docker](https://img.shields.io/badge/Docker-passing-green)
![Aws RDS](https://img.shields.io/badge/Aws-Rds-blue)
![github](https://img.shields.io/badge/Github-Actions-green)
![API](https://img.shields.io/badge/API-http-blue)
![Test](https://img.shields.io/badge/Test-unit-green)
![CI/CD](https://img.shields.io/badge/CI%20CD-automation-green)
![Go](https://img.shields.io/badge/Go-Paseto-green)
![Go](https://img.shields.io/badge/Go-Jwt-green)
![Postgresql](https://img.shields.io/badge/Go-Pq-green)

## License

[MIT](https://choosealicense.com/licenses/mit/)


## Developers:

<img src="https://github.com/husamettinarabaci/husamettinarabaci/blob/main/hsmtek-logo.png?raw=true" width="200"/>

[@husamettinarabaci](https://www.github.com/husamettinarabaci)

## Notes
DDD, Hex-Arc, Unit Test,Google Wire, testify

## Project Structure
Domain 
    Module Based
        Model 
            Value Object
            Entity
            Aggregate
        Service

Application
    Service
        Module Based
            C/Q Handlers
            Event Listeners
    Port - (Driver(Primary),Driven(Secondary))
    Adapter - (Driver(Primary))

Infrastructure
    Adapter - (Driven(Secondary))

Presentation

## Directory Structure