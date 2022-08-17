# gin-gonic-demo

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
