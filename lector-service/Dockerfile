FROM golang:alpine
#renombrar
COPY ./apigateway/main.go ./lector-service.go
#port
EXPOSE 9000
# run api-svc001.go
CMD ["go", "run", "lector-service.go"]