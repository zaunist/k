package install

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/zaunist/k/pkg/conf"
	"github.com/zaunist/k/pkg/use"
)

var address string

func Do(version string) {
	install(version)
}

func install(version string) {
	switch runtime.GOOS {
	case "linux":
		address = fmt.Sprintf("https://dl.k8s.io/release/%s/bin/%s/%s/kubectl", version, runtime.GOOS, runtime.GOARCH)
	case "windows":
		address = fmt.Sprintf("https://dl.k8s.io/release/%s/bin/%s/%s/kubectl.exe", version, runtime.GOOS, runtime.GOARCH)
	case "darwin":
		address = fmt.Sprintf("https://dl.k8s.io/release/%s/bin/%s/%s/kubectl", version, runtime.GOOS, runtime.GOARCH)
	}
	path := filepath.Join(conf.VersonDir, version)

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Printf("create directory failed: %v", err)
	}

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

	if _, err := io.Copy(f, resp.Body); err != nil {
		log.Fatalf("download error: %v", err)
	}

	use.Do(version)
}
