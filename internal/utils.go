package internal

import (
	"archive/tar"
	"compress/gzip"
	"github.com/GoFarsi/gvm/errors"
	"io"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"regexp"
	"strconv"
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

	if !removeOldData(goInstallPath, "/bin/go") {
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

	setGoEnvPath()

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

func setGoEnvPath() {
	homes, _, _ := getLinuxHomeUsers()

	for _, home := range homes {
		targetPath := home + "/.profile"
		b, err := os.ReadFile(targetPath)
		if err != nil {
			continue
		}

		if !strings.Contains(string(b), "/go/bin") {
			file, err := os.OpenFile(targetPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				continue
			}
			if _, err := file.WriteString(env); err != nil {
				continue
			}
		}
	}
}

func getLinuxHomeUsers() ([]string, int, int) {
	b, err := exec.Command("ls", "/home").Output()
	if err != nil {
		log.Fatalln(err)
	}

	users := strings.Split(string(b), "\n")
	homeList := []string{"/root"}
	userId := 0
	groupId := 0
	for _, u := range users {
		if u == "lost+found" || len(u) == 0 {
			continue
		}
		user, err := user.Lookup(u)
		if err != nil {
			continue
		}

		if u != "root" {
			uIdInt, _ := strconv.Atoi(user.Uid)
			gIdInt, _ := strconv.Atoi(user.Gid)
			userId = uIdInt
			groupId = gIdInt
		}

		homeList = append(homeList, user.HomeDir)
	}
	return homeList, userId, groupId
}

func removeOldData(path string, file string) bool {
	if _, err := os.Stat(path + file); err != nil {
		return true
	}

	if err := os.RemoveAll(path); err != nil {
		return false
	}

	if _, err := os.Stat(path + file); err != nil {
		return true
	} else {
		if err := exec.Command("rm", "-r", path).Start(); err != nil {
			return false
		}
	}

	return true
}
