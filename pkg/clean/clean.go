package clean

import (
	"log"
	"os"

	"github.com/zaunist/k/pkg/conf"
)

func Do() {
	clean()
}

func clean() {
	if err := os.RemoveAll(conf.VersonDir); err != nil {
		log.Fatalf("Remove versionDir delete: %v", err)
	}
	if err := os.RemoveAll(conf.BinDir); err != nil {
		log.Fatalf("Remove binDir delete: %v", err)
	}
}
