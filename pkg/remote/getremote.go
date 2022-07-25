package remote

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func List() []string {
	f, err := os.Open("./pkg/remote/versions.txt")
	if err != nil {
		log.Fatalf("open versions.txt error:%v", err)
		return nil
	}
	defer f.Close()
	ver := []string{}
	buf := bufio.NewReader(f)

	for {
		v, _, err := buf.ReadLine()
		if err != nil && err != io.EOF {
			log.Fatalf("read versions.txt error:%v", err)
		}
		if err == io.EOF {
			break
		}
		ver = append(ver, string(v))
	}
	for _, v := range ver {
		fmt.Println(v)
	}
	return ver
}
