package ls

import (
	"fmt"
	"io/ioutil"
)

func List(versionsDir string) {
	infos, err := ioutil.ReadDir(versionsDir)
	if err != nil || len(infos) <= 0 {
		fmt.Printf("No version installed yet\n\n")
		return
	}
	for _, file := range infos {
		fmt.Println(file.Name())
	}

}
