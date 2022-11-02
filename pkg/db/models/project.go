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

type Project struct {
	FuXiModel
	UserId                 int64  `gorm:"comment:用户ID" json:"userId"`
	ProjectName            string `gorm:"comment:应用名称" json:"projectName"`
	GitRepo                string `gorm:"comment:仓库地址" json:"gitRepo"`
	GitRepoType            uint8  `gorm:"comment:仓库类型" json:"gitRepoType"`
	GitBranch              string `gorm:"comment:仓库分支" json:"gitBranch"`
	DeployPath             string `gorm:"comment:部署路径" json:"deployPath"`
	DeployType             uint8  `gorm:"comment:部署类型" json:"deployType"`
	DeployNoticeType       uint8  `gorm:"comment:消息通知类型" json:"deployNoticeType"`
	DeployNoticeURL        string `gorm:"comment:消息通知Webhook地址" json:"deployNoticeURL"`
	HookPullCodeBefore     string `gorm:"comment:拉取代码之前" json:"hookPullCodeBefore"`
	HookPullCodeAfter      string `gorm:"comment:拉取代码之后" json:"hookPullCodeAfter"`
	HookPullCodeScriptType uint8  `gorm:"comment:脚本类型" json:"hookPullCodeScriptType"`
	HookDeployBefore       string `gorm:"comment:部署之前" json:"hookDeployBefore"`
	HookDeployAfter        string `gorm:"comment:部署之后" json:"hookDeployAfter"`
	HookDeployScriptType   uint8  `gorm:"comment:部署脚本类型" json:"hookDeployScriptType"`
	FilterRuleType         uint8  `gorm:"comment:过滤规则类型(contain包含，exclude排除)" json:"filterRuleType"`
	FilterRule             string `gorm:"comment:过滤规则" json:"filterRule"`
	DeployServers          string `gorm:"comment:服务器" json:"deployServers"`
}

func (p Project) TableName() string {
	return "project"
}
