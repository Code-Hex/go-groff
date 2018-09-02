package main

import (
	"io/ioutil"
	"os"

	groff "github.com/Code-Hex/go-groff"
	"github.com/k0kubun/pp"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	f, err := os.Open("file.man")
	if err != nil {
		return err
	}
	defer f.Close()
	src, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	nodes, err := groff.Parse(string(src))
	if err != nil {
		return err
	}
	pp.Println(nodes)
	return nil
}
