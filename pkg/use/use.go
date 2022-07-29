package use

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/zaunist/k/pkg/conf"
)

func Do(version string) {
	use(version)
}

func use(version string) {
	path := conf.BinDir

	if runtime.GOOS == "windows" {
		// Windows 10下无特权用户无法创建符号链接，优先调用mklink /j创建'目录联接'
		// Windows 的逻辑放在前面，避免提前创建目录导致
		// fmt.Println(path)
		// fmt.Println(filepath.Join(conf.VersonDir, version))
		if err := os.RemoveAll(conf.BinDir); err != nil {
			log.Fatalf("delete bin directory failed: %v", err)
		}
		if err := exec.Command("cmd", "/c", "mklink", "/j", path, filepath.Join(conf.VersonDir, version)).Run(); err != nil {
			log.Fatalf("create directory link failed in windows: %v", err)
		}
		fmt.Println("Use ", version)
		return
	}

	if err := os.MkdirAll(path, 0755); err != nil {
		log.Printf("create directory failed: %v", err)
	}
	targetVersion := filepath.Join(conf.VersonDir, version, "kubectl")

	symlink := filepath.Join(path, "kubectl")

	if finfo, err := os.Stat(targetVersion); err != nil || finfo.IsDir() {
		log.Fatalf(err.Error())
		return
	}

	if err := os.Remove(symlink); err != nil {
		log.Println("delete old symlink failed")
	}
	if err := os.Symlink(targetVersion, symlink); err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("Use ", version)
}
