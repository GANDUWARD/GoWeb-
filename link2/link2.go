package main

import (
	"fmt"
	"os"
)

func main() {
	fp, err := os.Create("./link2.TXT")
	if err != nil {
		fmt.Println(err)
	}
	defer fp.Close()
	os.Chmod("link2.TXT", 0777)
	err = os.Symlink("link2.TXT", "link3.TXT")
	if err != nil {
		fmt.Println(err)
	}
}
