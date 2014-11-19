package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func initCommand() {
	base := []byte(`base:     dev-spike          # image to launch Docker with
repopath: github.com/Xe/dev  # repo path for mounting $CURDIR
golang:   false              # Go has a more opinionated package store
ssh:      true               # pass through ssh keys?
user:     xena               # user in the docker container
projname: spike              # project name

# Overlay is the dockerfile to overlay into this for ` + "dev establish" + `
overlay: |
  FROM xena/dev-moonscript
  RUN echo "Hi mom!"`)

	err := ioutil.WriteFile(*manifest, base, 0666)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Wrote default manifest to " + *manifest)
}
