package util

import (
	"fmt"
	"io"
	"os"
)

// copyFileContents copies the contents from source to the file defined by dest.
func copyFileContents(src, dest string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dest)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

func CopyFile(src, dest string) (err error) {
	srcFile, err := os.Stat(src)
	if err != nil {
		return
	}
	if !srcFile.Mode().IsRegular() {
		return fmt.Errorf("copyfile: could not copy the file %s", srcFile.Name())
	}
	destFile, err := os.Stat(dest)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(destFile.Mode().IsRegular()) {
			return fmt.Errorf("copyfile: could not dest file")
		}
		if os.SameFile(srcFile, destFile) {
			return
		}
	}
	if err = os.Link(src, dest); err == nil {
		return
	}
	err = copyFileContents(src, dest)
	return
}