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

package utils

import (
	"bytes"
	"os/exec"
)

type Git struct {
	Dir    string
	Stdout bytes.Buffer
	Stderr bytes.Buffer
}

type GitInterface interface {
	Run(command string, args ...string) error
	Clone(args ...string) error
	List(args ...string) error
	Checkout(args ...string) error
	Pull(args ...string) error
	Fetch(args ...string) error
	Branch(args ...string) error
	Log(args ...string) error
	Current() error
	Ping(url string) error
}

func (g *Git) Run(command string, args ...string) error {
	g.Stdout.Reset()
	g.Stderr.Reset()
	cmd := exec.Command("git", append([]string{command}, args...)...)

	if len(g.Dir) != 0 {
		cmd.Dir = g.Dir
	}
	cmd.Stdout = &g.Stdout
	cmd.Stderr = &g.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (g *Git) Clone(args ...string) error {
	return g.Run("clone", args...)
}

func (g *Git) List(args ...string) error {
	return g.Run("ls-remote", args...)
}

func (g *Git) Checkout(args ...string) error {
	return g.Run("checkout", args...)
}

func (g *Git) Pull(args ...string) error {
	return g.Run("pull", args...)
}

func (g *Git) Fetch(args ...string) error {
	return g.Run("fetch", args...)
}

func (g *Git) Log(args ...string) error {
	return g.Run("log", args...)
}

func (g *Git) Current() error {
	return g.Run("symbolic-ref", "--short", "HEAD")
}

func (g *Git) Ping(url string) error {
	return g.List("-h", url)
}
