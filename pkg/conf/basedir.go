package conf

import (
	"log"
	"os/user"
	"path/filepath"
)

var (
	baseDir   string
	VersonDir string
	BinDir    string
)

func Init() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	baseDir = u.HomeDir
	VersonDir = filepath.Join(baseDir, ".k", "versions")
	BinDir = filepath.Join(baseDir, ".k", "bin")
}
