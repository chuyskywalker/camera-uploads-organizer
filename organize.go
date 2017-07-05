package main

import (
	"fmt"
	"regexp"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	// scan CWD for all files in iPhone naming pattern
	// for matching files, extract year/month
	// move file to year/month folder
	// catch no such directory and attempt to create
	// move again
	// done

	files, _ := ioutil.ReadDir("./")
	rx := regexp.MustCompile("(?i)^([0-9]{4})-([0-9]{2})-[0-9]{2}.+\\.(mov|png|jpg|jpeg|gif|avi|mp4)$") // won't be done multiple times
	for _, f := range files {
		if !f.IsDir() {
			r := rx.FindStringSubmatch(f.Name())
			if len(r) == 0 {
				// no match, carry on
				continue
			}
			mverr := os.Rename(f.Name(), filepath.Join(r[1], r[2], f.Name()))
			if mverr != nil {
				os.MkdirAll(filepath.Join(r[1], r[2]), 0777)
				mverr2 := os.Rename(f.Name(), filepath.Join(r[1], r[2], f.Name()))
				if mverr2 != nil {
					fmt.Println("Could not create directory/move file")
				}
			}
		}
	}
}
