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

package core

import (
	"context"

	"github.com/go-redis/redis/v8"

	"github.com/dnsjia/fuxi/cmd/app/config"
	"github.com/dnsjia/fuxi/pkg/db"
	"github.com/dnsjia/fuxi/pkg/db/models"
)

type DeployGetter interface {
	Deploy() DeployInterface
}

type DeployInterface interface {
	Get(ctx context.Context, taskId string) (deploy *models.DeployHistory, err error)
	List(ctx context.Context) (deploy []*models.DeployHistory, err error)
	Create(ctx context.Context, deploy *models.DeployHistory) error
	Updates(ctx context.Context, taskId string, deploy models.DeployHistory) error
}

type deploy struct {
	ComponentConfig config.Config
	factory         db.ShareDaoFactory
	redis           *redis.Client
}

func newDeploy(fx *fuxi) DeployInterface {
	return &deploy{
		ComponentConfig: fx.cfg,
		factory:         fx.factory,
		redis:           fx.redis,
	}
}

func (d *deploy) Get(ctx context.Context, taskId string) (deploy *models.DeployHistory, err error) {
	return d.factory.Deploy().Get(ctx, taskId)
}

func (d *deploy) List(ctx context.Context) (deploy []*models.DeployHistory, err error) {
	return d.factory.Deploy().List(ctx)
}

func (d *deploy) Create(ctx context.Context, deploy *models.DeployHistory) error {
	return d.factory.Deploy().Create(ctx, deploy)
}

func (d *deploy) Updates(ctx context.Context, taskId string, deploy models.DeployHistory) error {
	return d.factory.Deploy().Updates(ctx, taskId, deploy)
}
