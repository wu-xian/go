// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// go work init

package workcmd

import (
	"cmd/go/internal/base"
	"cmd/go/internal/modload"
	"context"
	"path/filepath"
)

// TODO(#49232) Add more documentation below. Though this is
// enough for those trying workspaces out, there should be more through
// documentation before Go 1.18 is released.

var cmdInit = &base.Command{
	UsageLine: "go work init [moddirs]",
	Short:     "initialize workspace file",
	Long: `Init initializes and writes a new go.work file in the current
directory, in effect creating a new workspace at the current directory.

go work init optionally accepts paths to the workspace modules as arguments.
If the argument is omitted, an empty workspace with no modules will be created.

See the workspaces design proposal at
https://go.googlesource.com/proposal/+/master/design/45713-workspace.md for
more information.
`,
	Run: runInit,
}

func init() {
	base.AddModCommonFlags(&cmdInit.Flag)
	base.AddWorkfileFlag(&cmdInit.Flag)
}

func runInit(ctx context.Context, cmd *base.Command, args []string) {
	modload.InitWorkfile()

	modload.ForceUseModules = true

	// TODO(matloob): support using the -workfile path
	// To do that properly, we'll have to make the module directories
	// make dirs relative to workFile path before adding the paths to
	// the directory entries

	workFile := filepath.Join(base.Cwd(), "go.work")

	modload.CreateWorkFile(ctx, workFile, args)
}
