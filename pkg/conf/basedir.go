package conf

import (
	"log"
	"os/user"
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
	VersonDir = baseDir + "/.k/versions/"
	BinDir = baseDir + "/.k/bin/"
}
