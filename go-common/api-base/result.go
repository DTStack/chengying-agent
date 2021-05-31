package apibase

import (
	"net/http"
	"runtime/debug"

	"github.com/kataras/iris/context"
)

const (
	UNKNOWN_ERR = iota + 100
	API_PARAM_ERR
	DB_MODEL_ERR
	RPC_HANDLE_ERR
)

type ApiResult struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func Feedback(ctx context.Context, result interface{}) {
	if err, ok := result.(error); ok {
		if IsApiParameterErrors(err) {
			errs, _ := err.(*ApiParameterErrors)
			data := map[string]string{}
			for pname, err := range errs.errors {
				data[pname] = err.Error()
			}
			ctx.JSON(&ApiResult{
				Code: API_PARAM_ERR,
				Msg:  "请求参数错误: Invalid parameter(s)",
				Data: data,
			})
		} else if IsDBModelError(err) {
			e, _ := err.(*DBModelError)
			ctx.JSON(&ApiResult{
				Code: DB_MODEL_ERR,
				Msg:  "获取agent数据失败: DB Model error",
				Data: e.err.Error(),
			})
		} else if IsRpcHandleError(err) {
			e, _ := err.(*RpcHandleError)
			ctx.JSON(&ApiResult{
				Code: RPC_HANDLE_ERR,
				Msg:  "与agent通信失败: Rpc handle error",
				Data: e.err.Error(),
			})
		} else {
			debug.PrintStack()
			ctx.JSON(&ApiResult{
				Code: UNKNOWN_ERR,
				Msg:  err.Error(),
				Data: err.Error(),
			})
		}
	} else if str, ok := result.(string); ok {
		ctx.WriteString(str)
	} else if bin, ok := result.([]byte); ok {
		ctx.Write(bin)
	} else {
		ctx.JSON(&ApiResult{
			Code: 0,
			Msg:  "ok",
			Data: result,
		})
	}
	ctx.StatusCode(http.StatusOK)
	ctx.Done()
}
