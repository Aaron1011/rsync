package main

import (
	"github.com/Aaron1011/rsync/rsync"
	"io/ioutil"
	"fmt"
)

func main() {
	file, _ := ioutil.ReadFile("mydata")
	fmt.Printf("File: %+v, ", rsync.NewRsyncFile(file))
}
