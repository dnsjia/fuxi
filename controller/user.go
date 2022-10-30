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

package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dnsjia/fuxi/api/response"
	"github.com/dnsjia/fuxi/api/types"
	"github.com/dnsjia/fuxi/pkg/fuxi"
)

func Login(c *gin.Context) {
	var u types.User

	if err := response.CheckParams(c, &u); err != nil {
		fmt.Println(err.Error())
		return
	}
	user, err := fuxi.CoreV1.User().GetByUserName(context.TODO(), u.UserName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error(), "code": 1000})
		return
	}
	if !user.Status {
		c.JSON(http.StatusOK, gin.H{"err": "用户已经禁用", "code": 1000})
		return
	}
}
