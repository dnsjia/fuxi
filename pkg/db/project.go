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

type ProjectInterface interface {
	Create(ctx context.Context, p models.Project) (id int, err error)
	List(ctx context.Context) (project []*models.Project, err error)
	Get(ctx context.Context, id int) (project *models.Project, err error)
	Delete(ctx context.Context, id int) error
}

type project struct {
	db *gorm.DB
}

func NewProjectFactory(db *gorm.DB) ProjectInterface {
	return &project{
		db: db,
	}
}

func (p *project) Create(ctx context.Context, project models.Project) (id int, err error) {
	err = p.db.Model(&project).Create(&project).Error
	return project.Id, err

}

func (p *project) List(ctx context.Context) (project []*models.Project, err error) {
	err = p.db.Model(&models.Project{}).Find(&project).Error
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (p *project) Get(ctx context.Context, id int) (project *models.Project, err error) {
	err = p.db.Model(&models.Project{}).Where("id = ?", id).First(&project).Error
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (p *project) Delete(ctx context.Context, id int) error {
	return p.db.Model(&models.Project{}).Where("id = ?", id).Delete(&models.Project{}).Error
}
