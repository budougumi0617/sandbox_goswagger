package controller

import (
	"github.com/budougumi0617/sandbox_goswagger/gen/models"
	"log"

	"github.com/budougumi0617/sandbox_goswagger/gen/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func PostAPIRegisterHandler(params operations.PostAPIRegisterParams) middleware.Responder {
	body := params.User
	log.Printf("request body: %#v\n", body)

	// operations/*_responses.goに定義ができる。
	// 別のハンドラー用の定義を使えてしまうことに注意する。
	// return operations.NewPostAPIRegisterNotFound()
	return operations.NewPostAPIRegistOK()
}

func GetGreetingHandler(params operations.GetGreetingParams) middleware.Responder {
	req := params.HTTPRequest // 生Request
	log.Printf("METHOD from raw request: %q\n", req.Method)
	rsp := &models.HelloResponse{}
	if params.Name != nil {
		rsp.Name = *(params.Name)
	}

	return operations.NewGetGreetingOK().WithPayload(rsp)
}