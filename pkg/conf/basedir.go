package conf

import (
	"os"
)

var (
	baseDir   string
	VersonDir string
	BinDir    string
)

func Init() {
	baseDir = os.Getenv("HOME")
	VersonDir = baseDir + "/.k/versions/"
	BinDir = baseDir + "/.k/bin/"
}
