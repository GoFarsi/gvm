package internal

import (
	"archive/tar"
	"compress/gzip"
	"github.com/GoFarsi/gvm/errors"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func checkGoInstalled() (string, bool) {
	goFile, err := findGoVersion()
	if err != nil {
		return "", false
	}

	resp, err := exec.Command(goFile, "version").Output()
	if err != nil {
		return "", false
	}

	re := regexp.MustCompile("(\\d+\\.\\d+(?:\\.\\d+)?)")
	match := re.FindStringSubmatch(string(resp))

	if len(match) != 0 {
		return match[0], true
	}

	return "", false
}

func installGo(filePath string) error {
	if !hasSudoAccess() {
		return errors.ERR_SUDO_ACCESS
	}

	if !removedOldGo() {
		return errors.ERR_CANT_REMOVE_OLD_VERSION
	}

	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	gzf, err := gzip.NewReader(f)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(gzf)

	count := 0
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}

		target := filepath.Join("/usr/local", header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					log.Println(err)
					os.Exit(1)
				}
			}
		case tar.TypeReg:
			log.Println("(", count, ")", "Name: ", target)

			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			// copy over contents
			if _, err := io.Copy(f, tarReader); err != nil {
				return err
			}

			f.Close()

		default:
			log.Printf("unable to untar file %s with type %c", target, header.Typeflag)
		}
		count++
	}

	if err := setGoEnvPath(); err != nil {
		return errors.ERR_CANT_SET_ENV
	}

	log.Printf("\ngo installed in %s path, please run 'source ~/.profile' for commit env.", goInstallPath)

	return nil
}

func findGoVersion() (string, error) {
	path, err := exec.LookPath("go")
	if err != nil {
		if _, err := os.Stat(goInstallPath + "/bin/go"); err != nil {
			return "", err
		}
		path = goInstallPath + "/bin/go"
	}

	return path, nil
}

func hasSudoAccess() bool {
	resp, err := exec.Command("groups").Output()
	if err != nil {
		return false
	}

	if strings.Contains(string(resp), "root") || strings.Contains(string(resp), "sudo") {
		return true
	}

	return false
}

func removedOldGo() bool {
	if _, err := os.Stat(goInstallPath + "/bin/go"); err != nil {
		return true
	}

	if err := os.RemoveAll(goInstallPath); err != nil {
		return false
	}

	return true
}

func setGoEnvPath() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	file, err := os.OpenFile(home+"/.profile", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(env); err != nil {
		return err
	}
	return nil
}
