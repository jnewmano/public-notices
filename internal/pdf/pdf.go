package pdf

import (
	"context"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func ExtractText(ctx context.Context, name string, r io.Reader) (string, error) {

	dir, err := ioutil.TempDir("", "pdftotext")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(dir)

	srcFilename := filepath.Join(dir, "source.pdf")
	dstFilename := filepath.Join(dir, "source.txt")

	err = WriteFile(srcFilename, r, 0744)
	if err != nil {
		return "", err
	}

	// Command: pdftotext -layout -nopgbrk <fn.pdf> <output.txt>
	args := []string{
		"-layout",
		"-nopgbrk",
		srcFilename,
		dstFilename,
	}

	cmd := exec.CommandContext(ctx, "pdftotext", args...)
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadFile(dstFilename)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func WriteFile(filename string, data io.Reader, perm os.FileMode) error {

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, data)
	if err1 := f.Close(); err == nil {
		err = err1
	}

	return err
}
