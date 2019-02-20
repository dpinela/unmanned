package main

import (
	"context"
	"io"
	"os/exec"
)

func renderMandoc(ctx context.Context, w io.Writer, r io.Reader) error {
	proc := exec.CommandContext(ctx, "mandoc", "-T", "html", "-O", "man=/%S/%N,style=/man.css")
	proc.Stdin = r
	proc.Stdout = w
	return proc.Run()
}
