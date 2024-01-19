package main

import (
	"fmt"
	"log"
	"os"
)


func version() string {
	return fmt.Sprintf("NBAQ v%s", APP_VERSION)
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
