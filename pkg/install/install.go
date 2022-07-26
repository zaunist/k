package install

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/zaunist/k/pkg/conf"
	"github.com/zaunist/k/pkg/use"
)

func Do(version string) {
	install(version)
}

func install(version string) {
	address := "https://dl.k8s.io/release/" + version + "/bin/linux/amd64/kubectl"
	path := filepath.Join(conf.VersonDir, version)

	os.MkdirAll(path, os.ModePerm)
	f, err := os.OpenFile(filepath.Join(path, "kubectl"), os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatalf("dowanload error: %v", err)
	}
	defer f.Close()
	resp, err := http.Get(address)
	if err != nil {
		log.Fatalf("get kubectl error: %v", err)
	}
	defer resp.Body.Close()

	io.Copy(f, resp.Body)

	use.Do(version)
}
