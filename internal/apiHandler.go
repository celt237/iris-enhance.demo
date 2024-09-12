package internal

import (
	"context"
	"fmt"
	"github.com/celt237/iris-enhance.demo/model"
	"github.com/kataras/iris/v12"
)

type ApiHandlerImpl struct {
}

func (a *ApiHandlerImpl) HandleCustomerAttribute(ctx iris.Context, attribute string, opt ...string) error {
	//TODO implement me
	//panic("implement me")
	if attribute == "@xPower" {
	}
	fmt.Sprintln("HandleCustomerAttribute:" + attribute)
	return nil
}

func (a *ApiHandlerImpl) WrapContext(ctx iris.Context) context.Context {
	return ctx
}

func (a *ApiHandlerImpl) HandleCustomerAnnotation(ctx iris.Context, annotation string, opt ...string) error {
	return nil
}

func (a *ApiHandlerImpl) Success(ctx iris.Context, produceType string, data interface{}) {
	result := &model.Result{
		Code:    200,
		Message: "success",
		Data:    data,
	}
	ctx.JSON(result)
	return
}
func (a *ApiHandlerImpl) CodeError(ctx iris.Context, produceType string, data interface{}, code int, err error) {
	result := &model.Result{
		Code:    code,
		Message: "fail",
		Data:    nil,
	}
	ctx.JSON(result)
	return
}
func (a *ApiHandlerImpl) Error(ctx iris.Context, produceType string, data interface{}, err error) {
	result := &model.Result{
		Code:    500,
		Message: "fail",
		Data:    nil,
	}
	ctx.JSON(result)
	return
}
