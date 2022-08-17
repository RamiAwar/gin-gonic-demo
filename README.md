# Gin Golang Demo Project

Opinionated example Golang web application with Gin Gonic.

Builds upon https://github.com/RamiAwar/go_demo_users_api.

## Build and run with Docker
```bash
docker build -t gin-demo:latest .
docker run -p 3000:3000 gin-demo
```


## Build and run with Go
```bash
go mod tidy
go run ./main.go
```


## DNS endpoint example:
```bash
❯ curl -X POST localhost:3000/v1/dns/find -H 'Content-Type: application/json' -d '{"x": "123.12", "y": "456.56", "z": "789.89", "vel": "20.0"}'
```


## Running the tests
```bash
go test ./...
```

## Understanding the structure
Check out my blog post on backend route design for an in depth dive into the structure: https://softgrade.org/in-depth-guide-to-backend-route-design/

### At a high level:
- routes: validation of requests, returning results, handling errors
- services: business logic, should not be concerned whatsoever with how they're used (ex. know nothing about HTTP or context)
- models: request models, response models, and everything in between goes here. Validation + default values are implemented as methods of those models.
- config, app: self explanatory

### Structural improvements:
Passing down context from requests to services is usually very useful. It is true that we want to isolate business logic from usecase, but in the case of web servers this is a very acceptable compromise as the context offers many advantages.
