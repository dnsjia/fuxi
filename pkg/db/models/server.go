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

type Server struct {
	FuXiModel `json:"fuXiModel"`
	Name      string `gorm:"comment:主机名称" json:"name" json:"name"`
	Ip        string `gorm:"comment:主机IP" json:"ip,omitempty"`
	Port      int    `gorm:"comment:端口" json:"port"`
	UserName  string `gorm:"comment:登陆用户" json:"username"`
	Password  string `gorm:"comment:登陆密码" json:"password"`
	Os        string `gorm:"comment:操作系统" json:"os"`
	OsType    string `gorm:"comment:操作系统版本" json:"osType"`
	Desc      string `gorm:"comment:描述" gorm:"type:text" json:"desc"`
}

func (s Server) TableName() string {
	return "servers"
}
