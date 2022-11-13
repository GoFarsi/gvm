package internal

import (
	"context"
	"fmt"
	"github.com/GoFarsi/gvm/errors"
	"github.com/Songmu/prompter"
	vers "github.com/hashicorp/go-version"
	"log"
	"os/exec"
	"path"
)

func InstallGo(ctx context.Context, version string, backup string) error {
	if !prompter.YN(fmt.Sprintf(warningSudoAccess, INSTALL), false) {
		return nil
	}

	ver, ok := checkGoInstalled()
	if ok {
		if !prompter.YN(fmt.Sprintf(warningInstalledGo, ver, goInstallPath, INSTALL), false) {
			return nil
		}
	}

	d, err := newDownload(version, backup)
	if err != nil {
		return err
	}

	t, err := d.download(ctx)
	if err != nil {
		return err
	}

	if err := installGo(path.Join(t.Path, t.FileName)); err != nil {
		return err
	}

	return nil
}

func UpgradeGo(ctx context.Context, version string, backup string) error {
	if !prompter.YN(fmt.Sprintf(warningSudoAccess, UPGRADE), false) {
		return nil
	}

	ver, ok := checkGoInstalled()
	installedVer, err := vers.NewVersion(ver)
	if err != nil {
		return errors.ERR_READ_VERSION_CODE
	}
	if len(version) != 0 {
		reqVer, err := vers.NewVersion(version)
		if err != nil {
			return errors.ERR_READ_VERSION_CODE
		}
		if ok {
			if reqVer.LessThan(installedVer) {
				return errors.ERR_UPGRADE_VERSION
			}
		}
	} else {
		list, err := NewList()
		if err != nil {
			return err
		}

		latestVer, err := vers.NewVersion(list.LastVersion())
		if err != nil {
			return errors.ERR_READ_VERSION_CODE
		}

		if latestVer.Equal(installedVer) {
			log.Printf("your go version is latest.")
			return nil
		}
	}

	d, err := newDownload(version, backup)
	if err != nil {
		return err
	}

	t, err := d.download(ctx)
	if err != nil {
		return err
	}

	if err := installGo(path.Join(t.Path, t.FileName)); err != nil {
		return err
	}

	return nil
}

func DowngradeGo(ctx context.Context, version string, backup string) error {
	if !prompter.YN(fmt.Sprintf(warningSudoAccess, DOWNGRADE), false) {
		return nil
	}

	if err := exec.Command("sudo", "-i").Start(); err != nil {
		return err
	}

	ver, ok := checkGoInstalled()
	installedVer, err := vers.NewVersion(ver)
	if err != nil {
		return errors.ERR_READ_VERSION_CODE
	}
	if len(version) != 0 {
		reqVer, err := vers.NewVersion(version)
		if err != nil {
			return errors.ERR_READ_VERSION_CODE
		}
		if ok {
			if reqVer.GreaterThan(installedVer) {
				return errors.ERR_DOWNGRADE_VERSION
			}
		}
	} else {
		list, err := NewList()
		if err != nil {
			return err
		}

		versions := list.GetVersions()
		for i, v := range versions {
			vr, err := vers.NewVersion(v)
			if err != nil {
				return errors.ERR_READ_VERSION_CODE
			}
			if vr.Equal(installedVer) {
				version = versions[i+1]
				break
			}
		}
	}

	d, err := newDownload(version, backup)
	if err != nil {
		return err
	}

	t, err := d.download(ctx)
	if err != nil {
		return err
	}

	if err := installGo(path.Join(t.Path, t.FileName)); err != nil {
		return err
	}

	return nil
}
