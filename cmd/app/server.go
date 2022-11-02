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

package app

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/dnsjia/fuxi/api/routers"
	"github.com/dnsjia/fuxi/cmd/app/options"
	"github.com/dnsjia/fuxi/pkg/fuxi"
)

func NewFuXiServerCommand() *cobra.Command {
	opts, err := options.NewOptions()
	if err != nil {
		fmt.Println("unable to initialize command options: ", zap.Any("err", err))
	}

	cmd := &cobra.Command{
		Use:   "fuxi",
		Short: "伏羲",
		Long: `欢迎大家加入我们,一起共创社区。 https://github.com/dnsjia/fuxi

		 _______     _    _ _ 
		(_______)   \ \  / (_)
		 _____ _   _ \ \/ / _ 
		|  ___) | | | )  ( | |
		| |   | |_| |/ /\ \| |
		|_|    \____/_/  \_\_|

`,

		Run: func(cmd *cobra.Command, args []string) {
			if err = opts.Complete(); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			if err := Run(opts); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}
	opts.BindFlags(cmd)

	return cmd
}

func Run(opts *options.Options) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fuxi.Setup(opts)

	initRouters(opts)
	opts.Run(ctx.Done())

	select {}
}

func initRouters(opts *options.Options) {
	routers.UserRouter(opts.GinEngine)
	routers.ProjectRouter(opts.GinEngine)
	routers.DeployRouter(opts.GinEngine)
}
