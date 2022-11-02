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

package models

type DeployHistory struct {
	FuXiModel
	TaskId      string `gorm:"comment:任务ID" json:"taskId,omitempty"`
	ProjectId   int    `gorm:"comment:项目ID" json:"projectId"`
	ProjectName string `gorm:"comment:项目名称" json:"projectName"`
	Status      int    `gorm:"comment:部署状态" json:"status"`
	DeployUser  string `gorm:"comment:操作用户" json:"deployUser"`
	Detail      string `gorm:"comment:部署信息" gorm:"type:text" json:"detail"`
	Extension   string `gorm:"comment:扩展信息" gorm:"type:text" json:"extension"`
}

func (d DeployHistory) TableName() string {
	return "deploy_history"
}
