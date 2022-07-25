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
	os.MkdirAll(path, 0755)
	targetVersion := filepath.Join(conf.VersonDir, version, "kubectl")

	symlink := filepath.Join(path, "kubectl")

	if finfo, err := os.Stat(targetVersion); err != nil || finfo.IsDir() {
		log.Fatalf(err.Error())
		return
	}

	err := os.Remove(symlink)
	if err != nil {
		log.Println("delete old symlink failed")
	}
	err = os.Symlink(targetVersion, symlink)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("Use ", version)
}
