package api

import (
	"net/http"

	"mono.thienhang.com/pkg/context"
)

func Ok(ctx *context.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"msg":  "ok",
	})
}

func OkWithMsg(ctx *context.Context, msg string) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"msg":  msg,
	})
}

func OkWithData(ctx *context.Context, data map[string]interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": data,
	})
}

func OkWithDataRaw(ctx *context.Context, data interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": data,
	})
}

func BadRequest(ctx *context.Context, msg string) {
	ctx.JSON(http.StatusBadRequest, map[string]interface{}{
		"code": http.StatusBadRequest,
		"msg":  msg,
	})
}

func Error(ctx *context.Context, msg string, datas ...map[string]interface{}) {
	res := map[string]interface{}{
		"code": http.StatusInternalServerError,
		"msg":  msg,
	}
	if len(datas) > 0 {
		res["data"] = datas[0]
	}
	ctx.JSON(http.StatusInternalServerError, res)
}

func Denied(ctx *context.Context, msg string) {
	ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
		"code": http.StatusForbidden,
		"msg":  msg,
	})
}
