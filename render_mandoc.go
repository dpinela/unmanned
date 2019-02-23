package main

import (
	"context"
	"io"
	"os/exec"
)

func renderMandoc(ctx context.Context, basedir string, w io.Writer, r io.Reader) error {
	proc := exec.CommandContext(ctx, "mandoc", "-T", "html", "-O", "man=/%S/%N,style=/mandoc.css")
	proc.Stdin = r
	proc.Stdout = w
	proc.Dir = basedir
	return proc.Run()
}
