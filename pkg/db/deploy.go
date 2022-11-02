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

package db

import (
	"context"

	"gorm.io/gorm"

	"github.com/dnsjia/fuxi/pkg/db/models"
)

type DeployInterface interface {
	Get(ctx context.Context, taskId string) (deploy *models.DeployHistory, err error)
	List(ctx context.Context) (deploy []*models.DeployHistory, err error)
	Create(ctx context.Context, deploy *models.DeployHistory) error
	Updates(ctx context.Context, taskId string, deploy models.DeployHistory) error
}

type deploy struct {
	db *gorm.DB
}

func NewDeployFactory(db *gorm.DB) DeployInterface {
	return &deploy{
		db: db,
	}
}

func (d *deploy) Get(ctx context.Context, taskId string) (deploy *models.DeployHistory, err error) {
	err = d.db.Model(&models.DeployHistory{}).Where("task_id = ?", taskId).First(deploy).Error
	if err != nil {
		return nil, err
	}
	return deploy, nil
}

func (d *deploy) List(ctx context.Context) (deploy []*models.DeployHistory, err error) {
	err = d.db.Model(&models.DeployHistory{}).Find(&deploy).Error
	if err != nil {
		return nil, err
	}
	return deploy, nil
}

func (d *deploy) Create(ctx context.Context, deploy *models.DeployHistory) error {
	return d.db.Model(&models.DeployHistory{}).Create(&deploy).Error
}

func (d *deploy) Updates(ctx context.Context, taskId string, deploy models.DeployHistory) error {
	return d.db.Model(&models.DeployHistory{}).Where("task_id = ?", taskId).Updates(deploy).Error
}
