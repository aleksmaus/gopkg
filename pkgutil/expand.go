package pkgutil

import (
	"compress/gzip"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/aleksmaus/gopkg/cpio"
	"github.com/aleksmaus/gopkg/xar"
)

var ErrPayloadNotFound = errors.New("pkg Payload file not found")

func Expand(pkgFile, dstDir string) error {
	f, err := os.Open(pkgFile)
	if err != nil {
		return err
	}
	defer f.Close()

	xr, err := xar.NewReader(f)

	if err != nil {
		return err
	}

	var payloadFile *xar.File
	for i := 0; i < len(xr.Files); i++ {
		f := xr.Files[i]
		if f.Name == "Payload" {
			payloadFile = &f
			break
		}
	}

	if payloadFile == nil {
		return ErrPayloadNotFound
	}

	return expandPayload(payloadFile, dstDir)
}

func expandPayload(f *xar.File, dstDir string) error {
	gzr, err := gzip.NewReader(f.Body)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dstDir, 0755)
	if err != nil {
		return err
	}

	r := cpio.NewReader(gzr)
	var entry cpio.Entry
	for entry, err = r.Next(); err == nil; entry, err = r.Next() {
		if entry.FilePath == "." {
			continue
		}

		body := entry.Body
		// Ignore symlinks, not needed for our purposes
		if entry.FileMode.IsDir() {
			err = os.MkdirAll(filepath.Join(dstDir, entry.FilePath), entry.FileMode.Perm())
			if err != nil {
				return err
			}
		} else if entry.FileMode.IsRegular() {
			body = nil
			err := writeFile(dstDir, &entry)
			if err != nil {
				return err
			}
		}

		if body != nil {
			// Discarding the body, need to read in full
			io.Copy(io.Discard, entry.Body)
		}
	}

	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	return nil
}

func writeFile(dstDir string, entry *cpio.Entry) error {
	fp := filepath.Join(dstDir, entry.FilePath)
	err := writeFileContent(fp, entry.Body, entry.FileMode.Perm())
	if err != nil {
		return err
	}

	return os.Chtimes(fp, time.Now().Local(), entry.FileMtime)
}

func writeFileContent(fp string, r io.Reader, mode fs.FileMode) error {
	f, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, r)
	if err != nil {
		return err
	}
	return nil
}
