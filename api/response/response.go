/*
Copyright (c) 2022 The DnsJia Authors.
WebSite:  https://github.com/dnsjia/fuxi
Email:    OpenSource@dnsjia.com

MIT License

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code   int         `json:"errCode"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	ErrMsg string      `json:"errMsg"`
}

const (
	SUCCESS    = 0
	ERROR      = -1
	ParamError = 8000
)

const (
	OkMsg         = "操作成功"
	NotOkMsg      = "操作失败"
	ParamErrorMsg = "参数绑定失败, 请检查数据类型"
)

var CustomError = map[int]string{
	SUCCESS:    OkMsg,
	ERROR:      NotOkMsg,
	ParamError: ParamErrorMsg,
}

func ResultFail(code int, data interface{}, msg string, c *gin.Context) {
	if msg == "" {
		c.JSON(http.StatusOK, Response{
			Code:   code,
			Data:   data,
			ErrMsg: CustomError[code],
		})
	} else {
		c.JSON(http.StatusOK, Response{
			Code:   code,
			Data:   data,
			ErrMsg: msg,
		})
	}
}

func ResultOk(code int, data interface{}, msg string, c *gin.Context) {

	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(c *gin.Context) {
	ResultOk(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	ResultOk(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	ResultOk(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	ResultOk(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	ResultFail(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(code int, message string, c *gin.Context) {
	ResultFail(code, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, code int, message string, c *gin.Context) {
	ResultFail(code, data, message, c)
}
