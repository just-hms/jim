package io

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"io"
	"io/ioutil"
	"jim/internal/constants"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ExecutableFolder() string {
	path, _ := os.Executable()
	path, _ = filepath.EvalSymlinks(path)

	return filepath.Dir(path)
}

func Executable() string {
	path, _ := os.Executable()
	path, _ = filepath.EvalSymlinks(path)

	return path
}

func CurrentFolder() string {
	path, _ := os.Getwd()
	return path
}

func ReadChar() rune {

	b := make([]byte, 1)
	os.Stdin.Read(b)

	if rune(string(b)[0]) == 13 {
		return 'y'
	}

	return rune(string(b)[0])
}

func ReplaceCurrentFolderFlag(input string) string {
	return strings.Replace(input, constants.CURRENT_FOLDER_FLAG, CurrentFolder(), -1)
}

func FileInput(file_default_content string) (string, error) {

	tmpDir := os.TempDir()
	tmpFile, tmpFileErr := ioutil.TempFile(tmpDir, "command")

	if tmpFileErr != nil {
		return "", errors.New("error while creating tempFile")
	}

	// set the file content to file_default_content
	tmpFile.WriteString(file_default_content)
	tmpFile.Close()

	// get the default editor
	editor := GetDefaultTextEditor()

	// get the default editor path
	path, err := exec.LookPath(editor)

	if err != nil {
		return "", errors.New("text editor not found")
	}

	// edit the tmp file with the default editor
	cmd := exec.Command(path, tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()

	if err != nil {
		return "", errors.New("text editor start failed")
	}

	cmd.Wait()

	data, err := os.ReadFile(tmpFile.Name())

	if err != nil {
		return "", errors.New("error reading the text file output")
	}

	return string(data), nil

}

func InterceptStdout(callback func()) string {

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	callback()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	return strings.TrimSpace(string(out))
}

func Untar(dst string, r io.Reader) error {

	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()

		switch {

		// if no more files are found return
		case err == io.EOF:
			return nil

		// return any other error
		case err != nil:
			return err

		// if the header is nil, just skip it (not sure how this happens)
		case header == nil:
			continue
		}

		// the target location where the dir/file should be created
		target := filepath.Join(dst, header.Name)

		// the following switch could also be done using fi.Mode(), not sure if there
		// a benefit of using one vs. the other.
		// fi := header.FileInfo()

		// check the file type
		switch header.Typeflag {

		// if its a dir and it doesn't exist create it
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}

		// if it's a file create it
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			// copy over contents
			if _, err := io.Copy(f, tr); err != nil {
				return err
			}

			// manually close here after each file operation; defering would cause each file close
			// to wait until all operations have completed.
			f.Close()
		}
	}
}
