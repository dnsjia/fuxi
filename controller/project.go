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
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dnsjia/fuxi/api/response"
	"github.com/dnsjia/fuxi/api/types"
	"github.com/dnsjia/fuxi/pkg/db/models"
	"github.com/dnsjia/fuxi/pkg/fuxi"
)

func CreateProject(c *gin.Context) {
	var project models.Project
	if err := response.CheckParams(c, &project); err != nil {
		response.FailWithMessage(response.ParamError, response.ParamErrorMsg, c)
		return
	}

	if _, err := fuxi.CoreV1.Project().Create(c.Request.Context(), project); err != nil {
		response.FailWithMessage(http.StatusOK, err.Error(), c)
		return
	}

	response.Ok(c)
}

func ListProject(c *gin.Context) {
	projectList, err := fuxi.CoreV1.Project().List(c.Request.Context())
	if err != nil {
		response.FailWithMessage(http.StatusOK, err.Error(), c)
		return
	}

	response.OkWithData(projectList, c)
}

func GetProject(c *gin.Context) {
	var projectOptions types.ProjectOptions
	if err := c.ShouldBindUri(&projectOptions); err != nil {
		response.FailWithMessage(response.ParamError, response.ParamErrorMsg, c)
		return
	}

	project, err := fuxi.CoreV1.Project().Get(c.Request.Context(), projectOptions.ProjectId)
	if err != nil {
		response.FailWithMessage(http.StatusOK, err.Error(), c)
		return
	}

	response.OkWithData(project, c)

}

func DeleteProject(c *gin.Context) {
	var req types.ProjectRequest
	if err := response.CheckParams(c, &req); err != nil {
		response.FailWithMessage(response.ParamError, response.ParamErrorMsg, c)
		return
	}

	if err := fuxi.CoreV1.Project().Delete(c.Request.Context(), req.ProjectId); err != nil {
		response.FailWithMessage(http.StatusOK, err.Error(), c)
		return
	}

	response.Ok(c)
}
