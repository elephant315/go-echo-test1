## About

This project is a web application built with the Echo framework in Go, demonstrating hexagonal architecture principles. 
The application includes features such as unit tests, Docker deployment, and Swagger API documentation (localhost:8080/apidoc/index.html).


## Getting Started

Use a Makefile

```makefile
make test       ; Run unit tests locally
make build      ; Build a Docker image
make up         ; Run App Server in a Docker on 8080 port exposed
make down       ; Down App Server in a Docker
make doc       ; Make docs
```

## Swagger docs

Can be accessed on `http://[you_server_ip:port]/apidoc/index.html`

## Directory Structure
Here's a directory structure. `RFU` are folders reserved for future use.
```
/go-echo-test1
|-- /cmd
|   |-- /server           # Application entry points for the
|
|-- /internal
|   |-- /app              # Application logic
|   |   |-- /domain       # RFU - Domain types and interfaces
|   |   |-- /ports        # RFU - Interfaces for primary and secondary adapters
|   |   |-- /services     # Business logic
|   |
|   |-- /adapters         # Adapters for primary and secondary ports
|       |-- /http         # HTTP server setup and handlers
|
|-- /pkg                  # RFU - Additional packages 
|-- /configs              # RFU - Configuration files
|-- /scripts              # RFU - Deployment, build scripts
|-- /test                 # RFU - Test files
```
