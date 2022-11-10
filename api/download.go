package api

import (
	"context"
	"fmt"
	"github.com/Code-Hex/pget"
	"github.com/GoFarsi/gvm/cli"
	"net/http"
	"os"
	"runtime"
)

type Download struct {
	version      string
	downloadPath string
}

type Target struct {
	FileName      string
	ContentLength int64
	Url           string
}

func newDownload(ver string, downloadPath string) (*Download, error) {
	def, err := defaultCfg()
	if err != nil {
		return nil, err
	}

	if len(ver) != 0 {
		def.version = ver
	}

	if len(downloadPath) != 0 {
		def.downloadPath = downloadPath
	}

	return def, nil
}

func defaultCfg() (*Download, error) {
	list, err := NewList()
	if err != nil {
		return nil, err
	}
	return &Download{
		version:      list.LastVersion(),
		downloadPath: os.TempDir(),
	}, nil
}

func (d *Download) Download(ctx context.Context) error {
	target := d.getFileInfo(ctx, fmt.Sprintf(filePattern, d.version, runtime.GOARCH))
	if target == nil {
		return nil
	}

	if err := target.download(ctx, d.downloadPath); err != nil {
		return err
	}

	return nil
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

func (t *Target) download(ctx context.Context, downloadPath string) error {
	dl := pget.New()
	args := []string{
		"-o",
		downloadPath,
		"-t",
		"30",
		t.Url,
	}
	if err := dl.Run(ctx, "", args); err != nil {
		return err
	}
	return nil
}
