package main

import (
	"net/http"

	"github.com/rodrities/lector-service/structure/endpoint"
	"github.com/rodrities/lector-service/structure/handler"
	"github.com/rodrities/lector-service/structure/service"
)

func main() {
	svc := service.NewPredictService()
	predictCovidEndpoint := endpoint.MakePredictCovidEndpoint(svc)
	handler.NewHttpHandler(predictCovidEndpoint)
	http.ListenAndServe(":8083", nil)

}
