package use

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/zaunist/k/pkg/conf"
)

func Do(version string) {
	use(version)
}

func use(version string) {
	path := conf.BinDir
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
