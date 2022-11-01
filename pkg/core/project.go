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
	"errors"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"

	"github.com/dnsjia/fuxi/api/types"
	"github.com/dnsjia/fuxi/cmd/app/config"
	"github.com/dnsjia/fuxi/pkg/db"
	"github.com/dnsjia/fuxi/pkg/db/models"
	fxtypes "github.com/dnsjia/fuxi/pkg/types"
	"github.com/dnsjia/fuxi/pkg/utils"
)

type ProjectGetter interface {
	Project() ProjectInterface
}

type ProjectInterface interface {
	Create(ctx context.Context, project models.Project) (id int, err error)
	List(ctx context.Context) (project []*models.Project, err error)
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (project *models.Project, err error)
	Ping(ctx context.Context, request types.PingRepoRequest) error
}

type project struct {
	ComponentConfig config.Config
	factory         db.ShareDaoFactory
	redis           *redis.Client
}

func newProject(fx *fuxi) ProjectInterface {
	return &project{
		ComponentConfig: fx.cfg,
		factory:         fx.factory,
		redis:           fx.redis,
	}
}

func (p *project) Create(ctx context.Context, project models.Project) (id int, err error) {
	return p.factory.Project().Create(ctx, project)
}

func (p *project) List(ctx context.Context) (project []*models.Project, err error) {
	return p.factory.Project().List(ctx)
}

func (p *project) Delete(ctx context.Context, id int) error {
	return p.factory.Project().Delete(ctx, id)
}

func (p *project) Get(ctx context.Context, id int) (project *models.Project, err error) {
	return p.factory.Project().Get(ctx, id)
}

func (p *project) Ping(ctx context.Context, request types.PingRepoRequest) error {
	protocol := strings.Split(request.URL, "://")[0]
	if protocol == fxtypes.RepoProtocolSSH {
		return errors.New("the ssh protocol is not supported")
	}

	var git utils.Git
	switch request.Type {
	case "git":
		return git.List(request.URL)
	case "svn":
	default:
		return fmt.Errorf("unsupported repository")
	}

	return nil

}
