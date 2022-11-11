package internal

import (
	"context"
	"fmt"
	"github.com/Code-Hex/pget"
	"github.com/GoFarsi/gvm/cli"
	"github.com/GoFarsi/gvm/errors"
	"log"
	"net/http"
	"os"
	"runtime"
)

type Download struct {
	version      string
	downloadPath string
	backup       bool
}

type Target struct {
	FileName      string
	Path          string
	ContentLength int64
	Url           string
	backup        bool
}

func newDownload(ver string, backup bool) (*Download, error) {
	def, err := defaultCfg(backup)
	if err != nil {
		return nil, err
	}

	if len(ver) != 0 {
		def.version = ver
	}

	return def, nil
}

func defaultCfg(backup bool) (*Download, error) {
	list, err := NewList()
	if err != nil {
		return nil, err
	}

	dl := &Download{
		version:      list.LastVersion(),
		downloadPath: os.TempDir(),
		backup:       backup,
	}

	if backup {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		dl.downloadPath = home
	}

	return dl, nil
}

func (d *Download) download(ctx context.Context) (*Target, error) {
	target := d.getFileInfo(ctx, fmt.Sprintf(filePattern, d.version, runtime.GOARCH))
	if target == nil {
		return nil, errors.ERR_CANT_FIND_ACTIVE_MIRROR
	}

	target.Path = d.downloadPath
	target.backup = d.backup

	if err := target.download(ctx); err != nil {
		return nil, err
	}

	return target, nil
}

func (*Download) getFileInfo(ctx context.Context, fileName string) *Target {
	for _, v := range mirrors {
		client := cli.DefaultClient()
		client.Timeout = _defaultMirrorTimeout
		req, err := http.NewRequest("HEAD", v+fileName, nil)
		if err != nil {
			continue
		}

		req = req.WithContext(ctx)

		resp, err := client.Do(req)
		if err != nil {
			continue
		}

		if resp.Header.Get("Accept-Ranges") != "bytes" && resp.ContentLength <= 0 {
			continue
		}

		return &Target{
			FileName:      fileName,
			ContentLength: resp.ContentLength,
			Url:           v + fileName,
		}
	}

	return nil
}

func (t *Target) download(ctx context.Context) error {
	log.Printf("started download %s with size %d MB...", t.FileName, t.ContentLength/1024/1024)
	dl := pget.New()
	args := []string{
		"-o",
		t.Path,
		"-t",
		"30",
		t.Url,
	}
	if err := dl.Run(ctx, "", args); err != nil {
		return err
	}

	if t.backup {
		log.Printf("download file %s completed and backed up in %s path.", t.FileName, t.Path)
		return nil
	}

	log.Printf("download file %s completed.", t.FileName)
	return nil
}
