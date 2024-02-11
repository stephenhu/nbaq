package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)


func version() string {
	return fmt.Sprintf("Nbaq v%s", APP_VERSION)
} // version


func getFiles(fp string) []os.DirEntry {

	dirs, err := os.ReadDir(fp)

	if err != nil {
		
		log.Println(err)
		return nil

	} else {
		return dirs
	}

} // getFiles


func getLatest() string {

	var ret string

	files := getFiles(filepath.Join(src, WAREHOUSE_DIR))

	for _, f := range files {

		if len(ret) == 0 {
			ret = f.Name()
		} else if ret < f.Name(){
			ret = f.Name()
		}

	}

	return ret

} // getLatest
