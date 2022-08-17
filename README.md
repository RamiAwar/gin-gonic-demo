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
