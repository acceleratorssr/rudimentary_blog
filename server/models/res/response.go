package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/utils"
)

const error_ = 9

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

//type ListResponse struct {
//	Count int   `json:"count"`
//	List  []any `json:"list"`
//}

// ListResponse 使用泛型，安全性更高：
type ListResponse[T any] struct {
	TotalCount int `json:"count"`
	List       T   `json:"list"`
}

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func OKWithAll(data any, msg string, c *gin.Context) {
	Result(http.StatusOK, data, msg, c)
}

func OKWithData(data any, c *gin.Context) {
	Result(http.StatusOK, data, "success", c)
}

func OKWithMessage(msg string, c *gin.Context) {
	Result(http.StatusOK, nil, msg, c)
}

func OK(c *gin.Context) {
	Result(http.StatusOK, nil, "success", c)
}

func OKWithList[T any](list T, totalCount int, c *gin.Context) {
	OKWithData(ListResponse[T]{
		List:       list,
		TotalCount: totalCount,
	}, c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(error_, data, msg, c)
}

func FailWithCode(code ErrorCode, c *gin.Context) {
	// go的类型是强类型
	// 即便底层类型相同，不同类型的值仍被认为是不同的
	// 所以此处需要类型转换
	msg, ok := ErrorMap[code]
	if !ok {
		// 处理错误，例如返回一个默认错误消息
		msg = "Unknown error"
	}

	Result(int(code), nil, msg, c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(error_, nil, msg, c)
}

func FailWithError(err error, obj any, c *gin.Context) {
	msg := utils.GetValidMsg(err, obj)
	FailWithMessage(msg, c)
}
