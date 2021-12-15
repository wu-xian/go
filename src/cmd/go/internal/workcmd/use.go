// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// go work use

package workcmd

import (
	"cmd/go/internal/base"
	"cmd/go/internal/fsys"
	"cmd/go/internal/modload"
	"context"
	"io/fs"
	"os"
	"path/filepath"
)

// TODO(#49232) Add more documentation below. Though this is
// enough for those trying workspaces out, there should be more thorough
// documentation before Go 1.18 is released.

var cmdUse = &base.Command{
	UsageLine: "go work use [-r] [moddirs]",
	Short:     "add modules to workspace file",
	Long: `Use provides a command-line interface for adding directories,
optionally recursively, to a go.work file.

The -r flag searches recursively for modules in the argument directories.`,
}

var useR = cmdUse.Flag.Bool("r", false, "")

func init() {
	cmdUse.Run = runUse // break init cycle

	base.AddModCommonFlags(&cmdUse.Flag)
	base.AddWorkfileFlag(&cmdUse.Flag)
}

func runUse(ctx context.Context, cmd *base.Command, args []string) {
	modload.InitWorkfile()

	modload.ForceUseModules = true

	var gowork string
	modload.InitWorkfile()
	gowork = modload.WorkFilePath()

	workFile, err := modload.ReadWorkFile(gowork)
	if err != nil {
		base.Fatalf("go: %v", err)
	}

	haveDirs := make(map[string]bool)
	for _, dir := range workFile.Use {
		haveDirs[filepath.Join(filepath.Dir(gowork), filepath.FromSlash(dir.Path))] = true
	}

	addDirs := make(map[string]bool)
	removeDirs := make(map[string]bool)
	lookDir := func(dir string) {
		absDir := filepath.Join(base.Cwd(), dir)
		// If the path is absolute, keep it absolute. If it's relative,
		// make it relative to the go.work file rather than the working directory.
		if !filepath.IsAbs(dir) {
			rel, err := filepath.Rel(filepath.Dir(gowork), absDir)
			if err == nil {
				dir = rel
			}
		}
		fi, err := os.Stat(filepath.Join(dir, "go.mod"))
		if err != nil {
			if os.IsNotExist(err) {

				if haveDirs[absDir] {
					removeDirs[dir] = true
				}
				return
			}
			base.Errorf("go: %v", err)
		}

		if !fi.Mode().IsRegular() {
			base.Errorf("go: %v is not regular", filepath.Join(dir, "go.mod"))
		}

		if !haveDirs[absDir] {
			addDirs[dir] = true
		}
	}

	for _, useDir := range args {
		if *useR {
			fsys.Walk(useDir, func(path string, info fs.FileInfo, err error) error {
				if !info.IsDir() {
					return nil
				}
				lookDir(path)
				return nil
			})
			continue
		}
		lookDir(useDir)
	}

	for dir := range removeDirs {
		workFile.DropUse(filepath.ToSlash(dir))
	}
	for dir := range addDirs {
		workFile.AddUse(filepath.ToSlash(dir), "")
	}
	modload.UpdateWorkFile(workFile)
	modload.WriteWorkFile(gowork, workFile)
}
