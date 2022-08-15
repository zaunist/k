package uninstall

import (
	"fmt"
	version2 "github.com/zaunist/k/pkg/version"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/zaunist/k/pkg/conf"
)

func Do(version string) {
	v := version2.NormalizedBuildVersion(version)
	if v == "" {
		log.Fatalf("Proviede version is invalid, please check again")
		return
	}
	uninstall(v)
}

func uninstall(version string) {
	infos, err := ioutil.ReadDir(conf.VersonDir)
	if err != nil || len(infos) <= 0 {
		fmt.Printf("No version installed yet\n\n")
		return
	}
	for _, file := range infos {
		if version == file.Name() {
			// delete symlink first
			symlink := filepath.Join(conf.BinDir, "kubectl")
			if err := os.Remove(symlink); err != nil {
				log.Println("delete old symlink failed")
			}
			if err := os.RemoveAll(filepath.Join(conf.VersonDir, version)); err != nil {
				log.Fatalf("Uninstall failed: %v", err)
			}
			fmt.Println("success uninstall: ", file.Name())
		}
	}
}
